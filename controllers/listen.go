package controllers

import (
	"github.com/goadesign/goa"
	"github.com/yinbaoqiang/goadame/app"
)

// ListenController implements the listen resource.
type ListenController struct {
	*goa.Controller
}

// NewListenController creates a listen controller.
func NewListenController(service *goa.Service) *ListenController {
	return &ListenController{Controller: service.NewController("ListenController")}
}

// Add runs the add action.
func (c *ListenController) Add(ctx *app.AddListenContext) error {
	// ListenController_Add: start_implement
	err := ctx.Payload.Validate()
	if err != nil {
		msg := "参数错误:" + err.Error()
		return ctx.BadRequest(&app.AntError{Msg: &msg})
	}
	action := *ctx.Payload.Action
	etype := ctx.Payload.Etype
	from := ctx.Payload.From
	hookurl := ctx.Payload.Hookurl
	// Put your logic here

	// ListenController_Add: end_implement

	res := &app.AntRegResult{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *ListenController) List(ctx *app.ListListenContext) error {
	// ListenController_List: start_implement
	ctx.Count
	ctx.Page
	ctx.Action
	ctx.Etype
	// Put your logic here

	// ListenController_List: end_implement
	res := &app.AntRegList{}
	return ctx.OK(res)
}

// Remove runs the remove action.
func (c *ListenController) Remove(ctx *app.RemoveListenContext) error {
	// ListenController_Remove: start_implement
	ctx.Rid
	// Put your logic here

	// ListenController_Remove: end_implement
	res := &app.AntResult{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *ListenController) Update(ctx *app.UpdateListenContext) error {
	// ListenController_Update: start_implement

	// Put your logic here

	// ListenController_Update: end_implement
	res := &app.AntRegResult{}
	return ctx.OK(res)
}
