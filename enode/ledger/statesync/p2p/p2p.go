/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 14:48
* @Description: The file is for
***********************************************************************/

package p2p

import (
	"EChain/enode/ledger"
	"EChain/enode/ledger/blockchain"
	"EChain/enode/utils"
	"fmt"
	DEATH "github.com/vrecan/death"
	"io/ioutil"
	"log"
	"net"
	"os"
	"runtime"
	SYS "syscall"
)

//此处节点的意义是区块链网络中的客户端节点IP地址，形如“IP:PORT”
//更贴切的称呼是peer

const (
	protocol      = "tcp"
	version       = 1
	commandLength = 12
)

var (
	// 这是已知节点集合
	KnownNodes     = []string{"127.0.0.1:3000",}
	blockInTransit [][]byte
	memoryPool     = make(map[string]blockchain.Tx)
)

//流程
//1.create Blockchain
//2.Wallet connects and Download Blockchain
//3.Miner connects and download Blockchain
//4.Wallets creates tx
//5.miner gets tx to memory pool
//6.enough tx -> mine block
//7.block sent to central node	//
//8.wallet syncs and verifies

// 处理连接，对请求做出处理
func HandleConnection(conn net.Conn, chain *blockchain.Chain) {
	//读取request
	req, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//从request获取command
	command := utils.BytesToCmd(req[:commandLength])
	fmt.Printf("Received %s command\n", command)

	//对命令作出对应处理
	switch command {
	case "addr":
		HandleAddr(req)
	case "block":
		HandleBlock(req, chain)
	case "inv":
		HandleInv(req, chain)
	case "getblocks":
		HandleGetBlocks(req, chain)
	case "getdata":
		HandleGetData(req, chain)
	case "tx":
		HandleTx(req, chain)
	case "version":
		HandleVersion(req, chain)

	default:
		fmt.Println(ledger.ErrUnknownCommand)
	}
}

/*开启服务器，监听请求并跳转至处理连接方法*/
func StartServer(cc *ledger.Conf) {

	//本地节点开启监听
	ln, err := net.Listen(protocol, cc.Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	//打开数据库获取区块链对象，关闭数据库
	//TODO:ContinueB...函数暂时不需要address参数
	chain, err := blockchain.ContinueChain(cc)
	if err != nil {
		log.Fatal(err)
	}
	defer chain.Db.Close()
	go CloseDB(chain)  //

	//如果本地节点不是已知节点第一个节点，那么发送向已知节点集第一个节点发送版本
	//TODO
	if cc.Addr != KnownNodes[0] {
		SendVersion(KnownNodes[0], chain)
	}

	//循环：接受请求，处理连接
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleConnection(conn, chain)
	}
}

/*检查某节点是否在已知节点集合中*/
func NodeIsKnown(addr string) bool {
	for _, node := range KnownNodes {
		if node == addr {
			return true
		}
	}

	return false
}

/*关闭数据库*/
func CloseDB(chain *blockchain.Chain) {

	//SIGINT标识中断信号；SIGTERM标识进程终结
	//当产生中断时，返回death对象
	d := DEATH.NewDeath(SYS.SIGINT, SYS.SIGTERM, os.Interrupt)

	//保证进程结束时关闭数据库
	d.WaitForDeathWithFunc(func() {
		defer os.Exit(1)
		defer runtime.Goexit()
		chain.Db.Close()
	})
}
