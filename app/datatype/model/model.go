package model

import (
	"database/sql"
	"time"
)

type DefaultModel struct {
	ID        uint         `gorm:"primaryKey"`
	CreatedAt time.Time    `gorm:"default:null"`
	UpdatedAt time.Time    `gorm:"default:null"`
	DeletedAt sql.NullTime `gorm:"index;default:null"`
}

type DefaultModelWithoutDlt struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:null"`
	UpdatedAt time.Time `gorm:"default:null"`
}
