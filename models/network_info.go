package models

import "time"

type NetworkInfo struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ZbxID        string    `gorm:"column:zbx_id;unique;not null;size:50" json:"zbx_id"`
	MacAddress   string    `gorm:"column:mac_address;not null;size:17" json:"mac_address"`
	EthName      string    `gorm:"column:eth_name;not null;size:15" json:"eth_name"`
	IDCCode      string    `gorm:"column:idc_code;not null;size:10" json:"idc_code"`
	NetType      string    `gorm:"column:net_type;not null;size:20" json:"net_type"`
	Vlan         string    `gorm:"column:vlan;not null;size:9" json:"vlan"`
	IPv4IP       string    `gorm:"column:ipv4_ip;not null;size:20" json:"ipv4_ip"`
	IPv4Gateway  string    `gorm:"column:ipv4_gateway;not null;size:20" json:"ipv4_gateway"`
	IPv6IP       string    `gorm:"column:ipv6_ip;size:50" json:"ipv6_ip"`
	IPv6Gateway  string    `gorm:"column:ipv6_gateway;size:50" json:"ipv6_gateway"`
	IPSpeed      int16     `gorm:"column:ip_speed;not null" json:"ip_speed"`
	IPStatus     string    `gorm:"column:ip_status;size:10" json:"ip_status"`
	IPNotes      string    `gorm:"column:ip_notes;size:255" json:"ip_notes"`
	SegmentNotes string    `gorm:"column:segment_notes;size:255" json:"segment_notes"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (NetworkInfo) TableName() string { return "network_info" }
