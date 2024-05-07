// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package lineargql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// GetTeamByIdResponse is returned by GetTeamById on success.
type GetTeamByIdResponse struct {
	// One specific team.
	Team GetTeamByIdTeam `json:"team"`
}

// GetTeam returns GetTeamByIdResponse.Team, and is useful for accessing the field via an interface.
func (v *GetTeamByIdResponse) GetTeam() GetTeamByIdTeam { return v.Team }

// GetTeamByIdTeam includes the requested fields of the GraphQL type Team.
// The GraphQL type's documentation follows.
//
// An organizational unit that contains issues.
type GetTeamByIdTeam struct {
	// The unique identifier of the entity.
	Id string `json:"id"`
	// The team's name.
	Name string `json:"name"`
	// The team's unique key. The key is used in URLs.
	Key string `json:"key"`
}

// GetId returns GetTeamByIdTeam.Id, and is useful for accessing the field via an interface.
func (v *GetTeamByIdTeam) GetId() string { return v.Id }

// GetName returns GetTeamByIdTeam.Name, and is useful for accessing the field via an interface.
func (v *GetTeamByIdTeam) GetName() string { return v.Name }

// GetKey returns GetTeamByIdTeam.Key, and is useful for accessing the field via an interface.
func (v *GetTeamByIdTeam) GetKey() string { return v.Key }

// MoveIssueToTeamIssueUpdateIssuePayload includes the requested fields of the GraphQL type IssuePayload.
type MoveIssueToTeamIssueUpdateIssuePayload struct {
	// The identifier of the last sync operation.
	LastSyncId float64 `json:"lastSyncId"`
}

// GetLastSyncId returns MoveIssueToTeamIssueUpdateIssuePayload.LastSyncId, and is useful for accessing the field via an interface.
func (v *MoveIssueToTeamIssueUpdateIssuePayload) GetLastSyncId() float64 { return v.LastSyncId }

// MoveIssueToTeamResponse is returned by MoveIssueToTeam on success.
type MoveIssueToTeamResponse struct {
	// Updates an issue.
	IssueUpdate MoveIssueToTeamIssueUpdateIssuePayload `json:"issueUpdate"`
}

// GetIssueUpdate returns MoveIssueToTeamResponse.IssueUpdate, and is useful for accessing the field via an interface.
func (v *MoveIssueToTeamResponse) GetIssueUpdate() MoveIssueToTeamIssueUpdateIssuePayload {
	return v.IssueUpdate
}

// __GetTeamByIdInput is used internally by genqlient
type __GetTeamByIdInput struct {
	Id string `json:"id"`
}

// GetId returns __GetTeamByIdInput.Id, and is useful for accessing the field via an interface.
func (v *__GetTeamByIdInput) GetId() string { return v.Id }

// __MoveIssueToTeamInput is used internally by genqlient
type __MoveIssueToTeamInput struct {
	IssueId string `json:"issueId"`
	TeamId  string `json:"teamId"`
}

// GetIssueId returns __MoveIssueToTeamInput.IssueId, and is useful for accessing the field via an interface.
func (v *__MoveIssueToTeamInput) GetIssueId() string { return v.IssueId }

// GetTeamId returns __MoveIssueToTeamInput.TeamId, and is useful for accessing the field via an interface.
func (v *__MoveIssueToTeamInput) GetTeamId() string { return v.TeamId }

// GetTeamById returns a team by its identifier or UUID
func GetTeamById(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*GetTeamByIdResponse, error) {
	req := &graphql.Request{
		OpName: "GetTeamById",
		Query: `
query GetTeamById ($id: String!) {
	team(id: $id) {
		id
		name
		key
	}
}
`,
		Variables: &__GetTeamByIdInput{
			Id: id,
		},
	}
	var err error

	var data GetTeamByIdResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// MoveIssueToTeam moves issue between teams
// - issueId: UUID or identifier of the issue to move
// - teamId: UUID of the team to move the issue to
func MoveIssueToTeam(
	ctx context.Context,
	client graphql.Client,
	issueId string,
	teamId string,
) (*MoveIssueToTeamResponse, error) {
	req := &graphql.Request{
		OpName: "MoveIssueToTeam",
		Query: `
mutation MoveIssueToTeam ($issueId: String!, $teamId: String!) {
	issueUpdate(id: $issueId, input: {teamId:$teamId}) {
		lastSyncId
	}
}
`,
		Variables: &__MoveIssueToTeamInput{
			IssueId: issueId,
			TeamId:  teamId,
		},
	}
	var err error

	var data MoveIssueToTeamResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
