package team_service

type CreateTeamRequest struct {
	Name string `json:"name" validate:"required"`
}
type MemberRequest struct {
	UserId string `json:"user_id" validate:"required"`
	TeamId string `json:"team_id"`
}
