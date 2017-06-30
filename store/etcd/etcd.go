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
	"github.com/yinbaoqiang/goadame/engine"
	"github.com/yinbaoqiang/goadame/store"
)

const (
	etcdDir                 = "/antEventEngine"
	etcdDirListen           = etcdDir + "/listen/"
	etcdDirIndexEtype       = etcdDir + "/index/etype/"
	etcdDirIndexEtypeAction = etcdDir + "/index/ea/"
	etcdDirTxnError         = etcdDir + "/txnerror/"
)

// Store 接口组合
type Store interface {
	engine.ListenerStore
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
	defer func() {
		_ = cli.Close()
	}()
	err = handler(cli)
	if err != nil {
		switch err {
		case context.Canceled:
			log.Printf("ctx被另一个例程取消: %v\n", err)
			return fmt.Errorf("ctx被另一个例程取消: %v", err)
		case context.DeadlineExceeded:
			log.Printf("ctx已经超时关闭: %v", err)
			return fmt.Errorf("ctx已经超时关闭: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Printf("客户端错误: %v", err)
			return fmt.Errorf("客户端错误: %v", err)
		default:
			log.Printf("bad cluster endpoints, which are not etcd servers: %v", err)
			return fmt.Errorf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	return nil
}

func (c client) execTimeout(handler func(ctx context.Context, cli *clientv3.Client) error) error {
	return c.exec(func(cli *clientv3.Client) (err error) {
		ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
		defer cancel()
		return handler(ctx, cli)

	})
}

// func (c client) put(key string, val string, opo ...clientv3.OpOption) error {
// 	return c.execTimeout(func(ctx context.Context, cli *clientv3.Client) (err error) {
// 		log.Println("----put", key, "=", val)
// 		if len(opo) == 0 {

// 			_, err = cli.Put(ctx, key, val)
// 		} else {

// 			_, err = cli.Put(ctx, key, val, opo...)
// 		}
// 		return err
// 	})

// }

func (c client) txn(ifcmps []clientv3.Cmp, thenop []clientv3.Op, elseop []clientv3.Op) error {
	return c.execTimeout(func(ctx context.Context, cli *clientv3.Client) (err error) {
		txn := cli.Txn(ctx)

		txn = txn.If(ifcmps...).Then(thenop...)
		if len(elseop) > 0 {
			txn.Else(thenop...)
		}
		_, err = txn.Commit()
		return
	})

}

// func (c client) delete(key string, opo ...clientv3.OpOption) error {
// 	return c.execTimeout(func(ctx context.Context, cli *clientv3.Client) (err error) {
// 		if len(opo) == 0 {
// 			_, err = cli.Delete(ctx, key)
// 		} else {

// 			_, err = cli.Delete(ctx, key, opo...)
// 		}
// 		return err
// 	})

// }
func (c client) get(key string, opo ...clientv3.OpOption) (resp *clientv3.GetResponse, err error) {
	_ = c.execTimeout(func(ctx context.Context, cli *clientv3.Client) (err error) {

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
		rch := cli.Watch(context.Background(), key, clientv3.WithPrefix())
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

func (c etcdStore) Watch(handler func(ctyp engine.ChgType, lis app.AntListen)) {
	go func() {
		err := c.watchPrefix(etcdDirListen, func(evt *clientv3.Event) {
			var lis app.AntListen
			var ctyp engine.ChgType
			fmt.Printf("------收到变化: %s %q : %q\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
			if evt.IsCreate() {
				ctyp = engine.ChgTypeAdd
				err := toAntListent(evt.Kv.Value, &lis)
				if err != nil {
					log.Printf("收到错误的删除:%v\n", err)
					return
				}
				handler(ctyp, lis)
			} else if evt.IsModify() {
				ctyp = engine.ChgTypeUpdate
				err := toAntListent(evt.Kv.Value, &lis)
				if err != nil {
					log.Printf("收到错误的删除:%v\n", err)
					return
				}
				handler(ctyp, lis)
			} else if evt.Type == clientv3.EventTypeDelete {
				ctyp = engine.ChgTypeRemove
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
}

func toAntListent(data []byte, lis *app.AntListen) error {
	if data == nil {
		return errors.New("数据不能为空")
	}
	return json.Unmarshal(data, &lis)
}
func getIndexEtypeKey(lis app.AntListen) string {
	return etcdDirIndexEtype + lis.Etype + "/" + lis.Rid
}

func getIndexEtypeActionKey(lis app.AntListen) string {
	action := ""
	if lis.Action != nil {
		action = *lis.Action
	}
	return etcdDirIndexEtypeAction + lis.Etype + "/" + action + "/" + lis.Rid
}

//  新增监听
func (c etcdStore) Add(lis app.AntListen) error {

	log.Println("-----------", lis)
	if lis.Rid == "" {
		lis.Rid = uuid.NewV4().String()
	}
	data, _ := json.Marshal(lis)
	key := etcdDirListen + lis.Rid
	value := string(data)
	err := c.txn([]clientv3.Cmp{clientv3.Compare(clientv3.Value(key), "!=", value)},
		[]clientv3.Op{
			clientv3.OpPut(key, value),
			clientv3.OpPut(getIndexEtypeKey(lis), value),
			clientv3.OpPut(getIndexEtypeActionKey(lis), value),
		}, []clientv3.Op{
			clientv3.OpPut(etcdDirTxnError+"add/"+lis.Rid, value),
		})
	log.Println(err)
	return err
}

// Update 修改监听
func (c etcdStore) Update(lis app.AntListen) error {
	return c.Add(lis)
}

// List 查询列表
func (c etcdStore) List(etype, action string, previd string, cnt int) (total int, res []*app.AntListen, err error) {
	log.Printf("List:etype=%s,action=%s,previd=%s,cnt=%d\n", etype, action, previd, cnt)
	switch {
	case etype == "":
		return c.list(etcdDirListen, previd, cnt)
	case action == "":
		return c.list(etcdDirIndexEtype+etype+"/", previd, cnt)
	default:
		return c.list(etcdDirIndexEtypeAction+etype+"/"+action+"/", previd, cnt)
	}
}

var (
	endWithRange = string([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
)

// listEtype 查询列表
func (c etcdStore) list(keypre, previd string, cnt int) (total int, res []*app.AntListen, err error) {
	log.Printf("List:keypre=%s,previd=%s,cnt=%d\n", keypre, previd, cnt)
	resp, err := c.get(keypre, clientv3.WithPrefix(), clientv3.WithCountOnly())
	if err != nil {
		return 0, nil, err
	}

	total = int(resp.Count)
	if total == 0 {
		return
	}
	key := keypre
	if previd != "" {
		key = keypre + previd
	}

	resp, err = c.get(key, clientv3.WithRange(keypre+endWithRange), clientv3.WithLimit(int64(cnt+1)), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return 0, nil, err
	}
	if resp == nil || len(resp.Kvs) == 0 {
		log.Println("----查询信息为空")
		return
	}
	res = make([]*app.AntListen, 0, len(resp.Kvs))
	for _, ev := range resp.Kvs {
		var lis app.AntListen
		err := toAntListent(ev.Value, &lis)
		if err != nil {
			continue
		}
		res = append(res, &lis)
	}
	if previd == "" {
		if len(res) > cnt {
			res = res[0:cnt]
		}
		return
	}
	if len(res) > 0 {
		res = res[1:]
	}
	return
}

// List 查询列表
func (c etcdStore) All() (res []*app.AntListen, err error) {
	//, clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend)
	log.Println("----List查询信息")
	resp, err := c.get(etcdDirListen, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	if resp == nil || len(resp.Kvs) == 0 {
		log.Println("----查询信息为空")
		return
	}
	log.Println("----查询信息不为空")
	res = make([]*app.AntListen, 0, len(resp.Kvs))
	for _, ev := range resp.Kvs {
		var lis app.AntListen
		err := toAntListent(ev.Value, &lis)
		if err != nil {
			continue
		}
		res = append(res, &lis)
	}
	return
}
func (c etcdStore) Rmove(rid string) error {
	r, err := c.get(etcdDirListen + rid)
	if err != nil {
		return err
	}
	if len(r.Kvs) == 0 {
		return nil
	}
	var lis app.AntListen
	err = toAntListent(r.Kvs[0].Value, &lis)
	if err != nil {
		return err
	}
	key := etcdDirListen + rid
	return c.txn(
		[]clientv3.Cmp{clientv3.Compare(clientv3.Value(key), "!=", "")},
		[]clientv3.Op{
			clientv3.OpDelete(getIndexEtypeKey(lis)),
			clientv3.OpDelete(getIndexEtypeActionKey(lis)),
		}, []clientv3.Op{
			clientv3.OpPut(etcdDirTxnError+"delete/"+lis.Rid, "删除失败"),
		})

}

// cli, err := clientv3.New(clientv3.Config{
//     Endpoints:   endpoints,
//     DialTimeout: dialTimeout,
// })
// if err != nil {
//     log.Fatal(err)
// }
// defer cli.Close()
