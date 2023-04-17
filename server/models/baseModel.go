package models

import (
				"time"
)

type BaseModel struct {
				ID uint    `gorm:"primary_key" json:"id"`
				CreatedAt *time.Time `json:"create_at"`
				UpdatedAt *time.Time `json:"update_at"`
				DeletedAt *time.Time `json:"delete_at"`
}