package p2p

import (
	"EChain/enode/ledger"
	"EChain/enode/utils"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

//维护本地存储的网络节点集合
type Addr struct {
	AddrList []string
}

// 处理请求节点信息地址
// 刚上线的节点A向某一节点B请求周遭所有已知节点信息，
// 随后B返回他指知道的节点，A调用这个方法更新本地维护的已知节点集合
// 随后向已知节点请求区块信息
func HandleAddr(request []byte) {
	//获取request内容
	var buff bytes.Buffer
	var payload Addr

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}

	//更新已知节点集和，并向已知节点集合的节点请求区块信息
	KnownNodes = append(KnownNodes, payload.AddrList...)
	fmt.Printf("there are %d known nodes\n", len(KnownNodes))
	RequestBlocks()
}

// 向指定节点发送本地已知节点集合信息
func SendAddr(address string) {
	nodes := Addr{KnownNodes}
	nodes.AddrList = append(nodes.AddrList, ledger.Config().Addr)
	payload := utils.GobEncode(nodes)
	request := append(utils.CmdToBytes("addr", commandLength), payload...)

	SendData(address, request)
}
