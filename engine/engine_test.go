package engine_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/yinbaoqiang/goadame/app"
	. "github.com/yinbaoqiang/goadame/engine"
)

const (
	hookError   = "HookError"
	hookSuccess = "HookSuccess"
	storeOne    = "StoreOne"
)

// Storer 数据存储
type testListenerStore struct {
}

func (s *testListenerStore) Watch(func(ctyp ChgType, lis app.AntListen)) {

}

// 获取所有的
func (s *testListenerStore) All() (res []*app.AntListen, err error) {
	return
}

// Storer 数据存储
type testStore struct {
	callback func(Event, string)
}

// 存储钩子回调事件失败
func (s *testStore) HookError(url string, ei Event, err error, start, end time.Time) {
	s.callback(ei, hookError)
}

// 存储钩子回调事件成功
func (s *testStore) HookSuccess(url string, ei Event, start, end time.Time) {
	s.callback(ei, hookSuccess)
}
func (s *testStore) SaveEvent(ei Event) {
	s.callback(ei, storeOne)
}

type testServer struct {
	gt        *testing.T
	defEvents map[string]Event
}

func (t *testServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	body := req.Body
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		resp.WriteHeader(500)
		return
	}
	var evt Event
	err = json.Unmarshal(data, &evt)
	if err != nil {
		resp.WriteHeader(400)
		return
	}
	old, ok := t.defEvents[evt.Eid]
	if !ok {
		t.gt.Errorf("收到错误的事件:%s\n", string(data))
		return
	}

	if old.Eid != evt.Eid {
		t.gt.Errorf("收到错误的事件:%s\n", string(data))
		return
	}
	t.gt.Logf("data:%s", string(data))
	resp.WriteHeader(200)
}
func (t *testServer) Start() {

	go http.ListenAndServe(":12345", t)
}

func TestMain(m *testing.M) {
	//(&testServer{}).Start()
	os.Exit(m.Run())
}

func TestOK(t *testing.T) {

	var (
		defEvents = map[string]Event{
			"10920192": Event{Eid: "10920192", Etype: "typ_test", Action: "action_test"},
			"10920193": Event{Eid: "10920193", Etype: "typ_test", Action: "action_test"},
			"10920194": Event{Eid: "10920194", Etype: "typ_test", Action: "action_test2"},
		}
	)
	enum := len(defEvents)
	server := httptest.NewServer(&testServer{gt: t, defEvents: defEvents})
	defer server.Close()
	storecnt := 0
	opt := Option{
		TimeOut: 3 * time.Second,
		Estore: &testStore{
			callback: func(ei Event, opresult string) {
				switch opresult {
				case hookError:
					t.Errorf("该处不应该失败")
					delete(defEvents, ei.Eid)
				case hookSuccess:
					delete(defEvents, ei.Eid)
				case storeOne:
					storecnt++
				}
			},
		},
		Lstore: &testListenerStore{},
	}
	engine := CreateEventEnginer(opt)
	engine.ListenManager().Add(server.URL, "typ_test", "action_test")
	engine.Start()
	for _, e := range defEvents {
		engine.Put(e)
	}

	engine.Stop()

	err := engine.Put(Event{Eid: "10920192", Etype: "typ_test", Action: "action_test"})
	if err == nil {
		t.Errorf("添加事件应该失败")
	}

	if len(defEvents) != 1 {
		t.Errorf("事件处理失败，还有剩余事件没处理:%d", len(defEvents))
	}
	if storecnt != enum {
		t.Errorf("事件存储数量错误:%d/%d", storecnt, enum)
	}
}

