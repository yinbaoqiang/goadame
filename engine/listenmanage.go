package engine

import "sync"

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
	return &lelem{
		etype:  etype,
		action: make(map[string]*hookURL),
	}
}

type lelem struct {
	etype  string
	all    hookURL
	action map[string]*hookURL
	lck    sync.RWMutex
}

// 通过行为找到注册的hook
func (e *lelem) getHooks(action string) *hookURL {
	if action == "" {
		return &e.all
	}
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
	if action == "" {
		return &e.all
	}

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
	hu := e.getHooks(action)
	if hu != nil {
		hu.remove(url)
	}
}
func (e *lelem) put(action, url string) {
	hu := e.getHooks(action)
	if hu == nil {
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
		return
	}
	e.lck.Lock()
	defer e.lck.Unlock()
	if hu.exists(url) {
		return
	}
	hu.add(url)

}

type hook string

type hookURL []hook

func (u *hookURL) exists(url string) bool {
	l := len(*u)
	for i := 0; i < l; i++ {
		if (*u)[i] == hook(url) {
			return true
		}
	}
	return false
}
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
	m.lmap[etype] = createLelem(etype)
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
