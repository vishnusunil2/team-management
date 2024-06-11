package team

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"reflect"
	"team-management/cmd/api/response"
	"testing"
)

const (
	AuthToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjZiYWVkNWItZDE3Ny00MTM2LThlYzktZTllODkyZTM0NTQ0IiwibmFtZSI6InZpc2hudSIsImVtYWlsIjoidmlzaG51c3VuaWwyNDNAZ21haWwuY29tIiwiZXhwIjoxNzE4MzUxNjQyLCJpYXQiOjE3MTgwOTI0NDIsInN1YiI6IjI2YmFlZDViLWQxNzctNDEzNi04ZWM5LWU5ZTg5MmUzNDU0NCJ9.6s_YxIVNnCI51myOM5EaKY3gk6hLslug4FM3jarEFc4"
)

type CreateTeamRequest struct {
	Name string `json:"name"`
}

func NewCreateTeamRequest(name string) CreateTeamRequest {
	return CreateTeamRequest{Name: name}
}

type CreateTeamResponse struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}

func NewCreateTeamResponse(name string, createdBy string) *CreateTeamResponse {
	return &CreateTeamResponse{
		Name:      name,
		CreatedBy: createdBy,
	}
}
func NewResponseWithoutData(statusCode int, message string, error interface{}) response.Response {
	return response.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
		Error:      error,
	}
}
func getRandName() string {
	length := 10
	str := uuid.New().String()
	return str[:length]
}

func TestCreateTeam(t *testing.T) {
	teamName := getRandName()
	method := "POST"
	url := "http://localhost:8080/api/v1/teams/"
	testCases := []struct {
		Name               string
		RequestBody        CreateTeamRequest
		ExpectedStatusCode int
		ExpectedBody       *CreateTeamResponse
	}{
		{
			Name:               "Success",
			RequestBody:        NewCreateTeamRequest(teamName),
			ExpectedStatusCode: 200,
			ExpectedBody:       NewCreateTeamResponse(teamName, "26baed5b-d177-4136-8ec9-e9e892e34544"),
		},
		{
			Name:               "Failiure",
			RequestBody:        NewCreateTeamRequest(""),
			ExpectedStatusCode: 400,
			ExpectedBody:       NewCreateTeamResponse("", ""),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			payloadBytes, _ := json.Marshal(tt.RequestBody)
			client := &http.Client{}
			req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", AuthToken)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("failed to execute request %v", err)
			}
			defer func(body io.ReadCloser) {
				err := body.Close()
				if err != nil {
					t.Fatalf("failed to close body: %v", err)
				}
			}(res.Body)
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Failed to read response %v", err)
			}
			if res.StatusCode != tt.ExpectedStatusCode {
				t.Errorf("expected %v , got %v ", tt.ExpectedStatusCode, res.StatusCode)
			}
			var apiResponse response.Response
			err = json.Unmarshal(body, &apiResponse)
			if err != nil {
				t.Fatalf("failed to unmarshal response %v", err)
			}
			if apiResponse.Data != nil {
				if apiResponse.Data.(map[string]interface{})["name"] != tt.ExpectedBody.Name {
					t.Fatalf("expected %v got %v", apiResponse.Data.(map[string]interface{})["name"], tt.ExpectedBody.Name)
				}
				if apiResponse.Data.(map[string]interface{})["created_by"] != tt.ExpectedBody.CreatedBy {
					t.Fatalf("expected %v got %v", apiResponse.Data.(map[string]interface{})["created_by"], tt.ExpectedBody.CreatedBy)
				}
			}
		})
	}
}

type MemberRequest struct {
	UserId string `json:"user_id"`
}

