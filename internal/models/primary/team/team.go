package team

import (
	"gorm.io/gorm"
	"team-management/common/utils"
	"team-management/internal/models/base"
)

const (
	OrgIdMaxLength = 6
)

type Team struct {
	Id     string `gorm:"primaryKey;index"`
	Name   string
	UserId string `gorm:"null"`
	*base.AuditFields
}

func (t *Team) BeforeCreate(tx *gorm.DB) (err error) {
	if t.Id == "" {
		for {
			id, err := utils.GenerateRandomString(utils.AlphaNumeric, OrgIdMaxLength)
			if err != nil {
				return err
			}
			var exists bool
			err = tx.Model(&Team{}).
				Select("1").
				Where("id=?", id).
				Limit(1).
				Scan(&exists).Error
			if err != nil {
				return err
			}
			if !exists {
				t.Id = id
				break
			}
		}
	}
	return nil
}
