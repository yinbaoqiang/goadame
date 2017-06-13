package engine

import (
	"sync"
	"testing"
)

func Test_listenManage_append(t *testing.T) {

	var lm = createListenManager()

	type fields struct {
		lmap map[string]*hookURL
		lck  sync.RWMutex
	}
	type args struct {
		etype  string
		action string
		url    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "测试1_1", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1"}},
		{name: "测试1_2", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_2"}},
		{name: "测试2", args: args{etype: "typ_2", action: "action_2", url: "http://test/test2"}},
		{name: "测试3", args: args{etype: "typ_1", action: "action_3", url: "http://test/test3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lm.Add(tt.args.url, tt.args.etype, tt.args.action)
			us := lm.GetAll(tt.args.etype, tt.args.action)
			r := false
			for _, u := range us {
				if string(u) == tt.args.url {
					r = true
					break
				}
			}
			if !r {
				t.Errorf("%s测试错误", tt.name)
			}
		})
	}

}
