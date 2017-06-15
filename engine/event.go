package engine

import "time"

// Event 事件
type Event struct {
	Eid     string    `json:"eid"`
	Action  string    `json:"action"`
	Etype   string    `json:"etype"`
	From    string    `json:"from"`
	OccTime time.Time `json:"occTime"`
	Data    interface{}
}
