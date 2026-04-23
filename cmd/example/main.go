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
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	aquil "github.com/aquil-safety/go-aquil"
)

func main() {
	token := os.Getenv("AQUIL_BEARER_TOKEN")
	if token == "" {
		fmt.Println("Set AQUIL_BEARER_TOKEN to a valid user JWT and rerun.")
		return
	}

	client := aquil.NewClient(&token)
	fmt.Println("Aquil Safety client created.")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Discover a valid severity_id for this workspace.
	severitiesResp, err := client.ListSeverities(ctx)
	if err != nil {
		fmt.Printf("Error listing severities: %v\n", err)
		return
	}

	var severities struct {
		OK   bool `json:"ok"`
		Data []struct {
			ID  string `json:"id"`
			Key string `json:"key"`
		} `json:"data"`
	}
	if err := json.Unmarshal(severitiesResp.Body, &severities); err != nil {
		fmt.Printf("Error parsing severities response: %v\n", err)
		return
	}
	if len(severities.Data) == 0 {
		fmt.Println("No severities found for this workspace.")
		return
	}

	severityID := severities.Data[0].ID
	fmt.Printf("Using severity: %s (%s)\n", severities.Data[0].Key, severityID)

	incidentResp, err := client.CreateIncident(ctx, aquil.CreateIncidentRequest{
		Name:           "Example incident from Go SDK",
		IdempotencyKey: fmt.Sprintf("go-sdk-example-%d", time.Now().Unix()),
		SeverityID:     severityID,
		Visibility:     "organization",
		Source:         "manual",
	})
	if err != nil {
		fmt.Printf("Error creating incident: %v\n", err)
		return
	}

	var incidentCreate struct {
		OK   bool `json:"ok"`
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(incidentResp.Body, &incidentCreate); err != nil {
		fmt.Printf("Error parsing incident creation response: %v\n", err)
		return
	}
	incidentID := incidentCreate.Data.ID
	fmt.Printf("Created incident with ID: %s\n", incidentID)

	fmt.Println("Gathering escalation policies...")

	escalationResp, err := client.ListEscalationPolicies(ctx)
	if err != nil {
		fmt.Printf("Error listing escalation policies: %v\n", err)
		return
	}

	fmt.Printf("Escalation policies:\n%s\n", string(escalationResp.Body))

	var escalationPolicies struct {
		OK   bool `json:"ok"`
		Data []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"data"`
	}
	if err := json.Unmarshal(escalationResp.Body, &escalationPolicies); err != nil {
		fmt.Printf("Error parsing escalation policies response: %v\n", err)
		return
	}

	if len(escalationPolicies.Data) == 0 {
		fmt.Println("No escalation policies found for this workspace.")
		return
	}

	escalationPolicyID := escalationPolicies.Data[0].ID // required for paging

	fmt.Printf("Using escalation policy: %s (%s)\n", escalationPolicies.Data[0].Name, escalationPolicyID)

	pagingResp, err := client.StartIncidentPaging(
		ctx,
		incidentID,
		aquil.StartIncidentPagingRequest{
			EscalationPolicyID: &escalationPolicyID,
		},
	)

	if err != nil {
		fmt.Printf("Error starting incident paging: %v\n", err)
		return
	}

	fmt.Printf("Started paging with response:\n%s\n", string(pagingResp.Body))
}
