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
	List(action, etype string, page, cnt int, total *int, res *[]*app.AntListen) error
	Rmove(rid string) error
}

// ListenrStore 监听存储
var ListenStore Listener
