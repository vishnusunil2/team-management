package team_repo

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"team-management/internal/models/primary/team"
)

type TeamRepo struct {
	DB *gorm.DB
}

func NewTeamRepo(db *gorm.DB) *TeamRepo {
	return &TeamRepo{DB: db}
}
func (t *TeamRepo) CreateTeam(ctx echo.Context, req CreateTeamRequest) (*team.Team, error) {
	teamObj := team.NewTeam(req.Name, req.CreatedBy)
	err := t.getBaseDbQuery(ctx).Create(teamObj).Error
	if err != nil {
		return nil, err
	}
	return teamObj, nil
}
func (t *TeamRepo) AddMember(ctx echo.Context, req MemberRequest) error {
	teamMemberObj := team.NewTeamMember(req.UserId, req.TeamId)
	if err := t.getBaseTeamMemberDbQuery(ctx).Create(teamMemberObj).Error; err != nil {
		return err
	}
	return nil
}
func (t *TeamRepo) RemoveMember(ctx echo.Context, req MemberRequest) error {
	if err := t.getBaseTeamMemberDbQuery(ctx).
		Where("team_id=? AND user_id=?", req.TeamId, req.UserId).
		Delete(&team.TeamMember{}).Error; err != nil {
		return err
	}
	return nil
}
func (t *TeamRepo) MakeAdmin(ctx echo.Context, req MemberRequest) error {
	if err := t.getBaseTeamMemberDbQuery(ctx).
		Where("team_id=? AND user_id=?", req.TeamId, req.UserId).
		Update("is_admin", true).Error; err != nil {
		return err
	}
	return nil
}
func (t *TeamRepo) MemberExists(ctx echo.Context, req MemberRequest) (bool, error) {
	var count int64
	if err := t.getBaseTeamMemberDbQuery(ctx).
		Where("team_id=? AND user_id=?", req.TeamId, req.UserId).
		Count(&count).Error; err != nil {
		return true, err
	}
	return count > 0, nil

}
func (t *TeamRepo) getBaseDbQuery(ctx echo.Context) *gorm.DB {
	return t.DB.Model(&team.Team{}).
		WithContext(ctx.Request().Context())
}
func (t *TeamRepo) getBaseTeamMemberDbQuery(ctx echo.Context) *gorm.DB {
	return t.DB.Model(&team.TeamMember{}).WithContext(ctx.Request().Context())
}
