package engine

import "time"

var defaultEnginer EventEnginer

func init() {
	defaultStore := &showStore{}
	opt := Option{
		TimeOut: 10 * time.Second,
		Estore:  defaultStore,
		Lstore:  defaultStore,
		TryCnt:  3,
	}
	defaultEnginer = CreateEventEnginer(opt)
}

// SetTryCnt 设置默认引擎的错误重试次数
func SetTryCnt(c int) {
	e, ok := defaultEnginer.(*eventEngine)
	if ok {
		e.tryCnt = c
	}
}

// SetTimeOut 设置默认引擎的超时时间
// to 必须大于1秒
func SetTimeOut(to time.Duration) {
	e, ok := defaultEnginer.(*eventEngine)
	if ok {
		if to > time.Second {

			e.timeOut = to
		}
	}
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
func Start() error {
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
