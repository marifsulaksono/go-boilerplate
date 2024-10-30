package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		ID       uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
		Name     string    `json:"name" gorm:"type:char(100)"`
		Email    string    `json:"email" gorm:"type:char(255)"`
		Password string    `json:"-"`
		Model
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
