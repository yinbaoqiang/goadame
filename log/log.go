package log

import (
	"git.oschina.net/antlinker/antmqtt/antserver/ulog"
)

const (
	// LogTag 日志标签
	LogTag = "EventEngine"
)

// Elog 日志管理器
var Elog = ulog.Ulog

func init() {
	Elog.SetLogTag(LogTag)
	//Mlog.SetEnabled(false)
}

func ElogInit(configs string) {
	Elog.ReloadConfig(configs)
	Elog.SetLogTag(LogTag)
}
