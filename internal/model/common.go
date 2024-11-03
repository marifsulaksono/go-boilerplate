package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedBy string         `json:"created_by,omitempty" gorm:"type:varchar(36)"`
	UpdatedBy string         `json:"updated_by,omitempty" gorm:"type:varchar(36)"`
	DeletedBy string         `json:"deleted_by,omitempty" gorm:"type:varchar(36)"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
