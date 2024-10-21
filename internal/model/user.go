package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		ID       uuid.UUID `json:"id" gorm:"type:char(36)"`
		Name     string    `json:"name" gorm:"type:char(100)"`
		Email    string    `json:"email" gorm:"type:char(255)"`
		Password string    `json:"-"`
		Model
	}

	UserResponse struct {
		ID    uuid.UUID `json:"id" gorm:"type:char(36)"`
		Name  string    `json:"name" gorm:"type:char(100)"`
		Email string    `json:"email" gorm:"type:char(255)"`
		Model
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
