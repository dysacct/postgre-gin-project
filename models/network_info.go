package models

import "time"

type NetworkInfo struct {
	ID           uint   `gorm:"primaryKey"`
	ZbxID        string `gorm:"unique;not null;size:50;column:zbx_id"`
	MacAddress   string `gorm:"not null;size:17"`
	EthName      string `gorm:"not null;size:15"`
	IDCCode      string `gorm:"not null;size:10"`
	NetType      string `gorm:"not null;size:20"`
	Vlan         string `gorm:"not null;size:9"`
	IPv4IP       string `gorm:"not null;size:20"`
	IPv4Gateway  string `gorm:"not null;size:20"`
	IPv6IP       string `gorm:"size:50"`
	IPv6Gateway  string `gorm:"size:50"`
	IPSpeed      int16  `gorm:"not null"`
	IPStatus     string `gorm:"size:10"`
	IPNotes      string `gorm:"size:255"`
	SegmentNotes string `gorm:"size:255"`
	CreatedAt    time.Time
}

func (NetworkInfo) TableName() string { return "network_info" }
