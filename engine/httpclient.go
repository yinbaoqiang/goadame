package engine

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the antevent service client.
type eventClient struct {
	*goaclient.Client
	Encoder *goa.HTTPEncoder
	Decoder *goa.HTTPDecoder
}

// newEventClient instantiates the client.
func newEventClient() *eventClient {
	client := &eventClient{
		Client:  goaclient.New(goaclient.HTTPClientDoer(http.DefaultClient)),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// 注册事件监听
func (c *eventClient) SendEvent(ctx context.Context, path string, payload Event) (*http.Response, error) {
	req, err := c.NewSendEventRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddListenRequest create the request corresponding to the add action endpoint of the listen resource.
func (c *eventClient) NewSendEventRequest(ctx context.Context, path string, payload Event) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	u, _ := url.Parse(path)
	c.Scheme = u.Scheme
	c.Host = u.Host
	req, err := http.NewRequest("PUT", path, &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}
