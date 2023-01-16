package main

import (
	"context"
	glog "github.com/mao888/go-log"
)

//func init()  {
//
//}
//
//Init(
////打开控制台日志，默认关闭
//WithConsoleStdout(),
////默认 level 为 info
//WithLevel(levelType),
////设置关闭自动压缩文件，默认打开
//WithOffCompress(),
////日志文件位置，默认 ./log.log
//WithFileLocation("test.log"),
//// 设置日志保存天数，默认30
//WithLogMaxAge(30),
////设置最大文件大小（MB），默认256
//WithLogMaxSize(250),
////设置全局自定义字段
//WithCustomizedGlobalField(globalFields),
////设置覆盖默认字段
//WithCoverDefaultKey(CoverDefaultKey{
//TimeKey:    "timestamp",
//CallerKey:  "label",
//MessageKey: "message"}),
//)
//
//// 日志打印
//glog.C(ctx).Debug("test debug")
//glog.C(ctx).Infof("test: %s","info")
//
//// 也支持打印时新加字段，但仅影响本次调用，不会影响全局字段，仅支持打印 info 日志
//glog.C(ctx).InfoWithField(map[string]interface{}{
//"temp_field":"glog is good "
//}, "msg1","msg2")

func main() {
	////默认 level 为 info
	glog.Init(glog.WithLevel("debug"),
		glog.WithConsoleStdout())

	// 日志打印
	glog.Debug(context.Background(), "test debug")
	glog.Infof(context.Background(), "test: %s", "info")
	//glog.Debugf(context.Background(), "debugf: %s", "dddd")
	glog.C(context.Background()).Debugf("debugf: %s", "dddd")
	// 也支持打印时新加字段，但仅影响本次调用，不会影响全局字段，仅支持打印 info 日志
	glog.InfoWithField(context.Background(), map[string]interface{}{
		"temp_field": "glog is good ",
	}, "msg1", "msg2")
}
