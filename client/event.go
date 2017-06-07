// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "antevent": event Resource Client
//
// Command:
// $ goagen
// --design=github.com/yinbaoqiang/goadame/design
// --out=$(GOPATH)/src/github.com/yinbaoqiang/goadame
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// PostEventPayload is the event post action payload.
type PostEventPayload struct {
	// 事件行为
	Action string `form:"action" json:"action" xml:"action"`
	// 事件类型
	Etype string `form:"etype" json:"etype" xml:"etype"`
	// 产生事件的服务器标识
	From string `form:"from" json:"from" xml:"from"`
	// 事件发生时间
	Occtime *string `form:"occtime,omitempty" json:"occtime,omitempty" xml:"occtime,omitempty"`
	// 事件发生时间
	Params *interface{} `form:"params,omitempty" json:"params,omitempty" xml:"params,omitempty"`
}

// PostEventPath computes a request path to the post action of event.
func PostEventPath() string {

	return fmt.Sprintf("/v1/event")
}

// 创建一个事件
func (c *Client) PostEvent(ctx context.Context, path string, payload *PostEventPayload) (*http.Response, error) {
	req, err := c.NewPostEventRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPostEventRequest create the request corresponding to the post action endpoint of the event resource.
func (c *Client) NewPostEventRequest(ctx context.Context, path string, payload *PostEventPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// PutEventPayload is the event put action payload.
type PutEventPayload struct {
	// 事件行为
	Action string `form:"action" json:"action" xml:"action"`
	// 事件唯一标识
	Eid string `form:"eid" json:"eid" xml:"eid"`
	// 事件类型
	Etype string `form:"etype" json:"etype" xml:"etype"`
	// 产生事件的服务器标识
	From string `form:"from" json:"from" xml:"from"`
	// 事件发生时间
	Occtime *string `form:"occtime,omitempty" json:"occtime,omitempty" xml:"occtime,omitempty"`
	// 事件发生时间
	Params *interface{} `form:"params,omitempty" json:"params,omitempty" xml:"params,omitempty"`
}

// PutEventPath computes a request path to the put action of event.
func PutEventPath() string {

	return fmt.Sprintf("/v1/event")
}

// 创建一个事件
func (c *Client) PutEvent(ctx context.Context, path string, payload *PutEventPayload) (*http.Response, error) {
	req, err := c.NewPutEventRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPutEventRequest create the request corresponding to the put action endpoint of the event resource.
func (c *Client) NewPutEventRequest(ctx context.Context, path string, payload *PutEventPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}
