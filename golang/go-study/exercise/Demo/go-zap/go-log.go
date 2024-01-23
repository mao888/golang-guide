package main

import (
	"context"
	glog "github.com/mao888/mao-glog"
	"github.com/mao888/mao-gutils/constants"
)

func init() {
	glog.Init(
		//打开控制台日志，默认关闭
		glog.WithConsoleStdout(),
		//默认 level 为 info
		glog.WithLevel(glog.DebugLevel),
		//设置关闭自动压缩文件，默认打开
		glog.WithOffCompress(),
		//日志文件位置，默认 ./log.log
		glog.WithFileLocation("test.log"),
		// 设置日志保存天数，默认30
		glog.WithLogMaxAge(30),
		//设置最大文件大小（MB），默认256
		glog.WithLogMaxSize(250),
		//设置全局自定义字段
		glog.WithCustomizedGlobalField(map[string]interface{}{constants.LoggerServerCode: constants.ServiceCode}),
		//设置覆盖默认字段
		glog.WithCoverDefaultKey(glog.CoverDefaultKey{
			LevelKey:      "",
			TimeKey:       "timestamp",
			CallerKey:     "label",
			MessageKey:    "message",
			StacktraceKey: "",
		}),
	)
}

func main() {
	// 日志打印
	// 方式一：返回全局 logger 变量方式
	glog.C(context.Background()).Debugf("debugf: %s", "dddd")
	glog.C(context.Background()).Infof("test: %s", "info")

	// 方式二：直接调用包方法
	glog.Debug(context.Background(), "test debug")
	glog.Infof(context.Background(), "test: %s", "info")
	glog.Debugf(context.Background(), "debugf: %s", "dddd")
	// 也支持打印时新加字段，但仅影响本次调用，不会影响全局字段，仅支持打印 info 日志
	glog.InfoWithField(context.Background(), map[string]interface{}{
		"temp_field": "glog is good ",
	}, "msg1", "msg2")
}
