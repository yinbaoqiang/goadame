package controllers

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

// Post runs the post action.
func (c *EventController) Post(ctx *app.PostEventContext) error {
	// EventController_Post: start_implement

	// Put your logic here

	// EventController_Post: end_implement
	res := &app.AntResult{}
	return ctx.OK(res)
}

// Put runs the put action.
func (c *EventController) Put(ctx *app.PutEventContext) error {
	// EventController_Put: start_implement

	// Put your logic here

	// EventController_Put: end_implement
	res := &app.AntResult{}
	return ctx.OK(res)
}
