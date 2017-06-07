//go:generate goagen bootstrap -d github.com/yinbaoqiang/goadame/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/yinbaoqiang/goadame/app"
)

func main() {
	// Create service
	service := goa.New("antevent")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "event" controller
	c := NewEventController(service)
	app.MountEventController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
