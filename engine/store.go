package engine

import (
	"fmt"
	"time"

	"github.com/yinbaoqiang/goadame/app"
)

// ChgType 变化类型
type ChgType int

const (
	// ChgTypeAdd 监听变化-新增
	ChgTypeAdd ChgType = iota
	// ChgTypeRemove 监听变化-移除
	ChgTypeRemove
	// ChgTypeUpdate 监听变化-修改
	ChgTypeUpdate
)

// ListenerStore 监听事件改变
type ListenerStore interface {
	Watch(func(ctyp ChgType, lis app.AntListen))
	// 获取所有的
	All() (res []*app.AntListen, err error)
}

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
}

var defaultEventStore = &showStore{}

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
func (s *showStore) Watch(func(ctyp ChgType, lis app.AntListen)) {
	fmt.Println("观察是否有新的监听变化")
}

// 获取所有的
func (s *showStore) All() (res []*app.AntListen, err error) {
	fmt.Println("加载所有监听")
	return nil, nil
}
