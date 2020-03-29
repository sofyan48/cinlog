package entity

import "time"

// LoggerEventHistory ...
type LoggerEventHistory struct {
	UUID      string          `json:"__uuid" bson:"__uuid"`
	Action    string          `json:"__action" bson:"__action"`
	Offset    uint64          `json:"__offset" bson:"__offset"`
	History   []LoggerHistory `json:"history" bson:"history"`
	Status    string          `json:"status" bson:"status"`
	CreatedAt *time.Time      `json:"created_at" bson:"created_at"`
	UpdateAt  *time.Time      `json:"update_at" bson:"update_at"`
}

// LoggerHistory ...
type LoggerHistory struct {
	Data map[string]string `json:"data" bson:"data"`
}

// LoggerSearch ...
type LoggerSearch struct {
	UUID string `json:"__uuid" bson:"__uuid"`
}
