package engine

import "time"

var defaultEnginer EventEnginer

func init() {
	if defaultStore == nil {
		defaultStore = &showStore{}
	}
	defaultEnginer = CreateEventEnginer(10 * time.Second)
}

// Put 向默认引擎添加事件
func Put(ei Event) error {
	return defaultEnginer.Put(ei)
}

// Start 启动默认引擎
func Start() {
	defaultEnginer.Start()
}

// Stop 停止默认引擎
func Stop() {
	defaultEnginer.Stop()
}

// DefaultEnginer 默认引擎
func DefaultEnginer() EventEnginer {
	return defaultEnginer
}

// AddListen 向默认引擎新增一个监听
func AddListen(url, etype, action string) {
	defaultEnginer.ListenManager().Add(url, etype, action)
}

// RemoveListen 从默认引擎移除一个监听
func RemoveListen(url, etype, action string) {
	defaultEnginer.ListenManager().Remove(url, etype, action)
}
