package model

import (
	"time"
)

// User is gorm model
type PushLogs struct {
	ID         uint   `json:"id"`
	PushId     string `json:"push_id"`
	Name       string `json:"name"`
	BatchId    string `json:"batch_id"`
	BusinessId string `json:"business_id"`
	AppId      string `json:"app_id"`
	Total      int    `json:"total"`
	Success    int    `json:"success"`
	Receive    int    `json:"receive"`
	PushMethod int    `json:"push_method"`
	Type       int    `json:"type"`

	//TargetType
	//IsDelay
	//IsDelete
	//IsSingle
	//Status
	//ProgressRate

	Message   string    `json:"message"`
	PushedAt  time.Time `json:"pushed_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
