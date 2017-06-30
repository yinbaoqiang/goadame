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
	List(etype, action string, lid string, cnt int) (total int, res []*app.AntListen, err error)
	Rmove(rid string) error
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
