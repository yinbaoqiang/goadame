package engine

import "time"

var defaultEnginer EventEnginer

func init() {
	defaultStore := &showStore{}

	defaultEnginer = CreateEventEnginer(10*time.Second, defaultStore, defaultStore)
}

// SetEventStorer 设置事件存储
func SetEventStorer(s EventStorer) {
	defaultEnginer.SetEventStorer(s)
}

// SetListenerStore 设置监听存储
func SetListenerStore(s ListenerStore) {
	defaultEnginer.SetListenerStore(s)
}

// Put 向默认引擎添加事件
func Put(ei Event) error {
	return defaultEnginer.Put(ei)
}

// Start 启动默认引擎
func Start() error{
	return defaultEnginer.Start()
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
