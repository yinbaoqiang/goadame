package main

import (
	"github.com/goadesign/goa"
	"github.com/yinbaoqiang/goadame/app"
)

// EventController implements the event resource.
type EventController struct {
	*goa.Controller
}

// NewEventController creates a event controller.
func NewEventController(service *goa.Service) *EventController {
	return &EventController{Controller: service.NewController("EventController")}
}

// Create runs the create action.
func (c *EventController) Create(ctx *app.CreateEventContext) error {
	// EventController_Create: start_implement

	// Put your logic here

	// EventController_Create: end_implement
	res := &app.AntEventCreResult{}
	return ctx.OK(res)
}

// Create2 runs the create2 action.
func (c *EventController) Create2(ctx *app.Create2EventContext) error {
	// EventController_Create2: start_implement

	// Put your logic here

	// EventController_Create2: end_implement
	res := &app.AntEventCreResult{}
	return ctx.OK(res)
}
