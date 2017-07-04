package store

import (
	"time"

	log "qiniupkg.com/x/log.v7"

	"github.com/yinbaoqiang/goadame/engine"
	"gopkg.in/mgo.v2"
)

const (
	collNameEventList        = "eventList"     // 事件发生集合
	collNameEventHookHistory = "eventHookList" // 事件或掉历史集合
)

// CreateEventStore 创建一个事件存储器
func CreateEventStore() engine.EventStorer {
	return &eventStore{}
}

// eventStore 数据存储
type eventStore struct {
}

// 存储钩子回调事件失败
func (s eventStore) HookError(url string, ei engine.Event, err error, start, end time.Time) {
	s.hookSave(url, ei, start, end, false, err)
}

// 存储钩子回调事件成功
func (s eventStore) HookSuccess(url string, ei engine.Event, start, end time.Time) {
	s.hookSave(url, ei, start, end, true, nil)
}

// storeEvent 事件
type storeHistoryHook struct {
	Eid       string      `bson:"eid"`
	HookURL   string      `bson:"hookUrl"`
	ExecOk    bool        `bson:"execOk"`
	StartTime time.Time   `bson:"startTime"`
	EndTime   time.Time   `bson:"endTime"`
	HookTime  int64       `bson:"hookTime"`
	Error     string      `bson:"error,omitempty"`
	Data      interface{} `bson:"data"`
}

// 存储钩子回调事件成功
func (s eventStore) hookSave(url string, ei engine.Event, start, end time.Time, execok bool, execerr error) {
	err := ExecSync(collNameEventHookHistory, func(coll *mgo.Collection) error {
		q := storeHistoryHook{
			Eid:       ei.Eid,
			HookURL:   url,
			ExecOk:    execok,
			StartTime: start,
			EndTime:   end,
			HookTime:  int64(end.Sub(start)),
			Data:      ei.Data,
		}
		if !execok {
			q.Error = execerr.Error()
		}
		return coll.Insert(q)
	})
	if err != nil {
		log.Errorf("HookSuccess存储失败:%v", err)
	}
}

// storeEvent 事件
type storeEvent struct {
	Eid     string      `bson:"_id"`
	Action  string      `bson:"action"`
	Etype   string      `bson:"etype"`
	From    string      `bson:"from"`
	OccTime time.Time   `bson:"occTime"`
	Data    interface{} `bson:"data"`
}

// 存储事件
func (s eventStore) SaveEvent(ei engine.Event) {
	err := ExecSync(collNameEventList, func(coll *mgo.Collection) error {
		return coll.Insert(storeEvent{
			Eid:     ei.Eid,
			Action:  ei.Action,
			Etype:   ei.Etype,
			From:    ei.From,
			OccTime: ei.OccTime,
			Data:    ei.Data,
		})
	})
	if err != nil {
		log.Errorf("SaveEvent失败:%v", err)
	}
}
