package engine

import "time"

// EventInfo 事件
type EventInfo struct {
	Eid     string    `json:"eid"`
	Action  string    `json:"action"`
	Etype   string    `json:"etype"`
	From    string    `json:"from"`
	OccTime time.Time `json:"occTime"`
}

// Event 事件
type Event struct {
	// GetEventInfo 获取事件信息
	Info EventInfo
	// GetData 获取数据
	Data map[string]interface{}
}
