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
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// ListIncidentsParams maps documented query filters for GET /v1/incidents.
type ListIncidentsParams struct {
	PageSize           *int
	After              string
	FilterMode         string
	SortBy             string
	StatusOneOf        string
	StatusNotIn        string
	SeverityOneOf      string
	SeverityNotIn      string
	CreatedAtGTE       string
	CreatedAtLTE       string
	CreatedAtDateRange string
	UpdatedAtGTE       string
	UpdatedAtLTE       string
}

func (p ListIncidentsParams) values() url.Values {
	q := url.Values{}
	if p.PageSize != nil {
		q.Set("page_size", strconv.Itoa(*p.PageSize))
	}
	if p.After != "" {
		q.Set("after", p.After)
	}
	if p.FilterMode != "" {
		q.Set("filter_mode", p.FilterMode)
	}
	if p.SortBy != "" {
		q.Set("sort_by", p.SortBy)
	}
	if p.StatusOneOf != "" {
		q.Set("status[one_of]", p.StatusOneOf)
	}
	if p.StatusNotIn != "" {
		q.Set("status[not_in]", p.StatusNotIn)
	}
	if p.SeverityOneOf != "" {
		q.Set("severity[one_of]", p.SeverityOneOf)
	}
	if p.SeverityNotIn != "" {
		q.Set("severity[not_in]", p.SeverityNotIn)
	}
	if p.CreatedAtGTE != "" {
		q.Set("created_at[gte]", p.CreatedAtGTE)
	}
	if p.CreatedAtLTE != "" {
		q.Set("created_at[lte]", p.CreatedAtLTE)
	}
	if p.CreatedAtDateRange != "" {
		q.Set("created_at[date_range]", p.CreatedAtDateRange)
	}
	if p.UpdatedAtGTE != "" {
		q.Set("updated_at[gte]", p.UpdatedAtGTE)
	}
	if p.UpdatedAtLTE != "" {
		q.Set("updated_at[lte]", p.UpdatedAtLTE)
	}
	return q
}

func addOptionalBool(q url.Values, key string, value *bool) {
	if value != nil {
		q.Set(key, strconv.FormatBool(*value))
	}
}

func escapedPathParam(param string) string {
	return url.PathEscape(param)
}

// System
func (c *Client) GetHealth(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/health", nil, nil, nil)
}

func (c *Client) GetMe(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/me", nil, nil, nil)
}

func (c *Client) PatchMe(ctx context.Context, body PatchMeRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPatch, "/v1/me", nil, body, nil)
}

// Workspaces
func (c *Client) ListWorkspaces(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/workspaces", nil, nil, nil)
}

func (c *Client) CreateWorkspace(ctx context.Context, body CreateWorkspaceRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPost, "/v1/workspaces", nil, body, nil)
}

func (c *Client) GetCurrentWorkspace(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/workspaces/current", nil, nil, nil)
}

func (c *Client) PatchCurrentWorkspace(ctx context.Context, body PatchCurrentWorkspaceRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPatch, "/v1/workspaces/current", nil, body, nil)
}

func (c *Client) SwitchCurrentWorkspace(ctx context.Context, body SwitchCurrentWorkspaceRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPost, "/v1/workspaces/current/switch", nil, body, nil)
}

func (c *Client) ListCurrentWorkspaceMembers(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/workspaces/current/members", nil, nil, nil)
}

func (c *Client) PatchCurrentWorkspaceMemberRole(ctx context.Context, body PatchCurrentWorkspaceMemberRoleRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPatch, "/v1/workspaces/current/members", nil, body, nil)
}

func (c *Client) DeleteCurrentWorkspaceMember(ctx context.Context, userID string) (*Response, error) {
	return c.doJSON(ctx, http.MethodDelete, "/v1/workspaces/current/members/"+escapedPathParam(userID), nil, nil, nil)
}

func (c *Client) ListCurrentWorkspaceInvites(ctx context.Context, includeHistory *bool) (*Response, error) {
	q := url.Values{}
	addOptionalBool(q, "include_history", includeHistory)
	return c.doJSON(ctx, http.MethodGet, "/v1/workspaces/current/invites", q, nil, nil)
}

func (c *Client) CreateCurrentWorkspaceInvite(ctx context.Context, body CreateCurrentWorkspaceInviteRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPost, "/v1/workspaces/current/invites", nil, body, nil)
}

func (c *Client) RevokeCurrentWorkspaceInvite(ctx context.Context, inviteID string) (*Response, error) {
	path := "/v1/workspaces/current/invites/" + escapedPathParam(inviteID) + "/revoke"
	return c.doJSON(ctx, http.MethodPost, path, nil, nil, nil)
}

