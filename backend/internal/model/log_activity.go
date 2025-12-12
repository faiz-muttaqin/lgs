package model

import (
	"time"

	"gorm.io/datatypes"
)

type LogActivity struct {
	ID uint `json:"id" gorm:"column:id;primaryKey"`

	// ===== Actor =====
	UserID    uint   `json:"user_id" gorm:"column:user_id;index"`
	IP        string `json:"ip" gorm:"column:ip;size:45"` // IPv4/IPv6
	UserAgent string `json:"user_agent" gorm:"column:user_agent;type:text"`

	// ===== Action =====
	Action     string `json:"action" gorm:"column:action;size:32;index"`           // CREATE | UPDATE | DELETE | LOGIN | APPROVE
	Resource   string `json:"resource" gorm:"column:resource;size:64;index"`       // merchant, terminal, user
	ResourceID string `json:"resource_id" gorm:"column:resource_id;size:64;index"` // "123", UUID, serial number

	// ===== Request Context =====
	ReqMethod string `json:"req_method" gorm:"column:req_method;size:8"`
	ReqURI    string `json:"req_uri" gorm:"column:req_uri;type:text"`

	// ===== Change Tracking =====
	BeforeData datatypes.JSON `json:"before_data,omitempty" gorm:"column:before_data;type:json"` // nil for CREATE
	AfterData  datatypes.JSON `json:"after_data,omitempty" gorm:"column:after_data;type:json"`   // nil for DELETE

	// ===== Result =====
	Status  string `json:"status" gorm:"column:status;size:16"` // SUCCESS | FAILED
	Message string `json:"message,omitempty" gorm:"column:message;type:text"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}

func (LogActivity) TableName() string {
	return "log_activities"
}
