package models

import "time"

type BusinessInfo struct {
	ID               uint   `gorm:"primaryKey"`
	ZbxID            string `gorm:"unique;not null;size:50;column:zbx_id"`
	BusinessName     string `gorm:"not null;size:100"`
	BusinessID       string `gorm:"not null;size:50"`
	OldBusinessName  string `gorm:"not null;size:100"`
	OldBusinessID    string `gorm:"not null;size:50"`
	BusinessSpeed    int16  `gorm:"not null"`
	OldBusinessSpeed int16  `gorm:"not null"`
	CreatedAt        time.Time
}

func (BusinessInfo) TableName() string { return "business_info" }
