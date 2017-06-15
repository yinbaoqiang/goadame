package engine

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

const (
	hookError   = "HookError"
	hookSuccess = "HookSuccess"
	storeOne    = "StoreOne"
)

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
func (s *testStore) StoreOne(ei Event) {
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
	old, ok := t.defEvents[evt.Info.Eid]
	if !ok {
		t.gt.Errorf("收到错误的事件:%s\n", string(data))
		return
	}

	if old.Info != evt.Info {
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
			"10920192": Event{Info: EventInfo{Eid: "10920192", Etype: "typ_test", Action: "action_test"}},
			"10920193": Event{Info: EventInfo{Eid: "10920193", Etype: "typ_test", Action: "action_test"}},
			"10920194": Event{Info: EventInfo{Eid: "10920194", Etype: "typ_test", Action: "action_test2"}},
		}
	)
	enum := len(defEvents)
	server := httptest.NewServer(&testServer{gt: t, defEvents: defEvents})
	defer server.Close()
	storecnt := 0
	engine := CreateEventEnginer(3*time.Second, &testStore{
		callback: func(ei Event, opresult string) {
			switch opresult {
			case hookError:
				t.Errorf("该处不应该失败")
				delete(defEvents, ei.Info.Eid)
			case hookSuccess:
				delete(defEvents, ei.Info.Eid)
			case storeOne:
				storecnt++
			}
		},
	})
	engine.ListenManager().Add(server.URL, "typ_test", "action_test")
	engine.Start()
	for _, e := range defEvents {
		engine.Put(e)
	}

	engine.Stop()

	err := engine.Put(Event{Info: EventInfo{Eid: "10920192", Etype: "typ_test", Action: "action_test"}})
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
			"10920192": Event{Info: EventInfo{Eid: "10920192", Etype: "typ_test", Action: "action_test"}},
			"10920193": Event{Info: EventInfo{Eid: "10920193", Etype: "typ_test", Action: "action_test"}},
			"10920194": Event{Info: EventInfo{Eid: "10920194", Etype: "typ_test", Action: "action_test2"}},
		}
	)
	enum := len(defEvents)
	server := httptest.NewServer(&testServer{gt: t, defEvents: defEvents})
	defer server.Close()
	storecnt := 0
	engine := CreateEventEnginer(3*time.Second, &testStore{
		callback: func(ei Event, opresult string) {
			switch opresult {
			case hookError:
				t.Errorf("该处不应该失败")
				delete(defEvents, ei.Info.Eid)
			case hookSuccess:
				delete(defEvents, ei.Info.Eid)
			case storeOne:
				storecnt++
			}
		},
	})
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
			"10920192": Event{Info: EventInfo{Eid: "10920192", Etype: "typ_test", Action: "action_test"}},
			"10920193": Event{Info: EventInfo{Eid: "10920193", Etype: "typ_test", Action: "action_test"}},
			"10920194": Event{Info: EventInfo{Eid: "10920194", Etype: "typ_test", Action: "action_test2"}},
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
		old, ok := defEvents[evt.Info.Eid]
		if !ok {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}

		if old.Info != evt.Info {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}
		delete(defEvents, old.Info.Eid)
		t.Logf("data:%s", string(data))
		resp.WriteHeader(500)
		resp.Write([]byte(`{"msg":"测试500错误"}`))
	}))
	defer server.Close()
	storecnt := 0
	engine := CreateEventEnginer(3*time.Second, &testStore{
		callback: func(ei Event, opresult string) {
			switch opresult {
			case hookError:
				delete(defEvents, ei.Info.Eid)
			case hookSuccess:
				t.Errorf("该处不应该成功")
				delete(defEvents, ei.Info.Eid)
			case storeOne:
				storecnt++
			}
		},
	})
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
			"10920192": Event{Info: EventInfo{Eid: "10920192", Etype: "typ_test", Action: "action_test"}},
			"10920193": Event{Info: EventInfo{Eid: "10920193", Etype: "typ_test", Action: "action_test"}},
			"10920194": Event{Info: EventInfo{Eid: "10920194", Etype: "typ_test", Action: "action_test2"}},
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
		old, ok := defEvents[evt.Info.Eid]
		if !ok {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}

		if old.Info != evt.Info {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}
		delete(defEvents, old.Info.Eid)
		t.Logf("data:%s", string(data))
		resp.WriteHeader(400)
		resp.Write([]byte(`{"msg":"测试400错误"}`))
	}))
	defer server.Close()
	storecnt := 0
	engine := CreateEventEnginer(3*time.Second, &testStore{
		callback: func(ei Event, opresult string) {
			switch opresult {
			case hookError:
				delete(defEvents, ei.Info.Eid)
			case hookSuccess:
				t.Errorf("该处不应该成功")
				delete(defEvents, ei.Info.Eid)
			case storeOne:
				storecnt++
			}
		},
	})
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
			"10920192": Event{Info: EventInfo{Eid: "10920192", Etype: "typ_test", Action: "action_test"}},
			"10920193": Event{Info: EventInfo{Eid: "10920193", Etype: "typ_test", Action: "action_test"}},
			"10920194": Event{Info: EventInfo{Eid: "10920194", Etype: "typ_test", Action: "action_test2"}},
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
		old, ok := defEvents[evt.Info.Eid]
		if !ok {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}

		if old.Info != evt.Info {
			t.Errorf("收到错误的事件:%s\n", string(data))
			return
		}
		delete(defEvents, old.Info.Eid)
		t.Logf("data:%s", string(data))
		time.Sleep(3 * time.Second)
		resp.WriteHeader(200)
	}))
	defer server.Close()
	storecnt := 0
	engine := CreateEventEnginer(100*time.Millisecond, &testStore{
		callback: func(ei Event, opresult string) {
			switch opresult {
			case hookError:
				delete(defEvents, ei.Info.Eid)
			case hookSuccess:
				t.Errorf("该处不应该成功")
				delete(defEvents, ei.Info.Eid)
			case storeOne:
				storecnt++
			}
		},
	})
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