func NewMemberRequest(userId string) *MemberRequest {
	return &MemberRequest{
		UserId: userId,
	}
}
func TestAddMember(t *testing.T) {
	userId := "5fa98ec2-9b27-41c9-93e7-089a4fb0ed23"
	method := "POST"
	url := "http://localhost:8080/api/v1/teams/2db4c285-d7c1-4c11-aa5f-57dea6491eb0/members"
	testCases := []struct {
		Name               string
		RequestBody        *MemberRequest
		ExpectedStatusCode int
		ExpectedBody       response.Response
	}{
		{
			Name:               "Success",
			RequestBody:        NewMemberRequest(userId),
			ExpectedStatusCode: 200,
			ExpectedBody:       NewResponseWithoutData(200, "member added successfully", nil),
		},
		{
			Name:               "Fail",
			RequestBody:        nil,
			ExpectedStatusCode: 400,
			ExpectedBody:       NewResponseWithoutData(400, "error adding member", "invalid name"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			payloadBytes, _ := json.Marshal(tt.RequestBody)
			client := &http.Client{}
			req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", AuthToken)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("failed to execute request %v", err)
			}
			defer func(body io.ReadCloser) {
				err := body.Close()
				if err != nil {
					t.Fatalf("failed to close body: %v", err)
				}
			}(res.Body)
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Failed to read response %v", err)
			}
			if res.StatusCode != tt.ExpectedStatusCode {
				t.Fatalf("expected %v got %v ", tt.ExpectedStatusCode, res.StatusCode)
			}
			var apiResponse response.Response
			err = json.Unmarshal(body, &apiResponse)
			if err != nil {
				t.Fatalf("failed to unmarshal response %v", err)
			}
			if !reflect.DeepEqual(apiResponse, tt.ExpectedBody) {
				t.Fatalf("expected %v got %v", tt.ExpectedBody, apiResponse)
			}
		})
	}

}
func TestRemoveMember(t *testing.T) {
	method := "DELETE"
	testCases := []struct {
		Name               string
		URL                string
		ExpectedStatusCode int
		ExpectedBody       response.Response
	}{
		{
			Name:               "Success",
			URL:                "http://localhost:8080/api/v1/teams/2db4c285-d7c1-4c11-aa5f-57dea6491eb0/members/5fa98ec2-9b27-41c9-93e7-089a4fb0ed23",
			ExpectedStatusCode: 200,
			ExpectedBody:       NewResponseWithoutData(200, "member removed successfully", nil),
		},
		{
			Name:               "Fail",
			URL:                "http://localhost:8080/api/v1/teams/2db4c285-d7c1-4c11-aa5f-57dea6491eb0/members/abcde",
			ExpectedStatusCode: 400,
			ExpectedBody:       NewResponseWithoutData(400, "error removing member", "member does not exist"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest(method, tt.URL, nil)
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", AuthToken)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("failed to execute request %v", err)
			}
			defer func(body io.ReadCloser) {
				err := body.Close()
				if err != nil {
					t.Fatalf("failed to close body: %v", err)
				}
			}(res.Body)
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Failed to read response %v", err)
			}
			if res.StatusCode != tt.ExpectedStatusCode {
				t.Fatalf("expected %v got %v ", tt.ExpectedStatusCode, res.StatusCode)
			}
			var apiResponse response.Response
			err = json.Unmarshal(body, &apiResponse)
			if err != nil {
				t.Fatalf("failed to unmarshal response %v", err)
			}
			if !reflect.DeepEqual(apiResponse, tt.ExpectedBody) {
				t.Fatalf("expected %v got %v", tt.ExpectedBody, apiResponse)
			}
		})
	}
}
func TestMakeAdmin(t *testing.T) {
	method := "PATCH"
	testCases := []struct {
		Name               string
		URL                string
		ExpectedStatusCode int
		ExpectedBody       response.Response
	}{
		{
			Name:               "Success",
			ExpectedStatusCode: 200,
			URL:                "http://localhost:8080/api/v1/teams/2db4c285-d7c1-4c11-aa5f-57dea6491eb0/members/5fa98ec2-9b27-41c9-93e7-089a4fb0ed23",
			ExpectedBody:       NewResponseWithoutData(200, "admin added successfully", nil),
		},
		{
			Name:               "Fail",
			ExpectedStatusCode: 400,
			URL:                "http://localhost:8080/api/v1/teams/2db4c285-d7c1-4c11-aa5f-57dea6491eb0/members/abcde",
			ExpectedBody:       NewResponseWithoutData(400, "error adding admin", "member does not exist"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest(method, tt.URL, nil)
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", AuthToken)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("failed to execute request %v", err)
			}
			defer func(body io.ReadCloser) {
				err := body.Close()
				if err != nil {
					t.Fatalf("failed to close body: %v", err)
				}
			}(res.Body)
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Failed to read response %v", err)
			}
			if res.StatusCode != tt.ExpectedStatusCode {
				t.Fatalf("expected %v got %v ", tt.ExpectedStatusCode, res.StatusCode)
			}
			var apiResponse response.Response
			err = json.Unmarshal(body, &apiResponse)
			if err != nil {
				t.Fatalf("failed to unmarshal response %v", err)
			}
			if !reflect.DeepEqual(apiResponse, tt.ExpectedBody) {
				t.Fatalf("expected %v got %v", tt.ExpectedBody, apiResponse)
			}
		})
	}
}
