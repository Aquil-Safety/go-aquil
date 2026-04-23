# go-aquil

Aquil Safety Go SDK provides a convenient interface for interacting with the Aquil API in Go applications. It abstracts away raw HTTP requests and offers typed models for request bodies and query parameters, making it easier to integrate Aquil into your Go projects.

## Go SDK Installation

```bash
go get github.com/aquil-safety/go-aquil
```

## Authentication

`NewClient` expects a **Bearer JWT** (not an internal ingestion key):

```go
token := "your_api_token_here"
client := aquil.NewClient(&token)
```

## SDK-First Usage

### 1. Typed request bodies

Use typed structs from [types.go](/Users/devin/Documents/Projects/Aquili%20Safety/go-sdk/types.go) instead of raw maps.

```go
incidentResp, err := client.CreateIncident(ctx, aquil.CreateIncidentRequest{
	Name:           "Example incident from Go SDK",
	IdempotencyKey: fmt.Sprintf("go-aquil-example-%d", time.Now().Unix()),
	SeverityID:     severityID,
	Visibility:     "organization",
	Source:         "manual",
})
```

### 2. Typed query params

```go
pageSize := 25
resp, err := client.ListIncidents(ctx, aquil.ListIncidentsParams{
	PageSize:      &pageSize,
	StatusOneOf:   "new,open",
	SeverityNotIn: "sev4",
})
```

### 3. OpenAPI defaults exposed as SDK constants

- `aquil.DefaultInviteExpiresInDays` (`14`)
- `aquil.DefaultAckTimeoutSeconds` (`60`)
- `aquil.DefaultMaxGapMinutes` (`0`)

## Main Typed Models

- Incident: `CreateIncidentRequest`, `PatchIncidentRequest`, `ResolveIncidentRequest`, `DeclineIncidentRequest`, `MergeIncidentRequest`, `CreateIncidentNoteRequest`, `StartIncidentPagingRequest`
- Workspaces: `CreateWorkspaceRequest`, `PatchCurrentWorkspaceRequest`, `SwitchCurrentWorkspaceRequest`, `PatchCurrentWorkspaceMemberRoleRequest`, `CreateCurrentWorkspaceInviteRequest`
- Escalations: `CreateEscalationPolicyRequest`, `PatchEscalationPolicyRequest`, `EscalationPolicyCreateStep`
- Schedules: `CreateScheduleRequest`, `PatchScheduleRequest`, `CreateScheduleShiftRequest`, `MergeScheduleShiftsRequest`
- Sensor events: `CreateSensorEventRequest`, `SensorEventLocation`
- Suppressions: `CreateSuppressionRequest`

## End-to-end Incident Example

```go
ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
defer cancel()

severitiesResp, err := client.ListSeverities(ctx)
if err != nil {
	log.Fatal(err)
}

var severities struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}
if err := json.Unmarshal(severitiesResp.Body, &severities); err != nil {
	log.Fatal(err)
}
if len(severities.Data) == 0 {
	log.Fatal("no severities configured")
}

incidentResp, err := client.CreateIncident(ctx, aquil.CreateIncidentRequest{
	Name:           "Example incident from Go SDK",
	IdempotencyKey: fmt.Sprintf("go-aquil-example-%d", time.Now().Unix()),
	SeverityID:     severities.Data[0].ID,
	Visibility:     "organization",
	Source:         "manual",
})
if err != nil {
	log.Fatal(err)
}

fmt.Println(string(incidentResp.Body))
```

View examples in [cmd/example](/cmd/example/) for more usage patterns.

View official API docs for more details on request/response formats: https://docs.aquil.com/go-sdk
