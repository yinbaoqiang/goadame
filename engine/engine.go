package engine

import (
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

// EventEnginer 事件引擎
type EventEnginer interface {
	// 获取监听管理器
	ListenManager() ListenManager

	Put(ei Event) error

	Start()
	Stop()
}

// CreateEventEnginer 创建事件引擎
func CreateEventEnginer(timeOut time.Duration, store ...Storer) EventEnginer {
	s := defaultStore
	if len(store) == 1 {
		s = store[0]
	}
	if s == nil {
		panic("没有注册store ,请使用RegStorer 注册.或传入store参数")
	}
	if timeOut == 0 {
		timeOut = 3 * time.Second
	}
	return &eventEngine{
		lm:      createListenManager(),
		store:   s,
		timeOut: timeOut,
	}
}

type eventEngine struct {
	lm ListenManager

	echan chan Event
	wg    sync.WaitGroup

	store   Storer
	timeOut time.Duration
}

func (e *eventEngine) ListenManager() ListenManager {
	return e.lm
}

func (e *eventEngine) Put(ei Event) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("事件引擎已经关闭或未开启:%v", e)
		}
	}()
	e.echan <- ei
	e.wg.Add(1)
	return nil
}
func (e *eventEngine) Start() {
	e.echan = make(chan Event, 10)
	go e.work()
}
func (e *eventEngine) Stop() {
	close(e.echan)
	e.wg.Wait()
}
func (e *eventEngine) work() {

	for ei := range e.echan {

		if ei.Etype != "" {
			e.one(ei)
		}

	}
}
func (e *eventEngine) one(ei Event) {
	defer e.wg.Done()
	// 持久化事件
	e.wg.Add(1)
	go e.storeOne(ei)
	// 查询事件监听
	hu := e.lm.GetAll(ei.Etype, ei.Action)

	for _, h := range hu {
		e.wg.Add(1)
		// 加入调用队列
		//	fmt.Printf("%s:%s=>%s\n", ei.Etype, ei.Action, h.url)
		nh := h
		nh.put(func() {
			defer e.wg.Done()
			// 执行钩子回调
			//	fmt.Printf("执行钩子回调:%s:%s=>%s\n", ei.Etype, ei.Action, nh.url)
			e.hook(nh.url, ei)
		})

	}
}

// 向外发送事件
// ctx 上下文
// url 发送地址
// ei 事件
func (e *eventEngine) _sendEvent(ctx context.Context, url string, ei Event) error {
	res, err := newEventClient().SendEvent(ctx, url, ei)
	if err != nil {

		return err
	}
	// 处理业务逻辑
	if res.StatusCode == http.StatusOK {
		return nil
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var r HookResult
	err = json.Unmarshal(data, &r)
	if err != nil {
		return err
	}
	return fmt.Errorf("请求%s失败:%s", url, r.Msg)
}

// 调用监听钩子
func (e *eventEngine) hook(url string, ei Event) {

	// 记录开始时间
	start := time.Now()
	rchan := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), e.timeOut)
	defer cancel()
	go func() {
		defer close(rchan)
		err := e._sendEvent(ctx, url, ei)
		if err != nil {
			rchan <- err
		}
		return

	}()

	select {
	case <-ctx.Done():
		if ctx.Err() != nil {
			// 超时失败
			e.hookError(url, ei, ctx.Err(), start, time.Now())
			return

		}
	case err := <-rchan:
		if err != nil {
			e.hookError(url, ei, err, start, time.Now())
			return
		}
		e.hookSuccess(url, ei, start, time.Now())
	}

}

// 存储钩子回调事件失败
func (e *eventEngine) hookError(url string, ei Event, err error, start, end time.Time) {
	e.store.HookError(url, ei, err, start, end)
}

// 存储钩子回调事件成功
func (e *eventEngine) hookSuccess(url string, ei Event, start, end time.Time) {
	e.store.HookSuccess(url, ei, start, end)
}
func (e *eventEngine) storeOne(ei Event) {
	defer e.wg.Done()
	e.store.SaveEvent(ei)
}
