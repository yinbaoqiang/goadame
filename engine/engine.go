package engine

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// HookResult 事件
type HookResult struct {
	Msg string `json:"msg"`
}

// EventInfo 事件
type EventInfo struct {
	Eid     string    `json:"eid"`
	Action  string    `json:"action"`
	Etype   string    `json:"etype"`
	From    string    `json:"from"`
	OccTime time.Time `json:"occTime"`
}

// Event 事件
type Event interface {
	// GetEventInfo 获取事件信息
	GetEventInfo() EventInfo
	// GetData 获取数据
	GetData() map[string]interface{}
}

// ListenManager 监听管理器
type ListenManager interface {
	Add(url, etype, action string)
	Remove(url, etype, action string)
	GetAll(etype, action string) hookURL
}

type lelem struct {
	etype  string
	action string
}
type hook string

func (h hook) call(ctx context.Context, etype, action string, data []byte) error {
	req, err := http.NewRequest("PUT", string(h), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	http.DefaultClient.Do()
	return req, nil
	return nil
}

type hookURL []hook

func (u *hookURL) add(url string) {
	*u = append(*u, hook(url))
}
func (u *hookURL) remove(url string) {
	l := len(*u)
	for i := 0; i < l; i++ {
		if (*u)[i] == hook(url) {
			n := (*u)[0:i]
			if i+1 != l {
				n = append(n, (*u)[i+1:l]...)
			}
			u = &n

		}
	}
}

func (u hookURL) trigger(ctx context.Context, etype, action string, data []byte) {
	for _, h := range u {
		go func(h hook) {
			cx, c := context.WithCancel(ctx)
			defer c()
			h.call(cx, etype, action, data)
		}(h)

	}
}

type listenManage struct {
	lmap map[lelem]*hookURL
	lck  sync.RWMutex
}

func (m *listenManage) lock(handler func()) {
	m.lck.Lock()
	defer m.lck.Unlock()
	handler()
}
func (m *listenManage) rlock(handler func()) {
	m.lck.RLock()
	defer m.lck.RUnlock()
	handler()
}
func (m *listenManage) Add(url, etype, action string) {
	m.lock(func() {
		m.add(url, etype, action)
	})
}
func (m *listenManage) add(url, etype, action string) {
	if m.lmap == nil {
		m.lmap = make(map[lelem]*hookURL)
		m.append(etype, action, url)
		return
	}
	u, ok := m.lmap[lelem{etype, action}]
	if !ok {
		m.append(etype, action, url)
		return
	}
	*u = append(*u, hook(url))

}
func (m *listenManage) append(etype, action, url string) {
	u := make(hookURL, 0, 2)
	u.add(url)
	m.lmap[lelem{etype, action}] = &u
}
func (m *listenManage) Remove(url, etype, action string) {
	m.lock(func() {
		m.remove(url, etype, action)
	})
}
func (m *listenManage) remove(url, etype, action string) {
	u, ok := m.lmap[lelem{etype, action}]
	if !ok {
		return
	}
	u.remove(url)
}

func (m *listenManage) GetAll(etype, action string) (hu hookURL) {
	m.rlock(func() {
		hu = m.getAll(etype, action)
	})
	return
}
func (m *listenManage) getAll(etype, action string) (hu hookURL) {
	u, ok := m.lmap[lelem{etype, action}]
	if !ok {
		return
	}
	hu = *u
	return
}

// EventEnginer 事件引擎
type EventEnginer interface {
	// 获取监听管理器
	ListenManager() ListenManager

	Receive(ei EventInfo)
}

type eventEngine struct {
	lm *listenManage

	client *eventClient

	echan chan Event
}

func (e *eventEngine) ListenManager() ListenManager {
	return e.lm
}

func (e *eventEngine) Put(ei Event) {
	e.echan <- ei
}
func (e *eventEngine) work() {

	for {

	}
}
func (e *eventEngine) one(ei Event) {
	// 持久话事件
	go e.storeOne(ei)
	// 查询事件监听

	hu := e.lm.GetAll(ei.GetEventInfo().Etype, ei.GetEventInfo().Action)

	for _, h := range hu {
		e.hook(string(h), ei)
	}
}
func (e *eventEngine) hook(url string, ei Event) {
	// 记录开始时间
	start := time.Now()
	rchan := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		defer close(rchan)
		res, err := e.client.SendEvent(ctx, url, ei)
		if err != nil {
			rchan <- err
			return
		}
		// 处理业务逻辑
		if res.StatusCode == http.StatusOK {
			return
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			rchan <- err
			return
		}
		var r HookResult
		err = json.Unmarshal(data, &r)
		if err != nil {
			rchan <- err
			return
		}
		rchan <- fmt.Errorf("请求%s失败:%s", url, r.Msg)
		return

	}()

	select {
	case <-ctx.Done():
		if ctx.Err() != nil {
			// 超时失败
			e.hookError(url, ei, ctx.Err(), start)
			return

		}
	case err := <-rchan:
		if err != nil {
			e.hookError(url, ei, err, start)
			return
		}
	}

}

// 存储钩子回调事件失败
func (e *eventEngine) hookError(url string, ei Event, err error, start time.Time) {

}

// 存储钩子回调事件成功
func (e *eventEngine) hookSuccess(url string, ei Event, start time.Time) {

}
func (e *eventEngine) storeOne(ei Event) {

}
