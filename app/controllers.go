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
	"github.com/goadesign/goa/cors"
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

// AnalysisController is the controller interface for the Analysis actions.
type AnalysisController interface {
	goa.Muxer
	Hook(*HookAnalysisContext) error
	List(*ListAnalysisContext) error
}

// MountAnalysisController "mounts" a Analysis resource controller on the given service.
func MountAnalysisController(service *goa.Service, ctrl AnalysisController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHookAnalysisContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Hook(rctx)
	}
	service.Mux.Handle("GET", "/v1/admin/event/analysis/hook/:eid", ctrl.MuxHandler("hook", h, nil))
	service.LogInfo("mount", "ctrl", "Analysis", "action", "Hook", "route", "GET /v1/admin/event/analysis/hook/:eid")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListAnalysisContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v1/admin/event/analysis", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Analysis", "action", "List", "route", "GET /v1/admin/event/analysis")
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

// ListenController is the controller interface for the Listen actions.
type ListenController interface {
	goa.Muxer
	Add(*AddListenContext) error
	List(*ListListenContext) error
	Remove(*RemoveListenContext) error
	Update(*UpdateListenContext) error
}

// MountListenController "mounts" a Listen resource controller on the given service.
func MountListenController(service *goa.Service, ctrl ListenController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddListenContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AddListenPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Add(rctx)
	}
	service.Mux.Handle("POST", "/v1/admin/listen", ctrl.MuxHandler("add", h, unmarshalAddListenPayload))
	service.LogInfo("mount", "ctrl", "Listen", "action", "Add", "route", "POST /v1/admin/listen")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListListenContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v1/admin/listen", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Listen", "action", "List", "route", "GET /v1/admin/listen")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRemoveListenContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Remove(rctx)
	}
	service.Mux.Handle("DELETE", "/v1/admin/listen/:rid", ctrl.MuxHandler("remove", h, nil))
	service.LogInfo("mount", "ctrl", "Listen", "action", "Remove", "route", "DELETE /v1/admin/listen/:rid")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateListenContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateListenPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/v1/admin/listen/:rid", ctrl.MuxHandler("update", h, unmarshalUpdateListenPayload))
	service.LogInfo("mount", "ctrl", "Listen", "action", "Update", "route", "PUT /v1/admin/listen/:rid")
}

// unmarshalAddListenPayload unmarshals the request body into the context request data Payload field.
func unmarshalAddListenPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &addListenPayload{}
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

// unmarshalUpdateListenPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateListenPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateListenPayload{}
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

// PublicController is the controller interface for the Public actions.
type PublicController interface {
	goa.Muxer
	goa.FileServer
}

// MountPublicController "mounts" a Public resource controller on the given service.
func MountPublicController(service *goa.Service, ctrl PublicController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/ui/*filepath", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/ui/*filepath", "dist")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "dist", "route", "GET /ui/*filepath")

	h = ctrl.FileHandler("/ui/", "dist/index.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "dist/index.html", "route", "GET /ui/")
}

// handlePublicOrigin applies the CORS response headers corresponding to the origin.
func handlePublicOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
