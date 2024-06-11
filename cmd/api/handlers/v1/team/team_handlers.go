package team

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"team-management/cmd/api/response"
	"team-management/internal/service/team_service"
)

type TeamHandler struct {
	teamService *team_service.TeamService
}

func NewInternalHandler(teamService *team_service.TeamService) *TeamHandler {
	return &TeamHandler{teamService: teamService}
}
func (h *TeamHandler) CreateTeam(c echo.Context) error {
	var req team_service.CreateTeamRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadGateway, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	teamObj, err := h.teamService.CreateTeam(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error creating team",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "team created successfully",
		Data:       teamObj,
		Error:      nil,
	})
}
func (h *TeamHandler) AddMember(c echo.Context) error {
	var req team_service.MemberRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	teamId := c.Param("teamId")
	if teamId == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding members",
			Data:       nil,
			Error:      "invalid team id",
		})
	}
	req.TeamId = teamId
	err := h.teamService.AddMember(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding member",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "member added successfully",
		Data:       nil,
		Error:      nil,
	})
}
func (h *TeamHandler) RemoveMember(c echo.Context) error {
	memberId := c.Param("memberId")
	teamId := c.Param("teamId")
	if teamId == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding members",
			Data:       nil,
			Error:      "invalid team id",
		})
	}
	if memberId == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding members",
			Data:       nil,
			Error:      "invalid member id",
		})
	}
	req := team_service.MemberRequest{
		UserId: memberId,
		TeamId: teamId,
	}
	if err := h.teamService.RemoveMember(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error removing member",
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "member removed successfully",
	})
}
func (h *TeamHandler) MakeAdmin(c echo.Context) error {
	memberId := c.Param("memberId")
	teamId := c.Param("teamId")
	if teamId == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding members",
			Data:       nil,
			Error:      "invalid team id",
		})
	}
	if memberId == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding members",
			Data:       nil,
			Error:      "invalid member id",
		})
	}
	req := team_service.MemberRequest{
		UserId: memberId,
		TeamId: teamId,
	}
	if err := h.teamService.MakeAdmin(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding admin",
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "admin added successfully",
		Data:       nil,
		Error:      nil,
	})
}
