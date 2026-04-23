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

const (
	// Workspace invite defaults from OpenAPI.
	DefaultInviteExpiresInDays = 14
	// Incident paging defaults from OpenAPI when target_user_id is used.
	DefaultAckTimeoutSeconds = 60
	// Shift merge default from OpenAPI.
	DefaultMaxGapMinutes = 0
)

type PatchMeRequest struct {
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email,omitempty"`
}

type CreateWorkspaceRequest struct {
	Name string `json:"name"`
}

type PatchCurrentWorkspaceRequest struct {
	Name string `json:"name"`
}

type SwitchCurrentWorkspaceRequest struct {
	OrganizationID string `json:"organization_id"`
}

type PatchCurrentWorkspaceMemberRoleRequest struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type CreateCurrentWorkspaceInviteRequest struct {
	Email         string  `json:"email"`
	Role          *string `json:"role,omitempty"`
	ExpiresInDays *int    `json:"expires_in_days,omitempty"`
}

type EscalationPolicyCreateStep struct {
	Order             int      `json:"order"`
	TargetType        string   `json:"target_type"`
	TargetID          string   `json:"target_id"`
	AckTimeoutSeconds *int     `json:"ack_timeout_seconds,omitempty"`
	NotifyVia         []string `json:"notify_via,omitempty"`
}

type CreateEscalationPolicyRequest struct {
	TeamID                *string                      `json:"team_id,omitempty"`
	Name                  string                       `json:"name"`
	RepeatEnabled         *bool                        `json:"repeat_enabled,omitempty"`
	RepeatIntervalSeconds *int                         `json:"repeat_interval_seconds,omitempty"`
	MaxLoops              *int                         `json:"max_loops,omitempty"`
	Steps                 []EscalationPolicyCreateStep `json:"steps"`
}

type PatchEscalationPolicyRequest struct {
	TeamID                *string                      `json:"team_id,omitempty"`
	Name                  *string                      `json:"name,omitempty"`
	RepeatEnabled         *bool                        `json:"repeat_enabled,omitempty"`
	RepeatIntervalSeconds *int                         `json:"repeat_interval_seconds,omitempty"`
	MaxLoops              *int                         `json:"max_loops,omitempty"`
	Steps                 []EscalationPolicyCreateStep `json:"steps,omitempty"`
}

type CreateScheduleRequest struct {
	TeamID   string `json:"team_id"`
	Name     string `json:"name"`
	Timezone string `json:"timezone,omitempty"`
}

type PatchScheduleRequest struct {
	Name     string `json:"name,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

type CreateScheduleShiftRequest struct {
	UserID   string `json:"user_id"`
	Level    int    `json:"level"`
	StartsAt string `json:"starts_at"`
	EndsAt   string `json:"ends_at"`
}

type MergeScheduleShiftsRequest struct {
	ShiftIDs      []string `json:"shift_ids"`
	MaxGapMinutes *int     `json:"max_gap_minutes,omitempty"`
}

type SensorEventLocation struct {
	SiteID     string `json:"site_id,omitempty"`
	BuildingID string `json:"building_id,omitempty"`
	Floor      string `json:"floor,omitempty"`
	Zone       string `json:"zone,omitempty"`
}

type CreateSensorEventRequest struct {
	IdempotencyKey string               `json:"idempotency_key"`
	SensorID       string               `json:"sensor_id"`
	SensorType     string               `json:"sensor_type,omitempty"`
	DetectedAt     string               `json:"detected_at"`
	SignalType     string               `json:"signal_type"`
	Confidence     float64              `json:"confidence"`
	Location       *SensorEventLocation `json:"location,omitempty"`
	Metadata       map[string]any       `json:"metadata,omitempty"`
}

type CreateIncidentRequest struct {
	Name               string  `json:"name"`
	IdempotencyKey     string  `json:"idempotency_key"`
	SeverityID         string  `json:"severity_id"`
	CategoryID         *string `json:"category_id,omitempty"`
	Visibility         string  `json:"visibility"`
	Source             string  `json:"source,omitempty"`
	SourceRef          *string `json:"source_ref,omitempty"`
	InitialStatus      string  `json:"initial_status,omitempty"`
	AutoPage           *bool   `json:"auto_page,omitempty"`
	EscalationPolicyID *string `json:"escalation_policy_id,omitempty"`
	PageReason         *string `json:"page_reason,omitempty"`
}

type PatchIncidentRequest struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	SeverityID  string  `json:"severity_id,omitempty"`
	CategoryID  *string `json:"category_id,omitempty"`
	Visibility  string  `json:"visibility,omitempty"`
	Status      string  `json:"status,omitempty"`
}

type ResolveIncidentRequest struct {
	ResolutionCode string `json:"resolution_code,omitempty"`
	Summary        string `json:"summary"`
}

type DeclineIncidentRequest struct {
	Reason string `json:"reason,omitempty"`
}

type MergeIncidentRequest struct {
	TargetIncidentID string `json:"target_incident_id"`
	Reason           string `json:"reason,omitempty"`
}

type CreateIncidentNoteRequest struct {
	Body string `json:"body"`
}

type StartIncidentPagingRequest struct {
	EscalationPolicyID *string  `json:"escalation_policy_id,omitempty"`
	TargetUserID       *string  `json:"target_user_id,omitempty"`
	Reason             string   `json:"reason,omitempty"`
	AckTimeoutSeconds  *int     `json:"ack_timeout_seconds,omitempty"`
	NotifyVia          []string `json:"notify_via,omitempty"`
}

type CreateSuppressionRequest struct {
	ScopeType string `json:"scope_type"`
	ScopeID   string `json:"scope_id"`
	Reason    string `json:"reason"`
	StartsAt  string `json:"starts_at"`
	EndsAt    string `json:"ends_at"`
}
