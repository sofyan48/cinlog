package entity

// LoggerRequest ...
type LoggerRequest struct {
	UUID   string                 `json:"uuid" bson:"uuid"`
	Action string                 `json:"action" bson:"action"`
	Data   map[string]interface{} `json:"data" bson:"data"`
	Status string                 `json:"status" bson:"status"`
}

// GetLoggerRequest ...
type GetLoggerRequest struct {
	UUID   string `json:"uuid" bson:"uuid"`
	Action string `json:"action" bson:"action"`
}

// GetAllLoggerRequest ...
type GetAllLoggerRequest struct {
	Action string `json:"action" bson:"action"`
}