func TestOK2(t *testing.T) {

	var (
		defEvents = map[string]Event{
			"10920192": Event{Eid: "10920192", Etype: "typ_test", Action: "action_test"},
			"10920193": Event{Eid: "10920193", Etype: "typ_test", Action: "action_test"},
			"10920194": Event{Eid: "10920194", Etype: "typ_test", Action: "action_test2"},
		}
	)
	enum := len(defEvents)
	server := httptest.NewServer(&testServer{gt: t, defEvents: defEvents})
	defer server.Close()
	storecnt := 0
	opt := Option{
		TimeOut: 3 * time.Second,
		Estore: &testStore{
			callback: func(ei Event, opresult string) {
				switch opresult {
				case hookError:
					t.Errorf("该处不应该失败")
					delete(defEvents, ei.Eid)
				case hookSuccess:
					delete(defEvents, ei.Eid)
				case storeOne:
					storecnt++
				}
			},
		},
		Lstore: &testListenerStore{},
	}
	engine := CreateEventEnginer(opt)
	engine.ListenManager().Add(server.URL, "typ_test", "")
	engine.Start()
	for _, e := range defEvents {
		err := engine.Put(e)
		if err != nil {
			t.Errorf("添加事件失败:%v\n", err)
		}
	}
	engine.Stop()
	if len(defEvents) != 0 {
		t.Errorf("事件处理失败，还有剩余事件没处理:%d", len(defEvents))
	}
	if storecnt != enum {
		t.Errorf("事件存储数量错误:%d/%d", storecnt, enum)
	}
}

func TestError500(t *testing.T) {

	var (
		defEvents = map[string]Event{
			"10920192": Event{Eid: "10920192", Etype: "typ_test", Action: "action_test"},
			"10920193": Event{Eid: "10920193", Etype: "typ_test", Action: "action_test"},
			"10920194": Event{Eid: "10920194", Etype: "typ_test", Action: "action_test2"},
		}
	)
	enum := len(defEvents)
	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		body := req.Body
		defer body.Close()
		data, err := ioutil.ReadAll(body)
		if err != nil {
			resp.WriteHeader(500)
			return
		}
		var evt Event
		err = json.Unmarshal(data, &evt)
		if err != nil {
			resp.WriteHeader(400)
			return
		}
		old, ok := defEvents[evt.Eid]
		if !ok {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}

		if old.Eid != evt.Eid {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}
		delete(defEvents, old.Eid)
		t.Logf("data:%s", string(data))
		resp.WriteHeader(500)
		resp.Write([]byte(`{"msg":"测试500错误"}`))
	}))
	defer server.Close()
	storecnt := 0

	opt := Option{
		TimeOut: 3 * time.Second,
		Estore: &testStore{
			callback: func(ei Event, opresult string) {
				switch opresult {
				case hookError:
					t.Errorf("该处不应该失败")
					delete(defEvents, ei.Eid)
				case hookSuccess:
					delete(defEvents, ei.Eid)
				case storeOne:
					storecnt++
				}
			},
		},
		Lstore: &testListenerStore{},
	}
	engine := CreateEventEnginer(opt)

	engine.ListenManager().Add(server.URL, "typ_test", "")
	engine.Start()
	for _, e := range defEvents {
		err := engine.Put(e)
		if err != nil {
			t.Errorf("添加事件失败:%v\n", err)
		}
	}
	engine.Stop()
	if len(defEvents) != 0 {
		t.Errorf("事件处理失败，还有剩余事件没处理:%d", len(defEvents))
	}
	if storecnt != enum {
		t.Errorf("事件存储数量错误:%d/%d", storecnt, enum)
	}
}

