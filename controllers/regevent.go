package controllers

import (
	"github.com/goadesign/goa"
	"github.com/yinbaoqiang/goadame/app"
)

// RegeventController implements the regevent resource.
type RegeventController struct {
	*goa.Controller
}

// NewRegeventController creates a regevent controller.
func NewRegeventController(service *goa.Service) *RegeventController {
	return &RegeventController{Controller: service.NewController("RegeventController")}
}

// Add runs the add action.
func (c *RegeventController) Add(ctx *app.AddRegeventContext) error {
	// RegeventController_Add: start_implement

	// Put your logic here

	// RegeventController_Add: end_implement
	res := &app.AntRegResult{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *RegeventController) List(ctx *app.ListRegeventContext) error {
	// RegeventController_List: start_implement

	// Put your logic here

	// RegeventController_List: end_implement
	res := &app.AntRegList{}
	return ctx.OK(res)
}

// Remove runs the remove action.
func (c *RegeventController) Remove(ctx *app.RemoveRegeventContext) error {
	// RegeventController_Remove: start_implement

	// Put your logic here

	// RegeventController_Remove: end_implement
	res := &app.AntResult{}
	return ctx.OK(res)
}
