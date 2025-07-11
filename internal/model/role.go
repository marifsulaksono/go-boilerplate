package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*

This file is for role model
You can adjust the structure's field as your need

*/

type (
	Role struct {
		ID   uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
		Name string    `json:"name" gorm:"unique;type:varchar(100)"`
		Model
	}

	RoleRequest struct {
		Page   int    `json:"page"`
		Limit  int    `json:"limit"`
		Search string `json:"search"`
	}
)

func (u *Role) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
