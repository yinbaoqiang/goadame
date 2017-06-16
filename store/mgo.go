package store

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"sync"

	util "github.com/antlinker/taskpool"

	"gopkg.in/mgo.v2"
)

// MgoTask 执行任务任务
type MgoTask struct {
	util.BaseTask
	dbname   string
	collname string
	task     func(coll *mgo.Collection, opt ...interface{}) error
	opt      []interface{}
}

func (t *MgoTask) call(coll *mgo.Collection, opt ...interface{}) error {
	return t.task(coll, opt...)
}

// CreateTask 创建一个mongodb任务,指定集,合任务
func CreateTask(taskname, collname string, task func(coll *mgo.Collection, opt ...interface{}) error, opt ...interface{}) *MgoTask {
	return &MgoTask{
		collname: collname,
		task:     task,
		opt:      opt,
	}
}

// CreateDBTask 创建一个mongodb任务,指定数据库,集合,任务
func CreateDBTask(taskname, dbname, collname string, task func(coll *mgo.Collection, opt ...interface{}) error, opt ...interface{}) *MgoTask {

	return &MgoTask{
		dbname:   dbname,
		collname: collname,
		task:     task,
	}
}

const (
	taskName = "mongodb任务"
)

var defaultMgo MyMgoer

func checkInit() {
	if defaultMgo == nil {
		panic("请初始化后在使用该方法")
	}
}

// ExecSync 使用默认配置同步执行任务
func ExecSync(collname string, task func(coll *mgo.Collection) error) error {
	checkInit()
	return defaultMgo.ExecSync(collname, task)
}

// ExecDBAsync 使用默认配置默认异步执行任务，可以指定数据库
func ExecDBAsync(db, collname string, task func(coll *mgo.Collection, opt ...interface{}) error, opt ...interface{}) {
	checkInit()
	defaultMgo.ExecDBAsync(db, collname, task, opt...)
}

// MyMgoer mgo数据库操作
type MyMgoer interface {
	ExecSync(collname string, task func(coll *mgo.Collection) error) error

	ExecDBAsync(db, collname string, task func(coll *mgo.Collection, opt ...interface{}) error, opt ...interface{})
}
type myMgo struct {
	url           string
	session       *mgo.Session
	defaultDBName string
}

var lock = &sync.Mutex{}

// MongodbConfig mongodb 配置
type MongodbConfig struct {
	URL    string `json:"url" yaml:"url"`
	DBName string `json:"dbName" yaml:"dbName"`
}

// CreateMongodbConfigForEnv 通过环境变量创建mongodb连接信息
func CreateMongodbConfigForEnv(pre string) *MongodbConfig {
	tmp := os.Getenv(pre + "_MGO_URL")
	if tmp == "" {
		return nil
	}
	var cfg MongodbConfig
	cfg.URL = tmp
	tmp = os.Getenv(pre + "_MGO_DBNAME")
	if tmp != "" {
		cfg.DBName = tmp
	}

	return &cfg
}

var one sync.Once

// InitDefautMyMgoForCfg 初始化默认数据连接
func InitDefautMyMgoForCfg(cfg MongodbConfig) {
	one.Do(func() {
		defaultMgo = CreateMyMgoForCfg(cfg)
	})
}

// InitDefautMyMgo 初始化默认数据连接
func InitDefautMyMgo(url, dbname string) {
	defaultMgo = CreateMyMgo(url, dbname)
}

// CreateMyMgoForCfg 通过配置参数创建db
func CreateMyMgoForCfg(cfg MongodbConfig) MyMgoer {

	lock.Lock()
	defer lock.Unlock()
	mymgo := &myMgo{
		url:           cfg.URL,
		defaultDBName: cfg.DBName,
	}
	tmp, err := mgo.Dial(mymgo.url)
	if err != nil {
		panic(errors.New("创建mongodb连接失败1:" + fmt.Sprintf("%v\n%v", err, cfg)))
	}
	mymgo.session = tmp
	// Optional. Switch the session to a monotonic behavior.
	mymgo.session.SetMode(mgo.Eventual, true)

	return mymgo
}

// CreateMyMgo 创建数据库连接,其他使用默认设置
func CreateMyMgo(url, dbname string) MyMgoer {
	lock.Lock()
	defer lock.Unlock()
	mymgo := &myMgo{
		url:           url,
		defaultDBName: dbname,
	}

	tmp, err := mgo.Dial(mymgo.url)
	if err != nil {
		debug.PrintStack()
		panic(errors.New("创建mongodb连接失败2:" + fmt.Sprintf("%v", err)))
	}
	mymgo.session = tmp
	// Optional. Switch the session to a monotonic behavior.
	mymgo.session.SetMode(mgo.Strong, true)

	return mymgo
}

func (m *myMgo) execTask(task *MgoTask) error {
	session, err := m.getSession()
	if err != nil {
		if err != nil {
			mlog.Errorf("mongodb获取Session失败:%v", err)
		}
		return err
	}
	defer session.Close()
	var dbname = task.dbname
	if task.dbname == "" {
		dbname = m.defaultDBName
	}
	err = task.call(session.DB(dbname).C(task.collname), task.opt...)
	if err != nil {
		mlog.Errorf("mongodb存储失败:%v", err)
	}
	return err
}

//同步执行
func (m *myMgo) ExecSync(collname string, task func(coll *mgo.Collection) error) error {
	session, err := m.getSession()
	if err != nil {
		return fmt.Errorf("获取mongodb数据库会话失败:%v", err)
	}
	defer session.Close()
	return task(session.DB(m.defaultDBName).C(collname))
}

//异步执行
func (m *myMgo) ExecAsyncTask(task *MgoTask) {
	m.asyncTaskOperater.ExecAsyncTask(task)

}

//异步执行
func (m *myMgo) ExecDBAsync(db, collname string, task func(coll *mgo.Collection, opt ...interface{}) error, opt ...interface{}) {
	m.ExecAsyncTask(CreateDBTask(taskName, db, collname, task, opt...))
}

//异步执行
func (m *myMgo) ExecAsync(collname string, task func(coll *mgo.Collection, opt ...interface{}) error, opt ...interface{}) {
	m.ExecAsyncTask(CreateTask(taskName, collname, task, opt...))
}
func (m *myMgo) getSession() (*mgo.Session, error) {

	return m.session.Clone(), nil
}

func (m *myMgo) ExecTask(task util.Task) error {
	mtask, ok := task.(*MgoTask)
	if ok {

		return m.execTask(mtask)
	}
	return nil
}
