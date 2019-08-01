/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/8/1 20:44
* @Description: The file is for
***********************************************************************/

package monitor

import (
	"net"
	"net/http"
	"time"
)

// 任务的HTTP接口
type ApiServer struct {
	httpServer *http.Server
}

var (
	// 单例对象，提供外部调用，这里InitApiServer需要在main/main.go去调用
	E_apiServer *ApiServer	// 初始为空
)

// 初始化服务
func InitApiServer() (err error) {

	var (
		mux *http.ServeMux
		listener net.Listener
		httpServer *http.Server
		staticDir http.Dir
		staticHandler http.Handler
	)
	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/enode/list", handleListAllNodes)
	//handlefunc是用来处理动态接口的

	// 静态文件目录
	staticDir = http.Dir(Config().serverWebRoot)
	staticHandler = http.FileServer(staticDir)
	staticHandler = http.StripPrefix("/", staticHandler)	//去掉url /index.html 的前缀"/"
	mux.Handle("/", staticHandler)	//分析url，看哪个最匹配（重合的长度最长）  / + index.html

	// 启动TCP监听
	if listener, err = net.Listen("tcp", Config().serverIpAddress); err != nil {
		return
	}

	// 创建一个HTTP服务
	httpServer = &http.Server{
		ReadTimeout:time.Duration(Config().serverReadTimeout) * time.Millisecond,
		WriteTimeout:time.Duration(Config().serverWriteTimeout) * time.Millisecond,
		Handler:mux,
	}

	// 赋值单例   创建单例对象，以便外部调用
	E_apiServer = &ApiServer{
		httpServer:httpServer,
	}

	// 启动服务端，在协程中监听
	go httpServer.Serve(listener)

	return
}

// 列举所有网关节点
func handleListAllNodes(w http.ResponseWriter, r *http.Request) {

}
