package controllers

import (
	"github.com/goadesign/goa"
	"github.com/yinbaoqiang/goadame/app"
	"github.com/yinbaoqiang/goadame/store"
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
	l := app.AntListen{
		Action:  ctx.Payload.Action,
		Etype:   ctx.Payload.Etype,
		From:    ctx.Payload.From,
		Hookurl: ctx.Payload.Hookurl,
	}
	// Put your logic here
	err = store.ListenStore().Add(l)
	// ListenController_Add: end_implement
	if err != nil {
		msg := "新增监听失败:" + err.Error()
		return ctx.InternalServerError(&app.AntError{Msg: &msg})
	}
	res := &app.AntRegResult{OK: true}

	return ctx.OK(res)
}

// List runs the list action.
func (c *ListenController) List(ctx *app.ListListenContext) error {
	// ListenController_List: start_implement

	count := 0

	if ctx.Count != nil {
		count = *ctx.Count
	}
	if count < 1 {
		count = 1
	}
	action, etype, pid := "", "", ""
	if ctx.Action != nil {
		action = *ctx.Action
	}

	if ctx.Etype != nil {
		etype = *ctx.Etype
	}
	if ctx.Previd != nil {
		pid = *ctx.Previd
	}

	res := &app.AntListenList{}
	total, liss, err := store.ListenStore().List(etype, action, pid, count)
	// ListenController_List: end_implement
	if err != nil {
		msg := "新增监听失败:" + err.Error()
		return ctx.InternalServerError(&app.AntError{Msg: &msg})
	}
	res.List = liss
	res.Total = total
	return ctx.OK(res)
}

// Remove runs the remove action.
func (c *ListenController) Remove(ctx *app.RemoveListenContext) error {
	// ListenController_Remove: start_implement
	rid := ctx.Rid
	if rid == "" {
		msg := "请求删除错误。"
		return ctx.BadRequest(&app.AntError{Msg: &msg})
	}
	// Put your logic here

	err := store.ListenStore().Rmove(rid)
	// ListenController_List: end_implement
	if err != nil {
		msg := "删除监听失败:" + err.Error()
		return ctx.InternalServerError(&app.AntError{Msg: &msg})
	}
	// ListenController_Remove: end_implement
	res := &app.AntResult{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *ListenController) Update(ctx *app.UpdateListenContext) error {
	err := ctx.Payload.Validate()
	if err != nil {
		msg := "参数错误:" + err.Error()
		return ctx.BadRequest(&app.AntError{Msg: &msg})
	}
	// ListenController_Update: start_implement
	l := app.AntListen{
		Rid:     ctx.Rid,
		Action:  ctx.Payload.Action,
		Etype:   ctx.Payload.Etype,
		From:    ctx.Payload.From,
		Hookurl: ctx.Payload.Hookurl,
	}
	// Put your logic here
	// Put your logic here
	err = store.ListenStore().Update(l)
	// ListenController_Add: end_implement
	if err != nil {
		msg := "新增监听失败:" + err.Error()
		return ctx.InternalServerError(&app.AntError{Msg: &msg})
	}
	// ListenController_Update: end_implement
	res := &app.AntRegResult{OK: true}
	return ctx.OK(res)
}
