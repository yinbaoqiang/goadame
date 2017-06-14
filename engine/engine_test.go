package engine

import (
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
	body, _ := req.GetBody()
	data, _ := ioutil.ReadAll(body)
	fmt.Println("data:", string(data))
	resp.WriteHeader(200)
}
func (t *testServer) Start() {
	go http.ListenAndServe(":12345", t)
}

func TestMain(m *testing.M) {
	(&testServer{}).Start()
	os.Exit(m.Run())
}

func TestOK(t *testing.T) {

	engine := CreateEventEnginer()
	engine.ListenManager().Add("http://localhost:12345", "typ_test", "action_test")
	engine.Start()
	engine.Put(Event{
		Info: EventInfo{
			Eid:    "10920192",
			Etype:  "type_test",
			Action: "action_test",
		},
	})
	time.Sleep(3 * time.Second)
	engine.Stop()
}
