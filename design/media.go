package design

import (
	"github.com/goadesign/goa/design" // Use . imports to enable the DSL
	"github.com/goadesign/goa/design/apidsl"
)

// CreResultMedia defines the media type used to render bottles.
var CreResultMedia = apidsl.MediaType("application/vnd.ant.result+json", func() {
	apidsl.Description("创建事件成功返回")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("eid", design.String, "事件唯一标识")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("eid") // Media types may have multiple views and must
	})
})

// ErrMedia 定义一个错误处理类型
var ErrMedia = apidsl.MediaType("application/vnd.ant.error+json", func() {
	apidsl.Description("处理失败")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("msg", design.String, "错误描述")

	})
	apidsl.View("default", func() {
		apidsl.Attribute("msg")
	})
})

// CreResultMedia defines the media type used to render bottles.
var RegResultMedia = apidsl.MediaType("application/vnd.ant.result+json", func() {
	apidsl.Description("注册事件监听成功")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("ok", design.Boolean, "成功标识")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("ok") // Media types may have multiple views and must
	})
})

// RegInfoMedia 注册监听事件列表
var RegInfoMedia = apidsl.MediaType("application/vnd.ant.reg+json", func() {
	apidsl.Description("事件监听信息")

	apidsl.Attributes(func() { // (shape of the request body).
		apidsl.Attribute("rid", design.String, "注册事件监听唯一标识")
		apidsl.Attribute("etype", design.String, "事件类型")
		apidsl.Attribute("action", design.String, "事件行为,不设置该项则注册监听所有行为变化")
		apidsl.Attribute("from", design.String, "产生事件的服务器标识")
		apidsl.Attribute("bakurl", design.String, "回调路径")
		apidsl.Required("rid")
		apidsl.Required("etype")
		apidsl.Required("bakurl")
		apidsl.Required("from")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("ridok")  // Media types may have multiple views and must
		apidsl.Attribute("etype")  // Media types may have multiple views and must
		apidsl.Attribute("action") // Media types may have multiple views and must
		apidsl.Attribute("from")   // Media types may have multiple views and must
		apidsl.Attribute("bakurl") // Media types may have multiple views and must
	})
})

// RegListMedia 注册监听事件列表
var RegListMedia = apidsl.MediaType("application/vnd.ant.reg.list+json", func() {
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