func (c *Client) ResendCurrentWorkspaceInvite(ctx context.Context, inviteID string) (*Response, error) {
	path := "/v1/workspaces/current/invites/" + escapedPathParam(inviteID) + "/resend"
	return c.doJSON(ctx, http.MethodPost, path, nil, nil, nil)
}

// Reference
func (c *Client) ListSeverities(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/severities", nil, nil, nil)
}

func (c *Client) ListIncidentCategories(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/categories", nil, nil, nil)
}

func (c *Client) ListIncidentStatuses(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/incident-statuses", nil, nil, nil)
}

// Escalation policies
func (c *Client) ListEscalationPolicies(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/escalation-policies", nil, nil, nil)
}

func (c *Client) CreateEscalationPolicy(ctx context.Context, body CreateEscalationPolicyRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPost, "/v1/escalation-policies", nil, body, nil)
}

func (c *Client) GetEscalationPolicy(ctx context.Context, policyID string) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/escalation-policies/"+escapedPathParam(policyID), nil, nil, nil)
}

func (c *Client) PatchEscalationPolicy(ctx context.Context, policyID string, body PatchEscalationPolicyRequest) (*Response, error) {
	path := "/v1/escalation-policies/" + escapedPathParam(policyID)
	return c.doJSON(ctx, http.MethodPatch, path, nil, body, nil)
}

func (c *Client) DeleteEscalationPolicy(ctx context.Context, policyID string) (*Response, error) {
	return c.doJSON(ctx, http.MethodDelete, "/v1/escalation-policies/"+escapedPathParam(policyID), nil, nil, nil)
}

// Teams / on-call
func (c *Client) ListTeams(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/teams", nil, nil, nil)
}

func (c *Client) GetTeam(ctx context.Context, teamID string) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/teams/"+escapedPathParam(teamID), nil, nil, nil)
}

func (c *Client) GetTeamOnCall(ctx context.Context, teamID string) (*Response, error) {
	path := "/v1/teams/" + escapedPathParam(teamID) + "/on-call"
	return c.doJSON(ctx, http.MethodGet, path, nil, nil, nil)
}

func (c *Client) ListUserTeams(ctx context.Context, userID string) (*Response, error) {
	path := "/v1/users/" + escapedPathParam(userID) + "/teams"
	return c.doJSON(ctx, http.MethodGet, path, nil, nil, nil)
}

// Schedules
func (c *Client) ListSchedules(ctx context.Context, teamID string) (*Response, error) {
	q := url.Values{}
	if teamID != "" {
		q.Set("team_id", teamID)
	}
	return c.doJSON(ctx, http.MethodGet, "/v1/schedules", q, nil, nil)
}

func (c *Client) CreateSchedule(ctx context.Context, body CreateScheduleRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPost, "/v1/schedules", nil, body, nil)
}

func (c *Client) GetSchedule(ctx context.Context, scheduleID string) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/schedules/"+escapedPathParam(scheduleID), nil, nil, nil)
}

func (c *Client) PatchSchedule(ctx context.Context, scheduleID string, body PatchScheduleRequest) (*Response, error) {
	path := "/v1/schedules/" + escapedPathParam(scheduleID)
	return c.doJSON(ctx, http.MethodPatch, path, nil, body, nil)
}

func (c *Client) DeleteSchedule(ctx context.Context, scheduleID string) (*Response, error) {
	return c.doJSON(ctx, http.MethodDelete, "/v1/schedules/"+escapedPathParam(scheduleID), nil, nil, nil)
}

func (c *Client) ListScheduleShifts(ctx context.Context, scheduleID string, activeOnly *bool) (*Response, error) {
	q := url.Values{}
	addOptionalBool(q, "active_only", activeOnly)
	path := "/v1/schedules/" + escapedPathParam(scheduleID) + "/shifts"
	return c.doJSON(ctx, http.MethodGet, path, q, nil, nil)
}

func (c *Client) CreateScheduleShift(ctx context.Context, scheduleID string, body CreateScheduleShiftRequest) (*Response, error) {
	path := "/v1/schedules/" + escapedPathParam(scheduleID) + "/shifts"
	return c.doJSON(ctx, http.MethodPost, path, nil, body, nil)
}

func (c *Client) MergeScheduleShifts(ctx context.Context, scheduleID string, body MergeScheduleShiftsRequest) (*Response, error) {
	path := "/v1/schedules/" + escapedPathParam(scheduleID) + "/shifts/merge"
	return c.doJSON(ctx, http.MethodPost, path, nil, body, nil)
}

func (c *Client) DeleteScheduleShift(ctx context.Context, scheduleID string, shiftID string) (*Response, error) {
	path := "/v1/schedules/" + escapedPathParam(scheduleID) + "/shifts/" + escapedPathParam(shiftID)
	return c.doJSON(ctx, http.MethodDelete, path, nil, nil, nil)
}

