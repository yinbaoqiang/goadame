//go:generate goagen bootstrap -d github.com/yinbaoqiang/goadame/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/yinbaoqiang/goadame/app"
	"github.com/yinbaoqiang/goadame/controllers"
	"github.com/yinbaoqiang/goadame/engine"
)

func main() {
	// Create service
	service := goa.New("antevent")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "analysis" controller
	c := controllers.NewAnalysisController(service)
	app.MountAnalysisController(service, c)
	// Mount "event" controller
	c2 := controllers.NewEventController(service)
	app.MountEventController(service, c2)
	// Mount "regevent" controller
	c3 := controllers.NewListenController(service)
	app.MountListenController(service, c3)
	c4 := controllers.NewPublicController(service)
	app.MountPublicController(service, c4)
	c5 := controllers.NewSwaggerController(service)
	app.MountSwaggerController(service, c5)

	// 初始化应用
	initApp([]string{"127.0.0.1"}, 3)
	// 启动事件引擎
	engine.DefaultEnginer().Start()
	defer engine.DefaultEnginer().Stop()
	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)

	}

}
