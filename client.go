/*
* COPYRIGHT 2026 AQUIL SAFETY LLC
*
* This file is protected under the Apache 2.0 License. Unauthorized use,
* reproduction, or distribution of this file is strictly prohibited. For
* more information, please refer to the LICENSE file in the root directory of
* this project.
*
* @file
 */

package aquil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const defaultBaseURL = "https://api.aquilsafety.com"

// Client is the Aquil SDK client.
type Client struct {
	httpClient  *http.Client
	token       string
	internalKey string
	baseURL     string
}

// Response is a raw HTTP response payload returned by SDK methods.
type Response struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

// APIError is returned for non-2xx API responses.
type APIError struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	if len(e.Body) == 0 {
		return fmt.Sprintf("aquil API error: status=%d", e.StatusCode)
	}
	return fmt.Sprintf("aquil API error: status=%d body=%s", e.StatusCode, string(e.Body))
}

// NewClient creates an SDK client configured with bearer auth.
func NewClient(token *string) *Client {
	tk := ""
	if token != nil {
		tk = *token
	}

	return &Client{
		httpClient: &http.Client{},
		token:      tk,
		baseURL:    defaultBaseURL,
	}
}

// SetHTTPClient allows dependency injection for tests/custom transports.
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	if httpClient != nil {
		c.httpClient = httpClient
	}
}

// SetBearerToken updates bearer auth token used for requests.
func (c *Client) SetBearerToken(token string) {
	c.token = token
}

// SetInternalKey sets X-Internal-Key header used by internal ingestion routes.
func (c *Client) SetInternalKey(internalKey string) {
	c.internalKey = internalKey
}

// SetBaseURL overrides API base URL (useful for local/dev/testing).
func (c *Client) SetBaseURL(baseURL string) {
	if strings.TrimSpace(baseURL) == "" {
		return
	}
	c.baseURL = strings.TrimRight(baseURL, "/")
}

func (c *Client) doJSON(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	body any,
	extraHeaders map[string]string,
) (*Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	endpoint := strings.TrimRight(c.baseURL, "/") + path
	if len(query) > 0 {
		endpoint += "?" + query.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		encodedBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encode request body: %w", err)
		}
		bodyReader = bytes.NewReader(encodedBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	if c.internalKey != "" {
		req.Header.Set("X-Internal-Key", c.internalKey)
	}
	for key, value := range extraHeaders {
		req.Header.Set(key, value)
	}

	httpResp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer httpResp.Body.Close()

	responseBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	resp := &Response{
		StatusCode: httpResp.StatusCode,
		Header:     httpResp.Header.Clone(),
		Body:       responseBody,
	}

	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: httpResp.StatusCode,
			Body:       responseBody,
		}
	}

	return resp, nil
}
