package models

import "time"

type IDCInfo struct {
	ID        uint      `gorm:"primaryKey"`
	ZbxID     string    `gorm:"unique;not null;size:50;column:zbx_id"`
	IDCCode   string    `gorm:"not null;size:10"`
	IDCName   string    `gorm:"not null;size:50"`
	IPMIIP    string    `gorm:"not null;size:16"`
	SSHIP     string    `gorm:"not null;size:16"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (IDCInfo) TableName() string {
	return "idc_info"
}
