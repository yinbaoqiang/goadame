package store

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collNameServerStatusHistory = "serverHis" // 事件发生集合
)

// SaveStart 存储启动成功事件
func SaveStart(server string) error {
	err := ExecSync(collNameServerStatusHistory, func(coll *mgo.Collection) error {
		q := bson.M{"serverName": server, "startTime": time.Now()}

		return coll.Insert(q)
	})
	if err != nil {
		log.Printf("SaveStart存储失败:%v\n", err)

	}
	return err
}