// Sensor events
func (c *Client) CreateSensorEvent(ctx context.Context, organizationID string, body CreateSensorEventRequest) (*Response, error) {
	headers := map[string]string{}
	if organizationID != "" {
		headers["X-Organization-Id"] = organizationID
	}
	return c.doJSON(ctx, http.MethodPost, "/v1/sensor-events", nil, body, headers)
}

// Incidents
func (c *Client) ListIncidents(ctx context.Context, params ListIncidentsParams) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/incidents", params.values(), nil, nil)
}

func (c *Client) CreateIncident(ctx context.Context, body CreateIncidentRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPost, "/v1/incidents", nil, body, nil)
}

func (c *Client) GetIncident(ctx context.Context, incidentID string) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/incidents/"+escapedPathParam(incidentID), nil, nil, nil)
}

func (c *Client) PatchIncident(ctx context.Context, incidentID string, body PatchIncidentRequest) (*Response, error) {
	path := "/v1/incidents/" + escapedPathParam(incidentID)
	return c.doJSON(ctx, http.MethodPatch, path, nil, body, nil)
}

func (c *Client) AcknowledgeIncident(ctx context.Context, incidentID string) (*Response, error) {
	path := "/v1/incidents/" + escapedPathParam(incidentID) + "/acknowledge"
	return c.doJSON(ctx, http.MethodPost, path, nil, map[string]any{}, nil)
}

func (c *Client) ResolveIncident(ctx context.Context, incidentID string, body ResolveIncidentRequest) (*Response, error) {
	path := "/v1/incidents/" + escapedPathParam(incidentID) + "/resolve"
	return c.doJSON(ctx, http.MethodPost, path, nil, body, nil)
}

func (c *Client) DeclineIncident(ctx context.Context, incidentID string, body DeclineIncidentRequest) (*Response, error) {
	path := "/v1/incidents/" + escapedPathParam(incidentID) + "/decline"
	return c.doJSON(ctx, http.MethodPost, path, nil, body, nil)
}

func (c *Client) MergeIncident(ctx context.Context, incidentID string, body MergeIncidentRequest) (*Response, error) {
	path := "/v1/incidents/" + escapedPathParam(incidentID) + "/merge"
	return c.doJSON(ctx, http.MethodPost, path, nil, body, nil)
}

func (c *Client) ListIncidentTimeline(ctx context.Context, incidentID string) (*Response, error) {
	path := "/v1/incidents/" + escapedPathParam(incidentID) + "/timeline"
	return c.doJSON(ctx, http.MethodGet, path, nil, nil, nil)
}

func (c *Client) CreateIncidentNote(ctx context.Context, incidentID string, body CreateIncidentNoteRequest) (*Response, error) {
	path := "/v1/incidents/" + escapedPathParam(incidentID) + "/notes"
	return c.doJSON(ctx, http.MethodPost, path, nil, body, nil)
}

func (c *Client) StartIncidentPaging(ctx context.Context, incidentID string, body StartIncidentPagingRequest) (*Response, error) {
	path := "/v1/incidents/" + escapedPathParam(incidentID) + "/page"
	return c.doJSON(ctx, http.MethodPost, path, nil, body, nil)
}

func (c *Client) ListIncidentEscalations(ctx context.Context, incidentID string, includeTargets *bool) (*Response, error) {
	q := url.Values{}
	addOptionalBool(q, "include_targets", includeTargets)
	path := "/v1/incidents/" + escapedPathParam(incidentID) + "/escalations"
	return c.doJSON(ctx, http.MethodGet, path, q, nil, nil)
}

func (c *Client) CancelEscalationRun(ctx context.Context, incidentID string, runID string) (*Response, error) {
	path := "/v1/incidents/" + escapedPathParam(incidentID) + "/escalations/" + escapedPathParam(runID) + "/cancel"
	return c.doJSON(ctx, http.MethodPost, path, nil, nil, nil)
}

// Suppressions
func (c *Client) ListSuppressions(ctx context.Context) (*Response, error) {
	return c.doJSON(ctx, http.MethodGet, "/v1/suppressions", nil, nil, nil)
}

func (c *Client) CreateSuppression(ctx context.Context, body CreateSuppressionRequest) (*Response, error) {
	return c.doJSON(ctx, http.MethodPost, "/v1/suppressions", nil, body, nil)
}

func (c *Client) DeleteSuppression(ctx context.Context, suppressionID string) (*Response, error) {
	return c.doJSON(ctx, http.MethodDelete, "/v1/suppressions/"+escapedPathParam(suppressionID), nil, nil, nil)
}
