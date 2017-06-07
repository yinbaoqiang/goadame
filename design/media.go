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
