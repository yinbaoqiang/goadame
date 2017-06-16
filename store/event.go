package store

import (
	"time"

	"github.com/yinbaoqiang/goadame/engine"
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

}

// 存储钩子回调事件成功
func (s eventStore) HookSuccess(url string, ei engine.Event, start, end time.Time) {

}

// 存储事件
func (s eventStore) SaveEvent(ei engine.Event) {

}
