/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/6/26 10:51
* @Description: The file is for
***********************************************************************/

package main

//
//import (
//	"EChain/enode/cmd"
//	"EChain/enode/conf"
//	"EChain/enode/ledger"
//	"EChain/enode/logger"
//	log "github.com/sirupsen/logrus"
//	flag "github.com/spf13/pflag"
//	"runtime"
//)
//
////目标流程
////初始化环境
////初始化配置
////初始化日志记录
////程序启动协程进行数据采集：采集器
////启动协程对数据做分析判断
////启动协程
//
//var (
//	configFile string  //配置文件路径
//)
//
//func main() {
//
//	var err error
//
//	initEnv()
//	initArgs()
//
//	if err = conf.InitConfigByViper(configFile); err != nil {
//		goto ERR
//	}
//
//
//
//	if err = logger.InitLogger(); err != nil {
//		goto ERR
//	}
//
//	if err = ledger.InitLedger(); err != nil {
//		goto ERR
//	}
//
//	// 执行命令行程序
//	cmd.Execute()
//
//
//	return
//ERR:
//	log.Error(err)
//}
//
//// 多线程
//func initEnv() {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//}
//
//// 初始化命令行参数 $ enode -config ./enode.json
//func initArgs() {
//	//// 注意这里的执行程序。这里是给了个默认config值，但是没有解析参数。解析参数丢给了cmd去做。
//	//// 这意味着，如果程序运行时只输入enode则默认加载默认配置文件。如果需要额外指定，需要enode --config configfile
//	//// 强调的是，这么写，程序一开始运行绝对是按默认配置文件来！！！！
//	flag.StringVar(&configFile, "config", "../conf/enode.yaml", "指定配置文件路径")
//
//	// 为了解决这个问题，将参数定义及解析全部合并到一处。代码至于cmd下flags.go下，此处调用。而cmd则不使用flags.go
//}
// TODO:设置默认配置文件路径

import "EChain/enode/cmd"

func main() {
	cmd.Execute()
}