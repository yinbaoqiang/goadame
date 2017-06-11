package engine

import (
	"bytes"
	"context"
	"net/http"
	"sync"
	"time"
)

// EventInfo 事件
type EventInfo struct {
	Eid     string
	Action  string
	Etype   string
	From    string
	OccTime time.Time
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
	Add(url, action, etype string)
	Remove(url, action, etype string)
	GetAll(action, etype string) hookURL
}

type lelem struct {
	action string
	etype  string
}
type hook string

func (h hook) call(ctx context.Context, etype, action string, data []byte) error {
	req, err := http.NewRequest("PUT", string(h), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
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
func (m *listenManage) Add(url, action, etype string) {
	m.lock(func() {
		m.add(url, action, etype)
	})
}
func (m *listenManage) add(url, action, etype string) {
	if m.lmap == nil {
		m.lmap = make(map[lelem]*hookURL)
		m.append(action, etype, url)
		return
	}
	u, ok := m.lmap[lelem{action, etype}]
	if !ok {
		m.append(action, etype, url)
		return
	}
	*u = append(*u, url)

}
func (m *listenManage) append(action, etype, url string) {
	u := make(hookURL, 0, 2)
	u.add(url)
	m.lmap[lelem{action, etype}] = &u
}
func (m *listenManage) Remove(url, action, etype string) {
	m.lock(func() {
		m.remove(url, action, etype)
	})
}
func (m *listenManage) remove(url, action, etype string) {
	u, ok := m.lmap[lelem{action, etype}]
	if !ok {
		return
	}
	u.remove(url)
}

func (m *listenManage) GetAll(action, etype string) (hu hookURL) {
	m.rlock(func() {
		hu = m.getAll(action, etype)
	})
	return
}
func (m *listenManage) getAll(action, etype string) (hu hookURL) {
	u, ok := m.lmap[lelem{action, etype}]
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
}

func (e *eventEngine) ListenManager() ListenManager {
	return e.lm
}

func (e *eventEngine) Receive(ei EventInfo) {

}
