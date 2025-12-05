package models

import "time"

type MachineInfo struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ZbxID        string    `gorm:"column:zbx_id;unique;not null;size:50" json:"zbx_id"`
	SystemType   string    `gorm:"column:system_type;not null;size:20" json:"system_type"`
	Manufacturer string    `gorm:"column:manufacturer;not null;size:20" json:"manufacturer"`
	ServerSN     string    `gorm:"column:server_sn;not null;size:50" json:"server_sn"`
	SystemDisk   string    `gorm:"column:system_disk;not null;size:20" json:"system_disk"`
	SSDCount     string    `gorm:"column:ssd_count;not null;size:20" json:"ssd_count"`
	HDDCount     string    `gorm:"column:hdd_count;not null;size:20" json:"hdd_count"`
	MemoryCount  string    `gorm:"column:memory_count;not null;size:20" json:"memory_count"`
	CPUInfo      string    `gorm:"column:cpu_info;type:text;not null" json:"cpu_info"`
	ServerHeight string    `gorm:"column:server_height;not null;size:10" json:"server_height"`
	CreatedAt    time.Time `json:"created_at"`
}

func (MachineInfo) TableName() string { return "machine_info" }
