// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "antevent": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/yinbaoqiang/goadame/design
// --out=$(GOPATH)/src/github.com/yinbaoqiang/goadame
// --version=v1.2.0-dirty

package client

import (
	"net/http"
	"time"

	"github.com/goadesign/goa"
)

// AntError 处理失败 (default view)
//
// Identifier: vnd.ant.error+json; view=default
type AntError struct {
	// 错误描述
	Msg *string `form:"msg,omitempty" json:"msg,omitempty" xml:"msg,omitempty"`
}

// DecodeAntError decodes the AntError instance encoded in resp body.
func (c *Client) DecodeAntError(resp *http.Response) (*AntError, error) {
	var decoded AntError
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AntEvenBack 事件监听信息 (default view)
//
// Identifier: vnd.ant.even.back+json; view=default
type AntEvenBack struct {
	// 事件唯一标识
	Eid string `form:"eid" json:"eid" xml:"eid"`
	// 执行结束时间
	EndTime time.Time `form:"endTime" json:"endTime" xml:"endTime"`
	// 执行错误时的错误信息
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
	// 执行时间,单位纳秒
	ExecTime int `form:"execTime" json:"execTime" xml:"execTime"`
	// 钩子url
	Hookurl string `form:"hookurl" json:"hookurl" xml:"hookurl"`
	// 开始执行时间
	StartTime time.Time `form:"startTime" json:"startTime" xml:"startTime"`
	// 执行是否成功
	Success bool `form:"success" json:"success" xml:"success"`
}

// Validate validates the AntEvenBack media type instance.
func (mt *AntEvenBack) Validate() (err error) {
	if mt.Eid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eid"))
	}
	if mt.Hookurl == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "hookurl"))
	}

	return
}

// DecodeAntEvenBack decodes the AntEvenBack instance encoded in resp body.
func (c *Client) DecodeAntEvenBack(resp *http.Response) (*AntEvenBack, error) {
	var decoded AntEvenBack
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AntEvenBackCollection is the media type for an array of AntEvenBack (default view)
//
// Identifier: vnd.ant.even.back+json; type=collection; view=default
type AntEvenBackCollection []*AntEvenBack

// Validate validates the AntEvenBackCollection media type instance.
func (mt AntEvenBackCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeAntEvenBackCollection decodes the AntEvenBackCollection instance encoded in resp body.
func (c *Client) DecodeAntEvenBackCollection(resp *http.Response) (AntEvenBackCollection, error) {
	var decoded AntEvenBackCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// AntEventHistoryList 事件监听列表 (default view)
//
// Identifier: vnd.ant.event.history.list+json; view=default
type AntEventHistoryList struct {
	// 事件类型
	List []*AntHistoryInfo `form:"list,omitempty" json:"list,omitempty" xml:"list,omitempty"`
	// 总数量
	Total int `form:"total" json:"total" xml:"total"`
}

// Validate validates the AntEventHistoryList media type instance.
func (mt *AntEventHistoryList) Validate() (err error) {
	for _, e := range mt.List {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeAntEventHistoryList decodes the AntEventHistoryList instance encoded in resp body.
func (c *Client) DecodeAntEventHistoryList(resp *http.Response) (*AntEventHistoryList, error) {
	var decoded AntEventHistoryList
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AntHistoryInfo 事件监听信息 (default view)
//
// Identifier: vnd.ant.history.info+json; view=default
type AntHistoryInfo struct {
	// 事件行为,不设置该项则注册监听所有行为变化
	Action *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// 事件唯一标识
	Eid string `form:"eid" json:"eid" xml:"eid"`
	// 事件类型
	Etype string `form:"etype" json:"etype" xml:"etype"`
	// 产生事件的服务器标识
	From string `form:"from" json:"from" xml:"from"`
}

// Validate validates the AntHistoryInfo media type instance.
func (mt *AntHistoryInfo) Validate() (err error) {
	if mt.Eid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eid"))
	}
	if mt.Etype == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "etype"))
	}
	if mt.From == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "from"))
	}
	return
}

// DecodeAntHistoryInfo decodes the AntHistoryInfo instance encoded in resp body.
func (c *Client) DecodeAntHistoryInfo(resp *http.Response) (*AntHistoryInfo, error) {
	var decoded AntHistoryInfo
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AntListen 事件监听信息 (default view)
//
// Identifier: vnd.ant.listen+json; view=default
type AntListen struct {
	// 事件行为,不设置该项则注册监听所有行为变化
	Action *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// 事件类型
	Etype string `form:"etype" json:"etype" xml:"etype"`
	// 产生事件的服务器标识
	From string `form:"from" json:"from" xml:"from"`
	// 钩子url
	Hookurl string `form:"hookurl" json:"hookurl" xml:"hookurl"`
	// 注册事件监听唯一标识
	Rid string `form:"rid" json:"rid" xml:"rid"`
}

// Validate validates the AntListen media type instance.
func (mt *AntListen) Validate() (err error) {
	if mt.Rid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rid"))
	}
	if mt.Etype == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "etype"))
	}
	if mt.Hookurl == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "hookurl"))
	}
	if mt.From == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "from"))
	}
	return
}

// DecodeAntListen decodes the AntListen instance encoded in resp body.
func (c *Client) DecodeAntListen(resp *http.Response) (*AntListen, error) {
	var decoded AntListen
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AntListenList 事件监听列表 (default view)
//
// Identifier: vnd.ant.listen.list+json; view=default
type AntListenList struct {
	// 事件类型
	List []*AntListen `form:"list,omitempty" json:"list,omitempty" xml:"list,omitempty"`
	// 总数量
	Total int `form:"total" json:"total" xml:"total"`
}

// Validate validates the AntListenList media type instance.
func (mt *AntListenList) Validate() (err error) {
	for _, e := range mt.List {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeAntListenList decodes the AntListenList instance encoded in resp body.
func (c *Client) DecodeAntListenList(resp *http.Response) (*AntListenList, error) {
	var decoded AntListenList
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AntRegResult 注册事件监听成功 (default view)
//
// Identifier: vnd.ant.reg.result+json; view=default
type AntRegResult struct {
	// 成功标识
	OK bool `form:"ok" json:"ok" xml:"ok"`
}

// AntRegResultFailed 注册事件监听成功 (failed view)
//
// Identifier: vnd.ant.reg.result+json; view=failed
type AntRegResultFailed struct {
	// 如果ok=false,失败原因
	Msg *string `form:"msg,omitempty" json:"msg,omitempty" xml:"msg,omitempty"`
	// 成功标识
	OK bool `form:"ok" json:"ok" xml:"ok"`
}

// DecodeAntRegResult decodes the AntRegResult instance encoded in resp body.
func (c *Client) DecodeAntRegResult(resp *http.Response) (*AntRegResult, error) {
	var decoded AntRegResult
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeAntRegResultFailed decodes the AntRegResultFailed instance encoded in resp body.
func (c *Client) DecodeAntRegResultFailed(resp *http.Response) (*AntRegResultFailed, error) {
	var decoded AntRegResultFailed
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AntResult 创建事件成功返回 (default view)
//
// Identifier: vnd.ant.result+json; view=default
type AntResult struct {
	// 事件唯一标识
	Eid *string `form:"eid,omitempty" json:"eid,omitempty" xml:"eid,omitempty"`
}

// DecodeAntResult decodes the AntResult instance encoded in resp body.
func (c *Client) DecodeAntResult(resp *http.Response) (*AntResult, error) {
	var decoded AntResult
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
