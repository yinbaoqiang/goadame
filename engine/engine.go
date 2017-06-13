package engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (e *eventEngine) Put(ei Event) error {
	select {
	case e.echan <- ei:
	default:
		return errors.New("事件引擎已经关闭或未开启")
	}
	return nil
}
func (e *eventEngine) work() {

	for ei := range e.echan {
		if ei != nil {
			e.one(ei)
		}

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
