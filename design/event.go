package design

import (
	"github.com/goadesign/goa/design" // Use . imports to enable the DSL
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.Resource("event", func() { // Resources group related API endpoints
	apidsl.BasePath("/event") // together. They map to REST resources for REST

	apidsl.Response(design.InternalServerError, ErrMedia)
	apidsl.Action("put", func() { // Actions define a single API endpoint together
		apidsl.Description("创建一个事件")        // with its path, parameters (both path
		apidsl.Routing(apidsl.PUT("/:eid")) // parameters and querystring values) and payload

		apidsl.Params(func() {
			apidsl.Param("eid", design.String, "事件唯一标识")
			apidsl.Required("eid")
		})
		apidsl.Payload(func() { // (shape of the request body).
			apidsl.Member("etype", design.String, "事件类型")
			apidsl.Member("action", design.String, "事件行为")
			apidsl.Member("from", design.String, "产生事件的服务器标识")
			apidsl.Member("occtime", design.String, "事件发生时间")
			apidsl.Member("params", design.Any, "事件发生时间")
			apidsl.Required("etype")
			apidsl.Required("action")
			apidsl.Required("from")

		})
		apidsl.Response(design.OK, CreResultMedia)
	})
	apidsl.Action("post", func() { // Actions define a single API endpoint together
		apidsl.Description("创建一个事件") // with its path, parameters (both path

		apidsl.Response(design.InternalServerError, ErrMedia)
		apidsl.Routing(apidsl.POST("", func() {

		})) // parameters and querystring values) and payload
		apidsl.Payload(func() { // (shape of the request body).
			apidsl.Member("etype", design.String, "事件类型")
			apidsl.Member("action", design.String, "事件行为")
			apidsl.Member("from", design.String, "产生事件的服务器标识")
			apidsl.Member("occtime", design.String, "事件发生时间")
			apidsl.Member("params", design.Any, "事件发生时间")
			apidsl.Required("etype")
			apidsl.Required("action")
			apidsl.Required("from")

		})

		apidsl.Response(design.OK, CreResultMedia)
	})
})
