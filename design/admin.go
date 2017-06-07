package design

import (
	"github.com/goadesign/goa/design" // Use . imports to enable the DSL
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.Resource("regevent", func() { // Resources group related API endpoints
	apidsl.BasePath("/admin/event") // together. They map to REST resources for REST

	apidsl.Action("add", func() { // Actions define a single API endpoint together
		apidsl.Description("注册事件监听")        // with its path, parameters (both path
		apidsl.Routing(apidsl.POST("/reg")) // parameters and querystring values) and payload

		apidsl.Payload(func() { // (shape of the request body).
			apidsl.Member("etype", design.String, "事件类型")
			apidsl.Member("action", design.String, "事件行为,不设置该项则注册监听所有行为变化")
			apidsl.Member("from", design.String, "产生事件的服务器标识")
			apidsl.Member("bakurl", design.String, "回调路径")
			apidsl.Required("etype")
			apidsl.Required("bakurl")
			apidsl.Required("from")

		})
		apidsl.Response(design.OK, RegResultMedia)
	})
	apidsl.Action("remove", func() { // Actions define a single API endpoint together
		apidsl.Description("取消事件监听") // with its path, parameters (both path

		apidsl.Routing(apidsl.DELETE("/:rid")) // parameters and querystring values) and payload
		apidsl.Params(func() {                 // (shape of the request body).
			apidsl.Param("rid", design.String, "事件监听唯一标识")

		})

		apidsl.Response(design.OK, CreResultMedia)
	})
	apidsl.Action("list", func() { // Actions define a single API endpoint together
		apidsl.Description("注册事件监听")   // with its path, parameters (both path
		apidsl.Routing(apidsl.GET("")) // parameters and querystring values) and payload
		apidsl.Params(func() {
			apidsl.Param("page", design.Integer, "事件类型")
			apidsl.Param("count", design.Integer, "事件行为,不设置该项则注册监听所有行为变化")
		})
		apidsl.Response(design.OK, RegListMedia)
	})
})
