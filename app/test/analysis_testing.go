// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "antevent": analysis TestHelpers
//
// Command:
// $ goagen
// --design=github.com/yinbaoqiang/goadame/design
// --out=$(GOPATH)/src/github.com/yinbaoqiang/goadame
// --version=v1.2.0-dirty

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/yinbaoqiang/goadame/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// HookAnalysisBadRequest runs the method Hook of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func HookAnalysisBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AnalysisController, eid string, action *int) (http.ResponseWriter, *app.AntError) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if action != nil {
		sliceVal := []string{strconv.Itoa(*action)}
		query["action"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/admin/event/analysis/hook/%v", eid),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["eid"] = []string{fmt.Sprintf("%v", eid)}
	if action != nil {
		sliceVal := []string{strconv.Itoa(*action)}
		prms["action"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AnalysisTest"), rw, req, prms)
	hookCtx, _err := app.NewHookAnalysisContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Hook(hookCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.AntError
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AntError)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AntError", resp)
		}
	}

	// Return results
	return rw, mt
}

// HookAnalysisInternalServerError runs the method Hook of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func HookAnalysisInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AnalysisController, eid string, action *int) (http.ResponseWriter, *app.AntError) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if action != nil {
		sliceVal := []string{strconv.Itoa(*action)}
		query["action"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/admin/event/analysis/hook/%v", eid),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["eid"] = []string{fmt.Sprintf("%v", eid)}
	if action != nil {
		sliceVal := []string{strconv.Itoa(*action)}
		prms["action"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AnalysisTest"), rw, req, prms)
	hookCtx, _err := app.NewHookAnalysisContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Hook(hookCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt *app.AntError
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AntError)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AntError", resp)
		}
	}

	// Return results
	return rw, mt
}

// HookAnalysisOK runs the method Hook of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func HookAnalysisOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AnalysisController, eid string, action *int) (http.ResponseWriter, app.AntEvenBackCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if action != nil {
		sliceVal := []string{strconv.Itoa(*action)}
		query["action"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/admin/event/analysis/hook/%v", eid),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["eid"] = []string{fmt.Sprintf("%v", eid)}
	if action != nil {
		sliceVal := []string{strconv.Itoa(*action)}
		prms["action"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AnalysisTest"), rw, req, prms)
	hookCtx, _err := app.NewHookAnalysisContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Hook(hookCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.AntEvenBackCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.AntEvenBackCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AntEvenBackCollection", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ListAnalysisBadRequest runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListAnalysisBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AnalysisController, action *string, count *int, etype *string, from *string, page *int) (http.ResponseWriter, *app.AntError) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if action != nil {
		sliceVal := []string{*action}
		query["action"] = sliceVal
	}
	if count != nil {
		sliceVal := []string{strconv.Itoa(*count)}
		query["count"] = sliceVal
	}
	if etype != nil {
		sliceVal := []string{*etype}
		query["etype"] = sliceVal
	}
	if from != nil {
		sliceVal := []string{*from}
		query["from"] = sliceVal
	}
	if page != nil {
		sliceVal := []string{strconv.Itoa(*page)}
		query["page"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/admin/event/analysis"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if action != nil {
		sliceVal := []string{*action}
		prms["action"] = sliceVal
	}
	if count != nil {
		sliceVal := []string{strconv.Itoa(*count)}
		prms["count"] = sliceVal
	}
	if etype != nil {
		sliceVal := []string{*etype}
		prms["etype"] = sliceVal
	}
	if from != nil {
		sliceVal := []string{*from}
		prms["from"] = sliceVal
	}
	if page != nil {
		sliceVal := []string{strconv.Itoa(*page)}
		prms["page"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AnalysisTest"), rw, req, prms)
	listCtx, _err := app.NewListAnalysisContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.AntError
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AntError)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AntError", resp)
		}
	}

	// Return results
	return rw, mt
}

// ListAnalysisInternalServerError runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListAnalysisInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AnalysisController, action *string, count *int, etype *string, from *string, page *int) (http.ResponseWriter, *app.AntError) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if action != nil {
		sliceVal := []string{*action}
		query["action"] = sliceVal
	}
	if count != nil {
		sliceVal := []string{strconv.Itoa(*count)}
		query["count"] = sliceVal
	}
	if etype != nil {
		sliceVal := []string{*etype}
		query["etype"] = sliceVal
	}
	if from != nil {
		sliceVal := []string{*from}
		query["from"] = sliceVal
	}
	if page != nil {
		sliceVal := []string{strconv.Itoa(*page)}
		query["page"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/admin/event/analysis"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if action != nil {
		sliceVal := []string{*action}
		prms["action"] = sliceVal
	}
	if count != nil {
		sliceVal := []string{strconv.Itoa(*count)}
		prms["count"] = sliceVal
	}
	if etype != nil {
		sliceVal := []string{*etype}
		prms["etype"] = sliceVal
	}
	if from != nil {
		sliceVal := []string{*from}
		prms["from"] = sliceVal
	}
	if page != nil {
		sliceVal := []string{strconv.Itoa(*page)}
		prms["page"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AnalysisTest"), rw, req, prms)
	listCtx, _err := app.NewListAnalysisContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt *app.AntError
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AntError)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AntError", resp)
		}
	}

	// Return results
	return rw, mt
}

// ListAnalysisOK runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListAnalysisOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AnalysisController, action *string, count *int, etype *string, from *string, page *int) (http.ResponseWriter, *app.AntEventHistoryList) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if action != nil {
		sliceVal := []string{*action}
		query["action"] = sliceVal
	}
	if count != nil {
		sliceVal := []string{strconv.Itoa(*count)}
		query["count"] = sliceVal
	}
	if etype != nil {
		sliceVal := []string{*etype}
		query["etype"] = sliceVal
	}
	if from != nil {
		sliceVal := []string{*from}
		query["from"] = sliceVal
	}
	if page != nil {
		sliceVal := []string{strconv.Itoa(*page)}
		query["page"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/admin/event/analysis"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if action != nil {
		sliceVal := []string{*action}
		prms["action"] = sliceVal
	}
	if count != nil {
		sliceVal := []string{strconv.Itoa(*count)}
		prms["count"] = sliceVal
	}
	if etype != nil {
		sliceVal := []string{*etype}
		prms["etype"] = sliceVal
	}
	if from != nil {
		sliceVal := []string{*from}
		prms["from"] = sliceVal
	}
	if page != nil {
		sliceVal := []string{strconv.Itoa(*page)}
		prms["page"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AnalysisTest"), rw, req, prms)
	listCtx, _err := app.NewListAnalysisContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.AntEventHistoryList
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AntEventHistoryList)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AntEventHistoryList", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
