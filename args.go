package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/goadesign/goa"
	"github.com/yinbaoqiang/goadame/engine"
	"github.com/yinbaoqiang/goadame/store"
	"gopkg.in/urfave/cli.v2"
)

func createStart() []cli.Flag {
	//jobid string, gorunCnt int, loadSpaceTime time.Duration, maxerrcnt
	fs := []cli.Flag{
		&cli.StringFlag{
			Name:    "servername",
			Aliases: []string{"sn"},
			Value:   "tmp",
			Usage:   "mongodb数据库连接",
		},
		&cli.StringFlag{
			Name:    "mgourl",
			Aliases: []string{"ml"},
			Value:   "localhost:27017",
			Usage:   "mongodb数据库连接",
		},
		&cli.StringFlag{
			Name:    "mgodbname",
			Aliases: []string{"mn"},
			Value:   "antevent",
			Usage:   "mongodb数据库名",
		},
		&cli.StringSliceFlag{
			Name:    "etcdendpoint",
			Aliases: []string{"eep"},
			Value:   cli.NewStringSlice("localhost:2379"),
			Usage:   "etcd数据库连接",
		},
		&cli.IntFlag{
			Name:    "etcdtimeout",
			Aliases: []string{"eto"},
			Value:   10,
			Usage:   "etcd数据库连接超时时间",
		},
		&cli.IntFlag{
			Name:    "hooktimeout",
			Aliases: []string{"hto"},
			Value:   10,
			Usage:   "回调请求超时时间",
		},
		&cli.IntFlag{
			Name:    "errtrycnt",
			Aliases: []string{"etc"},
			Value:   3,
			Usage:   "回调请求错误重试次数",
		},
	}
	return fs
}

func initArgs(args []string, service *goa.Service) {
	app := &cli.App{
		Authors: []*cli.Author{
			{Name: "@antlinker.com"},
		},
		Name:    "事件引擎",
		Version: "1.0",
		Usage:   "事件调度引擎,可以将事件发布给订阅者",

		Flags: createStart(),
		Action: func(c *cli.Context) error {

			initMgo(c.String("mgourl"), c.String("mgodbname"))

			err := store.SaveStart(c.String("servername"))
			if err != nil {
				return err
			}
			// 初始化应用
			err = initEngine(c.StringSlice("etcdendpoint"), c.Int("etcdtimeout"), c.Int("errtrycnt"), c.Int("hooktimeout"))
			// 启动事件引擎
			if err != nil {
				return err
			}
			defer engine.Stop()
			go regSigin()
			// Start service
			if err := service.ListenAndServe(":8080"); err != nil {
				service.LogError("startup", "err", err)
				return err
			}

			return nil

		},
	}

	err := app.Run(args)
	if err != nil {
		log.Println("启动失败:", err)
	}
}
func regSigin() {

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)
	<-sigc
	go func() {
		signal.Notify(sigc, os.Interrupt, os.Kill)
		<-sigc
		log.Println("在此收到了退出程序信号，强制退出")
		os.Exit(9)
	}()
	log.Println("收到了退出程序信号")
	engine.Stop()
	os.Exit(0)

}
