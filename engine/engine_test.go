package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

type testServer struct {
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
	old, ok := defEvents[evt.Info.Eid]
	if !ok {
		gt.Errorf("收到错误的事件:%s\n", string(data))
		return
	}

	if old.Info != evt.Info {
		gt.Errorf("收到错误的事件:%s\n", string(data))
		return
	}
	delete(defEvents, old.Info.Eid)
	fmt.Println("data:", string(data))
	resp.WriteHeader(200)
}
func (t *testServer) Start() {
	go http.ListenAndServe(":12345", t)
}

func TestMain(m *testing.M) {
	(&testServer{}).Start()
	time.Sleep(100 * time.Millisecond)
	os.Exit(m.Run())
}

var (
	gt        *testing.T
	defEvents = map[string]Event{

		"10920192": Event{
			Info: EventInfo{
				Eid:    "10920192",
				Etype:  "typ_test",
				Action: "action_test",
			},
		},
		"10920193": Event{
			Info: EventInfo{
				Eid:    "10920193",
				Etype:  "typ_test",
				Action: "action_test",
			},
		},
		"10920194": Event{
			Info: EventInfo{
				Eid:    "10920194",
				Etype:  "typ_test",
				Action: "action_test2",
			},
		},
	}
)

func TestOK(t *testing.T) {
	gt = t
	engine := CreateEventEnginer()
	engine.ListenManager().Add("http://localhost:12345", "typ_test", "action_test")
	engine.Start()
	for _, e := range defEvents {
		engine.Put(e)
	}
	time.Sleep(1 * time.Second)
	engine.Stop()
	if len(defEvents) != 0 {
		t.Errorf("事件处理失败，还有剩余事件没处理:%d", len(defEvents))
	}
}

func TestOK2(t *testing.T) {
	gt = t
	engine := CreateEventEnginer()
	engine.ListenManager().Add("http://localhost:12345", "typ_test", "")
	engine.Start()
	for _, e := range defEvents {
		engine.Put(e)
	}
	time.Sleep(1 * time.Second)
	engine.Stop()
	if len(defEvents) != 0 {
		t.Errorf("事件处理失败，还有剩余事件没处理:%d", len(defEvents))
	}
}
