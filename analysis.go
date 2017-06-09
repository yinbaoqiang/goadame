package main

import (
	"github.com/goadesign/goa"
	"github.com/yinbaoqiang/goadame/app"
)

// AnalysisController implements the analysis resource.
type AnalysisController struct {
	*goa.Controller
}

// NewAnalysisController creates a analysis controller.
func NewAnalysisController(service *goa.Service) *AnalysisController {
	return &AnalysisController{Controller: service.NewController("AnalysisController")}
}

// Hook runs the hook action.
func (c *AnalysisController) Hook(ctx *app.HookAnalysisContext) error {
	// AnalysisController_Hook: start_implement

	// Put your logic here

	// AnalysisController_Hook: end_implement
	res := app.AntEvenBackCollection{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *AnalysisController) List(ctx *app.ListAnalysisContext) error {
	// AnalysisController_List: start_implement

	// Put your logic here

	// AnalysisController_List: end_implement
	res := &app.AntRegList{}
	return ctx.OK(res)
}
