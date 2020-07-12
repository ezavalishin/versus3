package models

import (
	"time"
)

type BaseModel struct {
	// ID should use uuid_generate_v4() for the pk's
	ID        int        `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"index"`
}

type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `sql:"index"`
}
