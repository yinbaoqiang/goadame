// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "antevent": regevent Resource Client
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
	"strconv"
)

// AddRegeventPayload is the regevent add action payload.
type AddRegeventPayload struct {
	// 事件行为,不设置该项则注册监听所有行为变化
	Action *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// 回调路径
	Bakurl string `form:"bakurl" json:"bakurl" xml:"bakurl"`
	// 事件类型
	Etype string `form:"etype" json:"etype" xml:"etype"`
	// 产生事件的服务器标识
	From string `form:"from" json:"from" xml:"from"`
}

// AddRegeventPath computes a request path to the add action of regevent.
func AddRegeventPath() string {

	return fmt.Sprintf("/v1/admin/event/reg")
}

// 注册事件监听
func (c *Client) AddRegevent(ctx context.Context, path string, payload *AddRegeventPayload) (*http.Response, error) {
	req, err := c.NewAddRegeventRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddRegeventRequest create the request corresponding to the add action endpoint of the regevent resource.
func (c *Client) NewAddRegeventRequest(ctx context.Context, path string, payload *AddRegeventPayload) (*http.Request, error) {
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

// ListRegeventPath computes a request path to the list action of regevent.
func ListRegeventPath() string {

	return fmt.Sprintf("/v1/admin/event")
}

// 注册事件监听
func (c *Client) ListRegevent(ctx context.Context, path string, count *int, page *int) (*http.Response, error) {
	req, err := c.NewListRegeventRequest(ctx, path, count, page)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListRegeventRequest create the request corresponding to the list action endpoint of the regevent resource.
func (c *Client) NewListRegeventRequest(ctx context.Context, path string, count *int, page *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if count != nil {
		tmp10 := strconv.Itoa(*count)
		values.Set("count", tmp10)
	}
	if page != nil {
		tmp11 := strconv.Itoa(*page)
		values.Set("page", tmp11)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RemoveRegeventPath computes a request path to the remove action of regevent.
func RemoveRegeventPath(rid string) string {
	param0 := rid

	return fmt.Sprintf("/v1/admin/event/%s", param0)
}

// 取消事件监听
func (c *Client) RemoveRegevent(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRemoveRegeventRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRemoveRegeventRequest create the request corresponding to the remove action endpoint of the regevent resource.
func (c *Client) NewRemoveRegeventRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
