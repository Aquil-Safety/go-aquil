package aquil

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClientDefaults(t *testing.T) {
	token := "valid-token"
	client := NewClient(&token)

	if client == nil {
		t.Fatal("expected NewClient to return non-nil client")
	}
	if client.token != token {
		t.Fatalf("expected token %q, got %q", token, client.token)
	}
	if client.baseURL != defaultBaseURL {
		t.Fatalf("expected baseURL %q, got %q", defaultBaseURL, client.baseURL)
	}
	if client.httpClient == nil {
		t.Fatal("expected httpClient to be initialized")
	}
}

func TestListIncidentsBuildsQueryAndAuthorization(t *testing.T) {
	token := "jwt-token"
	client := NewClient(&token)

	var gotPath string
	var gotAuth string
	var gotQuery string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		gotQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer ts.Close()

	client.SetBaseURL(ts.URL)

	pageSize := 50
	_, err := client.ListIncidents(context.Background(), ListIncidentsParams{
		PageSize:      &pageSize,
		StatusOneOf:   "new,investigating",
		SeverityNotIn: "sev4",
	})
	if err != nil {
		t.Fatalf("ListIncidents error: %v", err)
	}

	if gotPath != "/v1/incidents" {
		t.Fatalf("expected path /v1/incidents, got %s", gotPath)
	}
	if gotAuth != "Bearer jwt-token" {
		t.Fatalf("expected bearer token header, got %q", gotAuth)
	}
	if !strings.Contains(gotQuery, "page_size=50") {
		t.Fatalf("expected page_size in query, got %s", gotQuery)
	}
	if !strings.Contains(gotQuery, "status%5Bone_of%5D=new%2Cinvestigating") {
		t.Fatalf("expected status filter in query, got %s", gotQuery)
	}
	if !strings.Contains(gotQuery, "severity%5Bnot_in%5D=sev4") {
		t.Fatalf("expected severity filter in query, got %s", gotQuery)
	}
}

func TestCreateSensorEventSetsInternalHeadersAndBody(t *testing.T) {
	token := "jwt-token"
	client := NewClient(&token)
	client.SetInternalKey("internal-key")

	var gotOrgHeader string
	var gotInternalKey string
	var gotMethod string
	var gotPath string
	var gotBody map[string]any

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotOrgHeader = r.Header.Get("X-Organization-Id")
		gotInternalKey = r.Header.Get("X-Internal-Key")
		gotMethod = r.Method
		gotPath = r.URL.Path
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer ts.Close()

	client.SetBaseURL(ts.URL)

	_, err := client.CreateSensorEvent(context.Background(), "org-123", CreateSensorEventRequest{
		IdempotencyKey: "evt-1",
		SensorID:       "sensor-1",
		DetectedAt:     "2026-04-22T20:00:00Z",
		SignalType:     "gunshot",
		Confidence:     0.98,
	})
	if err != nil {
		t.Fatalf("CreateSensorEvent error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Fatalf("expected POST, got %s", gotMethod)
	}
	if gotPath != "/v1/sensor-events" {
		t.Fatalf("expected /v1/sensor-events, got %s", gotPath)
	}
	if gotOrgHeader != "org-123" {
		t.Fatalf("expected X-Organization-Id org-123, got %q", gotOrgHeader)
	}
	if gotInternalKey != "internal-key" {
		t.Fatalf("expected X-Internal-Key internal-key, got %q", gotInternalKey)
	}
	if gotBody["sensor_id"] != "sensor-1" {
		t.Fatalf("expected sensor_id in body, got %#v", gotBody)
	}
}

func TestReturnsAPIErrorOnNon2xx(t *testing.T) {
	token := "jwt-token"
	client := NewClient(&token)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"error":"duplicate"}`))
	}))
	defer ts.Close()

	client.SetBaseURL(ts.URL)

	_, err := client.CreateIncident(context.Background(), CreateIncidentRequest{
		Name:           "test",
		IdempotencyKey: "idem-1",
		SeverityID:     "sev-1",
		Visibility:     "organization",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusConflict {
		t.Fatalf("expected status 409, got %d", apiErr.StatusCode)
	}
	if string(apiErr.Body) != `{"error":"duplicate"}` {
		t.Fatalf("unexpected error body: %s", string(apiErr.Body))
	}
}
