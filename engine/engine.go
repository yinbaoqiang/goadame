package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/yinbaoqiang/goadame/app"
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
	Start() error
	Stop()
	SetEventStorer(s EventStorer)
	SetListenerStore(s ListenerStore)
}

// Option 创建引擎参数
type Option struct {
	TimeOut time.Duration
	Estore  EventStorer
	Lstore  ListenerStore
	TryCnt  int
}

// CreateEventEnginer 创建事件引擎
func CreateEventEnginer(opt Option) EventEnginer {

	if opt.Estore == nil {
		opt.Estore = defaultEventStore
	}
	if opt.Lstore == nil {
		opt.Lstore = defaultEventStore
	}
	if opt.TimeOut == 0 {
		opt.TimeOut = 3 * time.Second
	}
	if opt.TryCnt < 1 {
		opt.TryCnt = 1
	}
	return &eventEngine{
		lstore:  opt.Lstore,
		store:   opt.Estore,
		timeOut: opt.TimeOut,
		tryCnt:  opt.TryCnt,
	}
}

type eventEngine struct {
	lm      ListenManager
	echan   chan Event
	wg      sync.WaitGroup
	lstore  ListenerStore
	store   EventStorer
	timeOut time.Duration
	tryCnt  int
}

func (e *eventEngine) SetEventStorer(s EventStorer) {
	e.store = s
}
func (e *eventEngine) SetListenerStore(s ListenerStore) {
	e.lstore = s
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
func (e *eventEngine) initListenManager() error {
	e.lm = createListenManager()
	ls, err := e.lstore.All()
	if err != nil {
		return err
	}
	for _, a := range ls {
		e.addListen(*a)
	}
	return nil
}
func (e *eventEngine) addListen(lis app.AntListen) {
	action := ""
	if lis.Action != nil {
		action = *lis.Action
	}
	e.lm.Add(lis.Hookurl, lis.Etype, action)
}
func (e *eventEngine) removeListen(lis app.AntListen) {
	action := ""
	if lis.Action != nil {
		action = *lis.Action
	}
	e.lm.Remove(lis.Hookurl, lis.Etype, action)
}
func (e *eventEngine) listenWatch() {
	e.lstore.Watch(func(ctyp ChgType, lis app.AntListen) {
		fmt.Printf("[%d]%v:%v=>%v\n", ctyp, lis.Etype, lis.Action, lis.Hookurl)
		switch ctyp {
		case ChgTypeAdd:
			e.addListen(lis)
		case ChgTypeRemove:
			e.removeListen(lis)
		case ChgTypeUpdate:
			e.addListen(lis)
		default:

		}
	})
}
func (e *eventEngine) Start() error {
	if e.tryCnt < 1 {
		e.tryCnt = 1
	}
	e.echan = make(chan Event, 10)
	err := e.initListenManager()
	if err != nil {
		return err
	}
	e.listenWatch()
	go e.work()
	return nil
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
		fmt.Printf("%s:%s=>%s\n", ei.Etype, ei.Action, h.url)
		nh := h
		nh.put(func() {
			defer e.wg.Done()
			// 执行钩子回调
			fmt.Printf("执行钩子回调:%s:%s=>%s\n", ei.Etype, ei.Action, nh.url)
			for i := 0; i < e.tryCnt; i++ {

				err := e.hook(nh.url, ei)
				if err == nil {
					return
				}
			}

		})

	}
}

// 向外发送事件
// ctx 上下文
// url 发送地址
// ei 事件
func (e *eventEngine) _sendEvent(ctx context.Context, url string, ei Event) error {
	res, err := newEventClient(e.timeOut).SendEvent(ctx, url, ei)
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
func (e *eventEngine) hook(url string, ei Event) error {

	// 记录开始时间
	start := time.Now()
	rchan := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), e.timeOut+5*time.Second)
	defer cancel()
	go func() {
		defer close(rchan)
		err := e._sendEvent(ctx, url, ei)
		if err != nil {
			rchan <- err
		}

	}()

	select {
	case <-ctx.Done():
		if ctx.Err() != nil {
			// 超时失败
			fmt.Printf("%s:%s=>%s 执行超时\n", ei.Etype, ei.Action, url)
			e.hookError(url, ei, ctx.Err(), start, time.Now())
			return ctx.Err()

		}
	case err := <-rchan:

		if err != nil {
			fmt.Printf("%s:%s=>%s 执行失败\n", ei.Etype, ei.Action, url)
			e.hookError(url, ei, err, start, time.Now())
			return err
		}
		fmt.Printf("%s:%s=>%s 执行成功\n", ei.Etype, ei.Action, url)
		e.hookSuccess(url, ei, start, time.Now())
	}
	return nil

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
