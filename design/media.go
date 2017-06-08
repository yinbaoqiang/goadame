package design

import (
	"github.com/goadesign/goa/design" // Use . imports to enable the DSL
	"github.com/goadesign/goa/design/apidsl"
)

// CreResultMedia defines the media type used to render bottles.
var CreResultMedia = apidsl.MediaType("vnd.ant.result+json", func() {
	apidsl.Description("创建事件成功返回")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("eid", design.String, "事件唯一标识")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("eid") // Media types may have multiple views and must
	})
})

// ErrMedia 定义一个错误处理类型
var ErrMedia = apidsl.MediaType("vnd.ant.error+json", func() {
	apidsl.Description("处理失败")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("msg", design.String, "错误描述")

	})
	apidsl.View("default", func() {
		apidsl.Attribute("msg")
	})
})

// RegResultMedia defines the media type used to render bottles.
var RegResultMedia = apidsl.MediaType("vnd.ant.reg.result+json", func() {
	apidsl.Description("注册事件监听成功")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("ok", design.Boolean, "成功标识")
		apidsl.Attribute("msg", design.String, "如果ok=false,失败原因")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("ok") // Media types may have multiple views and must
	})
	apidsl.View("failed", func() { // View defines a rendering of the media type.
		apidsl.Attribute("ok")  // Media types may have multiple views and must
		apidsl.Attribute("msg") // Media types may have multiple views and must
	})
})

// RegInfoMedia 注册监听事件列表
var RegInfoMedia = apidsl.MediaType("vnd.ant.reg+json", func() {
	apidsl.Description("事件监听信息")

	apidsl.Attributes(func() { // (shape of the request body).
		apidsl.Attribute("rid", design.String, "注册事件监听唯一标识")
		apidsl.Attribute("etype", design.String, "事件类型")
		apidsl.Attribute("action", design.String, "事件行为,不设置该项则注册监听所有行为变化")
		apidsl.Attribute("from", design.String, "产生事件的服务器标识")
		apidsl.Attribute("backurl", design.String, "回调路径")
		apidsl.Required("rid")
		apidsl.Required("etype")
		apidsl.Required("backurl")
		apidsl.Required("from")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("rid")     // Media types may have multiple views and must
		apidsl.Attribute("etype")   // Media types may have multiple views and must
		apidsl.Attribute("action")  // Media types may have multiple views and must
		apidsl.Attribute("from")    // Media types may have multiple views and must
		apidsl.Attribute("backurl") // Media types may have multiple views and must
	})
})

// RegListMedia 注册监听事件列表
var RegListMedia = apidsl.MediaType("vnd.ant.reg.list+json", func() {
	apidsl.Description("事件监听列表")

	apidsl.Attributes(func() { // (shape of the request body).
		apidsl.Attribute("total", design.Integer, "总数量")
		apidsl.Attribute("list", apidsl.ArrayOf(RegInfoMedia), "事件类型")
		apidsl.Required("total")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("total") // Media types may have multiple views and must
		apidsl.Attribute("list")  // Media types may have multiple views and must
	})
})

// EventHisInfoMedia 事件信息
var EventHisInfoMedia = apidsl.MediaType("vnd.ant.reg+json", func() {
	apidsl.Description("事件监听信息")

	apidsl.Attributes(func() { // (shape of the request body).
		apidsl.Attribute("eid", design.String, "事件唯一标识")
		apidsl.Attribute("etype", design.String, "事件类型")
		apidsl.Attribute("action", design.String, "事件行为,不设置该项则注册监听所有行为变化")
		apidsl.Attribute("from", design.String, "产生事件的服务器标识")
		apidsl.Attribute("occtime", design.DateTime, "事件发送时间")
		apidsl.Attribute("data", design.Any, "事件数据")

		apidsl.Required("eid")
		apidsl.Required("etype")
		apidsl.Required("backurl")
		apidsl.Required("from")
		apidsl.Required("occtime")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("rid")     // Media types may have multiple views and must
		apidsl.Attribute("etype")   // Media types may have multiple views and must
		apidsl.Attribute("action")  // Media types may have multiple views and must
		apidsl.Attribute("from")    // Media types may have multiple views and must
		apidsl.Attribute("backurl") // Media types may have multiple views and must
	})
})

// EventHisListMedia 事件历史列表
var EventHisListMedia = apidsl.MediaType("vnd.ant.event.history.list+json", func() {
	apidsl.Description("事件监听列表")

	apidsl.Attributes(func() { // (shape of the request body).
		apidsl.Attribute("total", design.Integer, "总数量")
		apidsl.Attribute("list", apidsl.ArrayOf(EventHisInfoMedia), "事件类型")
		apidsl.Required("total")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("total") // Media types may have multiple views and must
		apidsl.Attribute("list")  // Media types may have multiple views and must
	})
})

// EventBackInfoMedia 事件监听回调执行情况
var EventBackInfoMedia = apidsl.MediaType("vnd.ant.even.back+json", func() {
	apidsl.Description("事件监听信息")

	apidsl.Attributes(func() { // (shape of the request body).
		apidsl.Attribute("eid", design.String, "事件唯一标识")
		apidsl.Attribute("backurl", design.String, "回调连接")
		apidsl.Attribute("startTime", design.DateTime, "开始执行时间")
		apidsl.Attribute("endTime", design.DateTime, "执行结束时间")
		apidsl.Attribute("execTime", design.Integer, "执行时间,单位纳秒")
		apidsl.Attribute("success", design.Boolean, "执行是否成功")
		apidsl.Attribute("error", design.String, "执行错误时的错误信息")

		apidsl.Required("eid")
		apidsl.Required("backurl")
		apidsl.Required("startTime")
		apidsl.Required("endTime")
		apidsl.Required("execTime")
		apidsl.Required("success")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("eid")       // Media types may have multiple views and must
		apidsl.Attribute("backurl")   // Media types may have multiple views and must
		apidsl.Attribute("startTime") // Media types may have multiple views and must
		apidsl.Attribute("endTime")   // Media types may have multiple views and must
		apidsl.Attribute("execTime")  // Media types may have multiple views and must
		apidsl.Attribute("success")   // Media types may have multiple views and must
		apidsl.Attribute("error")     // Media types may have multiple views and must
	})
})
