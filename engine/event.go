package engine

import "time"

// Event 事件
type Event struct {
	Eid     string      `json:"eid" bson:"eid"`
	Action  string      `json:"action" bson:"action"`
	Etype   string      `json:"etype" bson:"etype"`
	From    string      `json:"from" bson:"from"`
	OccTime time.Time   `json:"occTime" bson:"occTime"`
	Data    interface{} `json:"data" bson:"data"`
}
