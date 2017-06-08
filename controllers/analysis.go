package controllers

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

// Back runs the back action.
func (c *AnalysisController) Back(ctx *app.BackAnalysisContext) error {
	// AnalysisController_Back: start_implement

	// Put your logic here

	// AnalysisController_Back: end_implement
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
