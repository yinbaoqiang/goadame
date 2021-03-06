package engine

import (
	"container/list"
	"fmt"
	"sync"
)

// ListenManager 监听管理器
type ListenManager interface {
	Add(url, etype, action string)
	Remove(url, etype, action string)
	GetAll(etype, action string) hookURL
}

func createListenManager() ListenManager {
	return &listenManage{}
}

func createLelem(etype string) *lelem {
	a := make(hookURL, 0, 1)
	return &lelem{
		etype:  etype,
		action: make(map[string]*hookURL),
		all:    a,
	}
}

type lelem struct {
	etype  string
	all    hookURL
	action map[string]*hookURL
	lck    sync.RWMutex
}

// 通过行为找到注册的hook
func (e *lelem) getHooks(action string) (out *hookURL) {
	e.lck.RLock()
	defer e.lck.RUnlock()
	defer func() {
		if len(e.all) > 0 {
			*out = append(*out, (e.all)...)
		}
	}()

	h := make(hookURL, 0, 2)
	out = &h
	if action == "" {
		return
	}

	hu, ok := e.action[action]
	if !ok {
		//fmt.Println("nil")
		return
	}

	if len(*hu) > 0 {

		*out = append(*out, (*hu)...)
	}

	return
}
func (e *lelem) _getLckHooks(action string) *hookURL {
	e.lck.RLock()
	defer e.lck.RUnlock()
	hu, ok := e.action[action]
	if !ok {
		return nil
	}
	return hu

}

// 通过行为找到注册的hook
func (e *lelem) _getHooks(action string) *hookURL {
	hu, ok := e.action[action]
	if !ok {
		return nil
	}
	return hu
}

// 返回所有的hookURL
func (e *lelem) getAllHooks() hookURL {
	e.lck.RLock()
	defer e.lck.RUnlock()
	out := make(hookURL, 0, 2)
	for _, hu := range e.action {
		out = append(out, *hu...)
	}
	if e.all != nil {
		out = append(out, e.all...)
	}
	return out
}
func (e *lelem) removeAction(action string) {
	e.lck.Lock()
	defer e.lck.Unlock()
	delete(e.action, action)
}
func (e *lelem) remove(action, url string) {
	e.lck.Lock()
	defer e.lck.Unlock()
	hu := e._getHooks(action)
	if hu != nil {
		hu.remove(url)
	}
}

func (e *lelem) put(action, url string) {
	fmt.Printf("put1======新加入或修改:%s,%s,%s\n", e.etype, action, url)
	if e.all.exists(url) {
		return
	}
	if action == "" {
		e.lck.Lock()
		defer e.lck.Unlock()
		//e.all = append(e.all, createHook(url))
		e.all.add(url)
		fmt.Printf("e.all add:%#v\n", e.all[0])
		for _, v := range e.action {
			if v.exists(url) {
				v.remove(url)
			}
		}
		return
	}
	fmt.Printf("put2======新加入或修改:%s,%s,%s\n", e.etype, action, url)
	hu := e._getLckHooks(action)
	if hu == nil {
		fmt.Printf("put3======新加入或修改:%s,%s,%s\n", e.etype, action, url)
		e.lck.Lock()
		hu = e._getHooks(action)
		if hu == nil {
			h := make(hookURL, 0, 2)
			hu = &h
			e.action[action] = hu
		}
		hu.add(url)
		e.lck.Unlock()
		return
	}
	if hu.exists(url) {
		fmt.Printf("put4======新加入或修改:%s,%s,%s\n", e.etype, action, url)
		return
	}
	e.lck.Lock()
	defer e.lck.Unlock()
	if hu.exists(url) {
		fmt.Printf("put5======新加入或修改:%s,%s,%s\n", e.etype, action, url)
		return
	}
	*hu = append(*hu, createHook(url))

	fmt.Printf("======新加入或修改:%s,%s,%s\n", e.etype, action, url)
}

type hook struct {
	url       string
	waitQueue *list.List
	cond      *sync.Cond
}

func (h *hook) putWait() {

}
func (h *hook) put(handler func()) {
	h.cond.L.Lock()
	h.waitQueue.PushBack(handler)
	h.cond.Signal()
	h.cond.L.Unlock()
}

// 等待队列中弹出一个元素
func (h *hook) pop() func() {
	e := h.waitQueue.Front()
	if e == nil {
		return nil
	}
	handler := e.Value.(func())
	h.waitQueue.Remove(e)
	return handler
}
func (h *hook) run() {
	for {
		h.cond.L.Lock()
		if h.waitQueue.Len() == 0 {

			h.cond.Wait()
			h.cond.L.Unlock()
			continue
		}
		handler := h.pop()
		h.cond.L.Unlock()
		if handler != nil {
			handler()
		}
	}
}

var (
	hookMap = make(map[string]*hook)
	hmlck   sync.Mutex
)

func createHook(url string) *hook {
	hmlck.Lock()
	defer hmlck.Unlock()
	h, ok := hookMap[url]
	if ok {
		return h
	}
	l := &sync.Mutex{}
	h = &hook{
		url:       url,
		waitQueue: list.New(),
		cond:      sync.NewCond(l),
	}
	go h.run()
	hookMap[url] = h
	return h
}

type hookURL []*hook

func (u *hookURL) exists(url string) bool {
	l := len(*u)
	for i := 0; i < l; i++ {
		if (*u)[i].url == url {
			return true
		}
	}
	return false
}
func (u *hookURL) add(url string) {
	*u = append(*u, createHook(url))
}
func (u *hookURL) remove(url string) {
	l := len(*u)
	for i := 0; i < l; i++ {
		if (*u)[i].url == url {
			if i == 0 {
				*u = (*u)[1:]
				break
			} else if i+1 == l {
				*u = (*u)[0:i]
				break
			}
			n := (*u)[0:i]
			n = append(n, (*u)[i+1:l]...)
			*u = n
			break
		}
	}
}

type listenManage struct {
	lmap map[string]*lelem
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
		m.lmap = make(map[string]*lelem)
		m.append(etype, action, url)
		return
	}
	lm, ok := m.lmap[etype]
	if !ok {
		m.append(etype, action, url)
		return
	}
	lm.put(action, url)

}
func (m *listenManage) append(etype, action, url string) {
	u := make(hookURL, 0, 2)
	u.add(url)
	le := createLelem(etype)
	le.put(action, url)
	m.lmap[etype] = le
}
func (m *listenManage) Remove(url, etype, action string) {
	m.lock(func() {
		m.remove(url, etype, action)
	})
}
func (m *listenManage) remove(url, etype, action string) {
	u, ok := m.lmap[etype]
	if !ok {
		return
	}
	u.remove(action, url)
}

func (m *listenManage) GetAll(etype, action string) (hu hookURL) {
	m.rlock(func() {
		hu = m.getAll(etype, action)
	})
	return
}
func (m *listenManage) getAll(etype, action string) (hu hookURL) {
	u, ok := m.lmap[etype]
	if !ok {
		return
	}
	u1 := u.getHooks(action)
	if u1 != nil {
		return *u1
	}
	return
}
