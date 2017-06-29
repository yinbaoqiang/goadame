package store

import (
	"github.com/yinbaoqiang/goadame/app"
)

// Listener 监听存储
type Listener interface {
	// Add 新增监听
	Add(app.AntListen) error
	// Update 修改监听
	Update(app.AntListen) error
	// List 查询列表
	List(action, etype string, page, cnt int, total *int) (res []*app.AntListen, err error)
	Rmove(rid string) error
}

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

// ChgListener 监听事件改变
type ChgListener interface {
	// 监控监听变化
	// ctyp 变化类型
	// lis 变化的监听数据
	Watch(func(ctyp ChgType, lis app.AntListen))
}

// ListenStore 监听存储
var defaultListenStore Listener

// SetDefaultListener 设置默认存储
func SetDefaultListener(ls Listener) {
	defaultListenStore = ls
}

// ListenStore 获取存储
func ListenStore() Listener {
	if defaultListenStore == nil {
		panic("没有初始话默认存储defaultListenStore,请使用 SetDefaultListener 进行初始化")
	}
	return defaultListenStore
}

// ChgListener 监听存储
var defaultChgListener ChgListener

// SetDefaultChgListener 设置默认存储
func SetDefaultChgListener(ls ChgListener) {
	defaultChgListener = ls
}

// ChgListen 获取存储
func ChgListen() ChgListener {
	if defaultChgListener == nil {
		panic("没有初始话默认存储 defaultChgListener ,请使用 SetDefaultChgListener 进行初始化")
	}
	return defaultChgListener
}
