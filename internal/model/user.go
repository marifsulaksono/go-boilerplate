package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*

This file is for user model
You can adjust the structure's field as your need

*/

type (
	User struct {
		ID       uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
		Name     string    `json:"name" gorm:"type:varchar(100)"`
		Email    string    `json:"email" gorm:"unique;not null;type:varchar(300)"`
		Password string    `json:"-"`
		RoleID   uuid.UUID `json:"role_id" gorm:"not null;type:varchar(36)"`
		Role     Role      `json:"role" gorm:"foreignKey:RoleID;references:ID"`
		Model
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
