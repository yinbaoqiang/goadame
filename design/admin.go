package design

import (
	"github.com/goadesign/goa/design" // Use . imports to enable the DSL
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.Resource("listen", func() {
	apidsl.BasePath("/admin/listen")
	apidsl.Response(design.InternalServerError, ErrMedia)
	apidsl.Action("add", func() {
		apidsl.Description("注册事件监听")
		apidsl.Routing(apidsl.POST(""))

		apidsl.Payload(func() {
			apidsl.Member("etype", design.String, "事件类型")
			apidsl.Member("action", design.String, "事件行为,不设置该项则注册监听所有行为变化")
			apidsl.Member("from", design.String, "注册事件监听的服务器标识")
			apidsl.Member("hookurl", design.String, "钩子url")
			apidsl.Required("etype")
			apidsl.Required("hookurl")
			apidsl.Required("from")

		})
		apidsl.Response(design.OK, RegResultMedia)
	})
	apidsl.Action("update", func() {
		apidsl.Description("修改注册事件监听")
		apidsl.Routing(apidsl.PUT("/:rid"))
		apidsl.Payload(func() {
			apidsl.Member("etype", design.String, "事件类型")
			apidsl.Member("action", design.String, "事件行为,不设置该项则注册监听所有行为变化")
			apidsl.Member("from", design.String, "产生事件的服务器标识")
			apidsl.Member("hookurl", design.String, "钩子url")
			apidsl.Required("etype")
			apidsl.Required("hookurl")
			apidsl.Required("from")

		})
		apidsl.Response(design.OK, RegResultMedia)
	})
	apidsl.Action("remove", func() {
		apidsl.Description("取消事件监听")

		apidsl.Routing(apidsl.DELETE("/:rid"))
		apidsl.Params(func() {
			apidsl.Param("rid", design.String, "事件监听唯一标识")

		})

		apidsl.Response(design.OK, CreResultMedia)
	})
	apidsl.Action("list", func() {
		apidsl.Description("获取注册事件监听列表")
		apidsl.Routing(apidsl.GET(""))
		apidsl.Params(func() {
			apidsl.Param("page", design.Integer, "分页", func() {
				apidsl.Minimum(1)
			})
			apidsl.Param("count", design.Integer, "分页数量", func() {
				apidsl.Minimum(5)
			})
			apidsl.Param("etype", design.String, "事件类型,不设置则查询所有事件类型")
			apidsl.Param("action", design.String, "事件行为,不设置该项则查询所有行为")
		})
		apidsl.Response(design.OK, RegListMedia)
	})

})
var _ = apidsl.Resource("analysis", func() { // 事件分析
	apidsl.BasePath("/admin/event/analysis") // 基础URL

	apidsl.Response(design.InternalServerError, ErrMedia)
	apidsl.Action("list", func() {
		apidsl.Description("事件发生历史")
		apidsl.Routing(apidsl.GET(""))
		apidsl.Params(func() {
			apidsl.Param("page", design.Integer, "查询分页", func() {
				apidsl.Minimum(1)
			})
			apidsl.Param("count", design.Integer, "分页数量", func() {
				apidsl.Minimum(5)
			})
			apidsl.Param("etype", design.String, "事件类型")
			apidsl.Param("action", design.String, "行为")
			apidsl.Param("from", design.String, "来源")
		})
		apidsl.Response(design.OK, RegListMedia)
	})

	apidsl.Action("hook", func() {
		apidsl.Description("事件回调执行情况")
		apidsl.Routing(apidsl.GET("/hook/:eid"))
		apidsl.Params(func() {
			apidsl.Param("eid", design.String, "事件事件唯一标识")
			apidsl.Param("action", design.Integer, "事件行为,指定该值则只返回该事件该行为的钩子调用情况,不指定,返回该事件所有行为调用情况")
		})
		apidsl.Response(design.OK, apidsl.CollectionOf("vnd.ant.even.back+json"))
	})
})
