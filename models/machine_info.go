package models

import "time"

type MachineInfo struct {
	ID           uint   `gorm:"primaryKey"`
	ZbxID        string `gorm:"unique;not null;size:50;column:zbx_id"`
	SystemType   string `gorm:"not null;size:20"`
	Manufacturer string `gorm:"not null;size:20"`
	ServerSN     string `gorm:"not null;size:50"`
	SystemDisk   string `gorm:"not null;size:20"`
	SSDCount     string `gorm:"size:20"`
	HDDCount     string `gorm:"size:20"`
	MemoryCount  string `gorm:"size:20"`
	CPUInfo      string `gorm:"type:text;not null"`
	ServerHeight string `gorm:"not null;size:10"`
	CreatedAt    time.Time
}

func (MachineInfo) TableName() string { return "machine_info" }
