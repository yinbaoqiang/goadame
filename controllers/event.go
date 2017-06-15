package controllers

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/satori/go.uuid"
	"github.com/yinbaoqiang/goadame/app"
	"github.com/yinbaoqiang/goadame/engine"
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
	eid := uuid.NewV4().String()
	occtime := time.Now()
	if ctx.Payload.Occtime != nil {
		occtime = *ctx.Payload.Occtime
	}
	var data interface{}
	if ctx.Payload.Params != nil {
		data = *ctx.Payload.Params
	}

	evt := engine.Event{
		Eid:     eid,
		Etype:   ctx.Payload.Etype,
		Action:  ctx.Payload.Action,
		OccTime: occtime,
		From:    ctx.Payload.From,
		Data:    data,
	}
	// Put your logic here
	engine.Put(evt)
	// EventController_Put: end_implement
	res := &app.AntResult{Eid: &eid}
	return ctx.OK(res)
}

// Put runs the put action.
func (c *EventController) Put(ctx *app.PutEventContext) error {
	// EventController_Put: start_implement

	eid := ctx.Eid
	occtime := time.Now()
	if ctx.Payload.Occtime != nil {
		occtime = *ctx.Payload.Occtime
	}
	var data interface{}
	if ctx.Payload.Params != nil {
		data = *ctx.Payload.Params
	}

	evt := engine.Event{
		Eid:     eid,
		Etype:   ctx.Payload.Etype,
		Action:  ctx.Payload.Action,
		OccTime: occtime,
		From:    ctx.Payload.From,
		Data:    data,
	}
	// Put your logic here
	engine.Put(evt)
	// EventController_Put: end_implement
	res := &app.AntResult{Eid: &eid}
	return ctx.OK(res)
}
