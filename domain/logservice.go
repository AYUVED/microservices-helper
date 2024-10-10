package domain

import "time"

type Logservice struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	App       string    `json:"app"`
	Data      string    `json:"data"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	ProcessID string    `json:"process_id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy string    `json:"deleted_by"`
}

func NewLogservice(app string, name string, data string, processId string,
	logType string, logStatus string, user string) Logservice {
	return Logservice{
		Name:      name,
		App:       app,
		Data:      data,
		Type:      logType,
		Status:    logStatus,
		ProcessID: processId,
		CreatedAt: time.Now(),
		CreatedBy: user,
	}
}
