package engine

import (
	"fmt"
	"time"
)

// HookInfo 钩子信息
type HookInfo struct {
	ID, Etype, Action, URL string
}

// EventStorer 数据存储
type EventStorer interface {
	// 存储钩子回调事件失败
	HookError(url string, ei Event, err error, start, end time.Time)
	// 存储钩子回调事件成功
	HookSuccess(url string, ei Event, start, end time.Time)
	// 存储事件
	SaveEvent(ei Event)
	// // 存储钩子监听
	// SaveHook(etype, action, url string)
	// // 加载所以的监听钩子
	// LoadAllHooks() []HookInfo
}

var defaultStore EventStorer

// RegEventStorer 注册store驱动
func RegEventStorer(store EventStorer) {
	defaultStore = store
}

// showStore 数据存储
type showStore struct {
}

// 存储钩子回调事件失败
func (s *showStore) HookError(url string, ei Event, err error, start, end time.Time) {
	fmt.Printf("(%v-%v)(%v) %s:%s->%s error:%v\n", start, end, end.Sub(start), ei.Etype, ei.Action, url, err)
}

// 存储钩子回调事件成功
func (s *showStore) HookSuccess(url string, ei Event, start, end time.Time) {

	fmt.Printf("(%v-%v)(%v) %s:%s->%s success.\n", start, end, end.Sub(start), ei.Etype, ei.Action, url)

}
func (s *showStore) SaveEvent(ei Event) {
	fmt.Printf("%v %s:%s->store.\n", ei.OccTime, ei.Etype, ei.Action)
}
