package team_repo

type CreateTeamRequest struct {
	Name      string `json:"name"`
	CreatedBy string
}
type MemberRequest struct {
	UserId string `json:"user_id"`
	TeamId string `json:"team_id"`
}
