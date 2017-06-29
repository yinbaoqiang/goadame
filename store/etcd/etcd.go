package etcd

import (
	"context"
	"errors"
	"fmt"
	"log"

	"time"

	"encoding/json"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/goadesign/goa/uuid"
	"github.com/yinbaoqiang/goadame/app"
	"github.com/yinbaoqiang/goadame/store"
)

const (
	etcdDir       = "antEventEngine"
	etcdDirListen = "antEventEngine/listen/"
)

// Store 接口组合
type Store interface {
	store.ChgListener
	store.Listener
}

// CreateStore 创建 Store 接口实现
func CreateStore(cfg clientv3.Config) Store {
	var es etcdStore
	es.cfg = cfg
	es.requestTimeout = cfg.DialTimeout
	if es.requestTimeout == 0 {
		es.requestTimeout = 5 * time.Second
	}
	return es
}

type client struct {
	cfg            clientv3.Config
	requestTimeout time.Duration
}

func (c client) exec(handler func(cli *clientv3.Client) error) error {
	cli, err := clientv3.New(c.cfg)
	if err != nil {
		log.Printf("连接etcd失败:%v", err)
		return err
	}
	defer cli.Close()
	return handler(cli)
}
func (c client) execTimeout(handler func(ctx context.Context, cli *clientv3.Client) error) error {
	return c.exec(func(cli *clientv3.Client) (err error) {
		ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
		defer cancel()
		err = handler(ctx, cli)
		switch err {
		case context.Canceled:
			return fmt.Errorf("ctx被另一个例程取消: %v", err)
		case context.DeadlineExceeded:
			return fmt.Errorf("ctx已经超时关闭: %v", err)
		case rpctypes.ErrEmptyKey:
			return fmt.Errorf("客户端错误: %v", err)
		default:
			return fmt.Errorf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	})
}
func (c client) put(key string, val string, opo ...clientv3.OpOption) error {
	return c.execTimeout(func(ctx context.Context, cli *clientv3.Client) (err error) {
		if len(opo) == 0 {
			_, err = cli.Put(ctx, key, val)
		} else {

			_, err = cli.Put(ctx, key, val, opo...)
		}
		return err
	})

}
func (c client) delete(key string, opo ...clientv3.OpOption) error {
	return c.execTimeout(func(ctx context.Context, cli *clientv3.Client) (err error) {
		if len(opo) == 0 {
			_, err = cli.Delete(ctx, key)
		} else {

			_, err = cli.Delete(ctx, key, opo...)
		}
		return err
	})

}
func (c client) get(key string, opo ...clientv3.OpOption) (resp *clientv3.GetResponse, err error) {
	c.execTimeout(func(ctx context.Context, cli *clientv3.Client) (err error) {

		if len(opo) == 0 {
			resp, err = cli.Get(ctx, key)
		} else {
			resp, err = cli.Get(ctx, key, opo...)
		}
		return err
	})
	return
}
func (c client) watchPrefix(key string, opevt func(evt *clientv3.Event)) error {
	return c.exec(func(cli *clientv3.Client) (err error) {
		rch := cli.Watch(context.Background(), "foo", clientv3.WithPrefix())
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				opevt(ev)
			}
		}
		return nil
	})

}

type etcdStore struct {
	client
}

func (c etcdStore) Watch(handler func(ctyp store.ChgType, lis app.AntListen)) {
	go func() {
		err := c.watchPrefix(etcdDirListen, func(evt *clientv3.Event) {
			var lis app.AntListen
			var ctyp store.ChgType
			if evt.IsCreate() {
				ctyp = store.ChgTypeAdd
				err := toAntListent(evt.Kv.Value, &lis)
				if err != nil {
					log.Printf("收到错误的删除:%v\n", err)
					return
				}
				handler(ctyp, lis)
			} else if evt.IsModify() {
				ctyp = store.ChgTypeUpdate
				err := toAntListent(evt.Kv.Value, &lis)
				if err != nil {
					log.Printf("收到错误的删除:%v\n", err)
					return
				}
				handler(ctyp, lis)
			} else if evt.Type == clientv3.EventTypeDelete {
				ctyp = store.ChgTypeRemove
				err := toAntListent(evt.PrevKv.Value, &lis)
				if err != nil {
					log.Printf("收到错误的删除:%v\n", err)
					return
				}
				handler(ctyp, lis)
			}
		})
		if err != nil {
			log.Printf("监测监听配置发送错误关闭:%v\n", err)
			return
		}
		log.Println("监测监听配置关闭")
	}()
	return

}

func toAntListent(data []byte, lis *app.AntListen) error {
	if data == nil {
		return errors.New("数据不能为空")
	}
	return json.Unmarshal(data, &lis)
}

// Add 新增监听
func (c etcdStore) Add(lis app.AntListen) error {
	if lis.Rid == "" {
		lis.Rid = uuid.NewV4().String()
	}
	data, _ := json.Marshal(lis)
	return c.put(etcdDirListen+lis.Rid, string(data))
}

// Update 修改监听
func (c etcdStore) Update(lis app.AntListen) error {
	return c.Add(lis)
}

// List 查询列表
func (c etcdStore) List(action, etype string, page, cnt int, total *int) (res []*app.AntListen, err error) {
	resp, err := c.get(etcdDirListen, clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, err
	}
	res = make([]*app.AntListen, 0, len(resp.Kvs))
	for _, ev := range resp.Kvs {
		var lis app.AntListen
		toAntListent(ev.Value, &lis)
		res = append(res, &lis)
	}
	return
}
func (c etcdStore) Rmove(rid string) error {
	return c.delete(etcdDirListen + rid)
}

// cli, err := clientv3.New(clientv3.Config{
//     Endpoints:   endpoints,
//     DialTimeout: dialTimeout,
// })
// if err != nil {
//     log.Fatal(err)
// }
// defer cli.Close()
