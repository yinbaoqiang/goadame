//go:generate goagen bootstrap -d github.com/yinbaoqiang/goadame/design

package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/yinbaoqiang/goadame/app"
	"github.com/yinbaoqiang/goadame/controllers"
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

	initArgs(os.Args, service)

}
