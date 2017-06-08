// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "antevent": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/yinbaoqiang/goadame/design
// --out=$(GOPATH)/src/github.com/yinbaoqiang/goadame
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// 处理失败 (default view)
//
// Identifier: vnd.ant.error+json; view=default
type AntError struct {
	// 错误描述
	Msg *string `form:"msg,omitempty" json:"msg,omitempty" xml:"msg,omitempty"`
}

// 事件监听信息 (default view)
//
// Identifier: vnd.ant.even.back+json; view=default
type AntEvenBack struct {
	// 回调连接
	Backurl string `form:"backurl" json:"backurl" xml:"backurl"`
	// 事件唯一标识
	Eid string `form:"eid" json:"eid" xml:"eid"`
	// 执行结束时间
	EndTime time.Time `form:"endTime" json:"endTime" xml:"endTime"`
	// 执行错误时的错误信息
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
	// 执行时间,单位纳秒
	ExecTime int `form:"execTime" json:"execTime" xml:"execTime"`
	// 开始执行时间
	StartTime time.Time `form:"startTime" json:"startTime" xml:"startTime"`
	// 执行是否成功
	Success bool `form:"success" json:"success" xml:"success"`
}

// Validate validates the AntEvenBack media type instance.
func (mt *AntEvenBack) Validate() (err error) {
	if mt.Eid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eid"))
	}
	if mt.Backurl == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "backurl"))
	}

	return
}

// AntEvenBackCollection is the media type for an array of AntEvenBack (default view)
//
// Identifier: vnd.ant.even.back+json; type=collection; view=default
type AntEvenBackCollection []*AntEvenBack

// Validate validates the AntEvenBackCollection media type instance.
func (mt AntEvenBackCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// 事件监听列表 (default view)
//
// Identifier: vnd.ant.event.history.list+json; view=default
type AntEventHistoryList struct {
	// 事件类型
	List []*AntHistoryInfo `form:"list,omitempty" json:"list,omitempty" xml:"list,omitempty"`
	// 总数量
	Total int `form:"total" json:"total" xml:"total"`
}

// Validate validates the AntEventHistoryList media type instance.
func (mt *AntEventHistoryList) Validate() (err error) {
	for _, e := range mt.List {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// 事件监听信息 (default view)
//
// Identifier: vnd.ant.history.info+json; view=default
type AntHistoryInfo struct {
	// 事件行为,不设置该项则注册监听所有行为变化
	Action *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// 事件唯一标识
	Eid string `form:"eid" json:"eid" xml:"eid"`
	// 事件类型
	Etype string `form:"etype" json:"etype" xml:"etype"`
	// 产生事件的服务器标识
	From string `form:"from" json:"from" xml:"from"`
}

// Validate validates the AntHistoryInfo media type instance.
func (mt *AntHistoryInfo) Validate() (err error) {
	if mt.Eid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eid"))
	}
	if mt.Etype == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "etype"))
	}
	if mt.From == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "from"))
	}
	return
}

// 事件监听信息 (default view)
//
// Identifier: vnd.ant.reg+json; view=default
type AntReg struct {
	// 事件行为,不设置该项则注册监听所有行为变化
	Action *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// 回调路径
	Backurl string `form:"backurl" json:"backurl" xml:"backurl"`
	// 事件类型
	Etype string `form:"etype" json:"etype" xml:"etype"`
	// 产生事件的服务器标识
	From string `form:"from" json:"from" xml:"from"`
	// 注册事件监听唯一标识
	Rid string `form:"rid" json:"rid" xml:"rid"`
}

// Validate validates the AntReg media type instance.
func (mt *AntReg) Validate() (err error) {
	if mt.Rid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rid"))
	}
	if mt.Etype == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "etype"))
	}
	if mt.Backurl == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "backurl"))
	}
	if mt.From == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "from"))
	}
	return
}

// 事件监听列表 (default view)
//
// Identifier: vnd.ant.reg.list+json; view=default
type AntRegList struct {
	// 事件类型
	List []*AntReg `form:"list,omitempty" json:"list,omitempty" xml:"list,omitempty"`
	// 总数量
	Total int `form:"total" json:"total" xml:"total"`
}

// Validate validates the AntRegList media type instance.
func (mt *AntRegList) Validate() (err error) {
	for _, e := range mt.List {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// 注册事件监听成功 (default view)
//
// Identifier: vnd.ant.reg.result+json; view=default
type AntRegResult struct {
	// 成功标识
	OK *bool `form:"ok,omitempty" json:"ok,omitempty" xml:"ok,omitempty"`
}

// 注册事件监听成功 (failed view)
//
// Identifier: vnd.ant.reg.result+json; view=failed
type AntRegResultFailed struct {
	// 如果ok=false,失败原因
	Msg *string `form:"msg,omitempty" json:"msg,omitempty" xml:"msg,omitempty"`
	// 成功标识
	OK *bool `form:"ok,omitempty" json:"ok,omitempty" xml:"ok,omitempty"`
}

// 创建事件成功返回 (default view)
//
// Identifier: vnd.ant.result+json; view=default
type AntResult struct {
	// 事件唯一标识
	Eid *string `form:"eid,omitempty" json:"eid,omitempty" xml:"eid,omitempty"`
}
