package domain

import "time"

type Logservice struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	App       string      `json:"app"`
	Data      interface{} `json:"data"`
	Type      string      `json:"type"`
	Status    string      `json:"status"`
	ProcessId string      `json:"process_id"`
	User      string      `json:"user"`
	CreatedAt time.Time   `json:"created_at"`
	CreatedBy string      `json:"created_by"`
	UpdatedAt time.Time   `json:"updated_at"`
	UpdatedBy string      `json:"updated_by"`
	DeletedAt time.Time   `json:"deleted_at"`
	DeletedBy string      `json:"deleted_by"`
}

func NewLogservice(app string, name string, data interface{}, processId string,
	logType string, logStatus string, user string) Logservice {
	return Logservice{
		Name:      name,
		App:       app,
		Data:      data,
		Type:      logType,
		Status:    logStatus,
		ProcessId: processId,
		CreatedAt: time.Now(),
		CreatedBy: user,
	}
}
