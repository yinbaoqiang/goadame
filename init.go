package main

import (
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/yinbaoqiang/goadame/engine"
	"github.com/yinbaoqiang/goadame/store"
	"github.com/yinbaoqiang/goadame/store/etcd"
)

func initEngine(endpoints []string, dialTimeout int) error {
	// 初始化store
	var cfg clientv3.Config
	cfg.Endpoints = endpoints
	cfg.DialTimeout = time.Duration(dialTimeout) * time.Second
	s := etcd.CreateStore(cfg)
	store.SetDefaultListener(s)
	engine.SetEventStorer(store.CreateEventStore())
	engine.SetListenerStore(s)
	err := engine.Start()

	if err != nil {
		return err
	}
	return nil
}
