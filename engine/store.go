package engine

import "time"

// Storer 数据存储
type Storer interface {
	// 存储钩子回调事件失败
	HookError(url string, ei Event, err error, start, end time.Time)
	// 存储钩子回调事件成功
	HookSuccess(url string, ei Event, start, end time.Time)
	StoreOne(ei Event)
}

var defaultStore Storer

// RegStorer 注册store驱动
func RegStorer(store Storer) {
	defaultStore = store
}