func TestError400(t *testing.T) {

	var (
		defEvents = map[string]Event{
			"10920192": Event{Eid: "10920192", Etype: "typ_test", Action: "action_test"},
			"10920193": Event{Eid: "10920193", Etype: "typ_test", Action: "action_test"},
			"10920194": Event{Eid: "10920194", Etype: "typ_test", Action: "action_test2"},
		}
	)
	enum := len(defEvents)
	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		body := req.Body
		defer body.Close()
		data, err := ioutil.ReadAll(body)
		if err != nil {
			resp.WriteHeader(500)
			return
		}
		var evt Event
		err = json.Unmarshal(data, &evt)
		if err != nil {
			resp.WriteHeader(400)
			return
		}
		old, ok := defEvents[evt.Eid]
		if !ok {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}

		if old.Eid != evt.Eid {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}
		delete(defEvents, old.Eid)
		t.Logf("data:%s", string(data))
		resp.WriteHeader(400)
		resp.Write([]byte(`{"msg":"测试400错误"}`))
	}))
	defer server.Close()
	storecnt := 0
	opt := Option{
		TimeOut: 3 * time.Second,
		Estore: &testStore{
			callback: func(ei Event, opresult string) {
				switch opresult {
				case hookError:
					t.Errorf("该处不应该失败")
					delete(defEvents, ei.Eid)
				case hookSuccess:
					delete(defEvents, ei.Eid)
				case storeOne:
					storecnt++
				}
			},
		},
		Lstore: &testListenerStore{},
	}
	engine := CreateEventEnginer(opt)
	engine.ListenManager().Add(server.URL, "typ_test", "")
	engine.Start()
	for _, e := range defEvents {
		err := engine.Put(e)
		if err != nil {
			t.Errorf("添加事件失败:%v\n", err)
		}
	}
	engine.Stop()
	if len(defEvents) != 0 {
		t.Errorf("事件处理失败，还有剩余事件没处理:%d", len(defEvents))
	}
	if storecnt != enum {
		t.Errorf("事件存储数量错误:%d/%d", storecnt, enum)
	}
}
func TestErrorTimeOut(t *testing.T) {

	var (
		defEvents = map[string]Event{
			"10920192": Event{Eid: "10920192", Etype: "typ_test", Action: "action_test"},
			"10920193": Event{Eid: "10920193", Etype: "typ_test", Action: "action_test"},
			"10920194": Event{Eid: "10920194", Etype: "typ_test", Action: "action_test2"},
		}
	)
	enum := len(defEvents)
	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		body := req.Body
		defer body.Close()
		data, err := ioutil.ReadAll(body)
		if err != nil {
			resp.WriteHeader(500)
			return
		}
		var evt Event
		err = json.Unmarshal(data, &evt)
		if err != nil {
			resp.WriteHeader(400)
			return
		}
		old, ok := defEvents[evt.Eid]
		if !ok {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}

		if old.Eid != evt.Eid {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}
		delete(defEvents, old.Eid)
		t.Logf("data:%s", string(data))
		time.Sleep(3 * time.Second)
		resp.WriteHeader(200)
	}))
	defer server.Close()
	storecnt := int64(0)

	opt := Option{
		TimeOut: 100 * time.Millisecond,
		Estore: &testStore{
			callback: func(ei Event, opresult string) {
				switch opresult {
				case hookError:
					t.Errorf("该处不应该失败")
					delete(defEvents, ei.Eid)
				case hookSuccess:
					delete(defEvents, ei.Eid)
				case storeOne:
					atomic.AddInt64(&storecnt, 1)
				}
			},
		},
		Lstore: &testListenerStore{},
	}
	engine := CreateEventEnginer(opt)
	engine.ListenManager().Add(server.URL, "typ_test", "")
	engine.Start()
	for _, e := range defEvents {
		err := engine.Put(e)

		if err != nil {
			t.Errorf("添加事件失败:%v\n", err)
		}
	}
	engine.Stop()
	if len(defEvents) != 0 {
		t.Errorf("事件处理失败，还有剩余事件没处理:%d", len(defEvents))
	}
	if storecnt != int64(enum) {
		t.Errorf("事件存储数量错误:%d/%d", storecnt, enum)
	}
}

func BenchmarkPutEvent(b *testing.B) {

	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		resp.WriteHeader(200)
	}))
	defer server.Close()
	storecnt := int64(0)
	errorcnt := 0
	opt := Option{
		TimeOut: 3000 * time.Millisecond,
		Estore: &testStore{
			callback: func(ei Event, opresult string) {
				switch opresult {
				case hookError:
					errorcnt++
				case hookSuccess:

				case storeOne:
					atomic.AddInt64(&storecnt, 1)
				}
			},
		},
		Lstore: &testListenerStore{},
	}
	engine := CreateEventEnginer(opt)
	engine.ListenManager().Add(server.URL, "typ_test", "")
	engine.Start()
	for i := 0; i < b.N; i++ {
		engine.Put(Event{Eid: "id_" + strconv.Itoa(i), Etype: "typ_test", Action: "action_test"})
	}
	engine.Stop()
	if storecnt != int64(b.N) {
		b.Errorf("事件存储数量错误:%d/%d", storecnt, b.N)
	}
}
