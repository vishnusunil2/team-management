package team

import (
	"github.com/google/uuid"
	"team-management/internal/models/primary/user"
	"time"
)

const (
	OrgIdMaxLength = 6
)

type Team struct {
	Id        uuid.UUID `gorm:"primaryKey;index" json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
type TeamMember struct {
	TeamId  string `gorm:"primaryKey"`
	UserId  string `gorm:"primaryKey"`
	IsAdmin bool
	Team    Team      `gorm:"foreignKey:TeamId"`
	User    user.User `gorm:"foreignKey:UserId"`
}

func NewTeam(name string, createdBy string) *Team {
	return &Team{
		Id:        uuid.New(),
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}
}
func NewTeamMember(userId string, teamId string) *TeamMember {
	return &TeamMember{
		TeamId:  teamId,
		UserId:  userId,
		IsAdmin: false,
	}
}
