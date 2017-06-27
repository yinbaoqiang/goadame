package main

import (
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/yinbaoqiang/goadame/store"
	"github.com/yinbaoqiang/goadame/store/etcd"
)

func initApp(endpoints []string, dialTimeout int) {
	// 初始化store
	var cfg clientv3.Config
	cfg.Endpoints = endpoints
	cfg.DialTimeout = time.Duration(dialTimeout) * time.Second
	s := etcd.CreateStore(cfg)
	store.SetDefaultChgListener(s)
	store.SetDefaultListener(s)
}
