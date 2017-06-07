package design

import (
	"github.com/goadesign/goa/design" // Use . imports to enable the DSL
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.API("antevent", func() { // API defines the microservice endpoint and
	apidsl.Title("蚁动事件引擎")           // other global properties. There should be one
	apidsl.Description("这是一个事件服务引擎") // and exactly one API definition appearing in
	apidsl.Scheme("http")            // the design.
	apidsl.Host("localhost:8080")
	apidsl.Consumes("application/json") // Media types supported by the API
	apidsl.Produces("application/json") // Media types generated by the API
	apidsl.BasePath("/v1")
})

var _ = apidsl.Resource("event", func() { // Resources group related API endpoints
	apidsl.BasePath("/event") // together. They map to REST resources for REST

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
		apidsl.Response(design.NotFound)
	})
	apidsl.Action("post", func() { // Actions define a single API endpoint together
		apidsl.Description("创建一个事件") // with its path, parameters (both path

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

// CreResultMedia defines the media type used to render bottles.
var CreResultMedia = apidsl.MediaType("application/vnd.ant.event.cre.result+json", func() {
	apidsl.Description("创建事件成功返回")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("eid", design.String, "事件唯一标识")

	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("eid") // Media types may have multiple views and must
	})
})
