// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "antevent": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/yinbaoqiang/goadame/design
// --out=$(GOPATH)/src/github.com/yinbaoqiang/goadame
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// EventController is the controller interface for the Event actions.
type EventController interface {
	goa.Muxer
	Post(*PostEventContext) error
	Put(*PutEventContext) error
}

// MountEventController "mounts" a Event resource controller on the given service.
func MountEventController(service *goa.Service, ctrl EventController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPostEventContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PostEventPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Post(rctx)
	}
	service.Mux.Handle("POST", "/v1/event", ctrl.MuxHandler("post", h, unmarshalPostEventPayload))
	service.LogInfo("mount", "ctrl", "Event", "action", "Post", "route", "POST /v1/event")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPutEventContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PutEventPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Put(rctx)
	}
	service.Mux.Handle("PUT", "/v1/event/:eid", ctrl.MuxHandler("put", h, unmarshalPutEventPayload))
	service.LogInfo("mount", "ctrl", "Event", "action", "Put", "route", "PUT /v1/event/:eid")
}

// unmarshalPostEventPayload unmarshals the request body into the context request data Payload field.
func unmarshalPostEventPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &postEventPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalPutEventPayload unmarshals the request body into the context request data Payload field.
func unmarshalPutEventPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &putEventPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
