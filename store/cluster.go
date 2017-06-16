package store

import "github.com/yinbaoqiang/goadame/app"

const (
	// ChgTypAdd 事件变化类型 新增
	ChgTypAdd ListenChgTyp = iota + 1
	// ChgTypUpdate 事件变化类型 修改
	ChgTypUpdate
	// ChgTypDelete 事件变化类型 删除
	ChgTypDelete
)

// ListenChgTyp 监听变化类型
type ListenChgTyp int

// ListenChgEvent 监听变化事件
type ListenChgEvent struct {
	Type ListenChgTyp // 类型
	ID   string       // 唯一标识
	app.AntListen
}
type AntListen struct {
	// 事件行为,不设置该项则注册监听所有行为变化
	Action *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// 事件类型
	Etype string `form:"etype" json:"etype" xml:"etype"`
	// 产生事件的服务器标识
	From string `form:"from" json:"from" xml:"from"`
	// 钩子url
	Hookurl string `form:"hookurl" json:"hookurl" xml:"hookurl"`
	// 注册事件监听唯一标识
	Rid string `form:"rid" json:"rid" xml:"rid"`
}
type ClusterStorer interface {
	RegListen(app.AntListen)
	UnRegListen(app.AntListen)
	WatchRegListen() chan ListenChgEvent
}
