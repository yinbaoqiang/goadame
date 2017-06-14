package engine

import (
	"fmt"
	"sync"
	"testing"
)

func Test_listenManage1(t *testing.T) {

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
		{name: "测试1_1_0", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1"}},
		{name: "测试1_1_1", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1"}},
		{name: "测试1_1_2", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1"}},
		{name: "测试1_1_3", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1"}},
		{name: "测试1_2", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_2"}},
		{name: "测试1_3", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_3"}},
		{name: "测试1_4", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_4"}},
		{name: "测试1_5", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_5"}},
		{name: "测试2", args: args{etype: "typ_2", action: "action_2", url: "http://test/test2"}},
		{name: "测试3", args: args{etype: "typ_1", action: "action_3", url: "http://test/test3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lm.Add(tt.args.url, tt.args.etype, tt.args.action)
			us := lm.GetAll(tt.args.etype, tt.args.action)
			r := false
			for _, u := range us {
				if u.url == tt.args.url {
					r = true
					break
				}
			}
			if !r {
				t.Errorf("%s测试错误", tt.name)
				return
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lm.Remove(tt.args.url, tt.args.etype, tt.args.action)
			us := lm.GetAll(tt.args.etype, tt.args.action)
			r := false
			for _, u := range us {
				if u.url == tt.args.url {
					r = true
					break
				}
			}
			if r {
				t.Errorf("%s测试remove错误", tt.name)
			}
		})
	}

}
func Test_listenManage2(t *testing.T) {

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
		{name: "测试1_0", args: args{etype: "typ_1", action: "", url: "http://test/test1_0"}},
		{name: "测试1_1", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1"}},
		{name: "测试1_2", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_2"}},
		{name: "测试1_3", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_3"}},
		{name: "测试1_4", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_4"}},
		{name: "测试2", args: args{etype: "typ_2", action: "action_2", url: "http://test/test2"}},
		{name: "测试3", args: args{etype: "typ_1", action: "action_3", url: "http://test/test3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lm.Add(tt.args.url, tt.args.etype, tt.args.action)
			us := lm.GetAll(tt.args.etype, tt.args.action)
			r := false
			for _, u := range us {
				if u.url == tt.args.url {
					r = true
					break
				}
			}
			if !r {
				t.Errorf("%s测试错误", tt.name)
			}
		})
	}
	us := lm.GetAll("typ_1", "abc")
	if us != nil {
		t.Error("测试错误")
		return
	}
	us = lm.GetAll("typ_1", "")
	if us == nil {
		t.Error("测试错误")
		return
	}
	tests = []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "测试1_1", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1"}},
		{name: "测试1_4", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_4"}},
		{name: "测试1_2", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_2"}},
		{name: "测试1_3", args: args{etype: "typ_1", action: "action_1", url: "http://test/test1_3"}},
		{name: "测试2", args: args{etype: "typ_2", action: "action_2", url: "http://test/test2"}},
		{name: "测试3", args: args{etype: "typ_1", action: "action_3", url: "http://test/test3"}},
		{name: "测试4", args: args{etype: "typ_1", action: "action_3", url: "http://test/test4"}},
	}
	l := len(tests)
	for i := l - 1; i >= 0; i-- {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			lm.Remove(tt.args.url, tt.args.etype, tt.args.action)
			us := lm.GetAll(tt.args.etype, tt.args.action)
			r := false
			for _, u := range us {
				if u.url == tt.args.url {
					r = true
					break
				}
			}
			if r {
				t.Errorf("%s测试remove错误", tt.name)
			}
		})
	}

}

func Test_createHook(t *testing.T) {
	h := createHook("test")
	n := 10
	c := make(chan int, n*2)
	for i := 0; i < n; i++ {
		j := i
		h.put(func() {
			fmt.Println("send:", j)
			c <- j
		})
	}
	m := 0
	for i := range c {
		fmt.Println("reveive:", i)
		if i != m {
			t.Error("结果错误")
		}
		m++
		if m == n {
			break
		}
	}

}
