package team_service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"team-management/common/echo_ctx"
	"team-management/internal/models/primary/team"
	"team-management/internal/repo/team_repo"
)

type TeamService struct {
	teamRepo *team_repo.TeamRepo
}

func NewTeamService(teamRepo *team_repo.TeamRepo) *TeamService {
	return &TeamService{
		teamRepo: teamRepo,
	}
}
func (t *TeamService) CreateTeam(ctx echo.Context, req CreateTeamRequest) (*team.Team, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("invalid name")
	}
	createTeamRequest := team_repo.CreateTeamRequest{
		Name:      req.Name,
		CreatedBy: echo_ctx.GetUserId(ctx),
	}
	teamObj, err := t.teamRepo.CreateTeam(ctx, createTeamRequest)
	if err != nil {
		return nil, err
	}
	return teamObj, nil
}
func (t *TeamService) AddMember(ctx echo.Context, req MemberRequest) error {
	if req.UserId == "" {
		return fmt.Errorf("invalid name")
	}
	request := team_repo.MemberRequest{
		UserId: req.UserId,
		TeamId: req.TeamId,
	}
	if err := t.teamRepo.AddMember(ctx, request); err != nil {
		return err
	}
	return nil
}
func (t *TeamService) RemoveMember(ctx echo.Context, req MemberRequest) error {
	request := team_repo.MemberRequest{
		UserId: req.UserId,
		TeamId: req.TeamId,
	}
	memberExists, err := t.teamRepo.MemberExists(ctx, request)
	if err != nil {
		return err
	}
	if !memberExists {
		return fmt.Errorf("member does not exist")
	}
	if err := t.teamRepo.RemoveMember(ctx, request); err != nil {
		return err
	}
	return nil
}
func (t *TeamService) MakeAdmin(ctx echo.Context, req MemberRequest) error {
	request := team_repo.MemberRequest{
		UserId: req.UserId,
		TeamId: req.TeamId,
	}
	memberExists, err := t.teamRepo.MemberExists(ctx, request)
	if err != nil {
		return err
	}
	if !memberExists {
		return fmt.Errorf("member does not exist")
	}
	if err := t.teamRepo.MakeAdmin(ctx, request); err != nil {
		return err
	}
	return nil
}
