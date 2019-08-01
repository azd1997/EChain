package p2p

import (
	"EChain/enode/ledger"
	"EChain/enode/ledger/blockchain"
	"EChain/enode/utils"
	"bytes"
	"encoding/gob"
	"log"
)

//Version主要用来处理最长合法链问题
type Version struct {
	Version    int
	BestHeight int
	AddrFrom   string
}

//接收到Version请求时
func HandleVersion(request []byte, chain *blockchain.Chain) {
	//将request中command以后内容解码并写入payload
	var buff bytes.Buffer
	var payload Version

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}

	//取当前区块链最长高度；及payload中最长高度
	//注意这里payload代表的是收到的其他节点发过来的版本信息（其中维护了最长高度的等信息）
	//可见，该方法主要用于解决最长合法链共识
	bestHeight, err := chain.GetHeight()
	if err != nil {
		log.Fatal(err)
	}
	otherHeight := payload.BestHeight

	//本地区块链不是最长链则向对方发送自己的地址，请求区块
	if bestHeight < int64(otherHeight) {
		SendGetBlocks(payload.AddrFrom)
	} else if bestHeight > int64(otherHeight) {
		//本地区块链若为最长合法链而对方不是，则给对方发一个Version
		SendVersion(payload.AddrFrom, chain)
	}

	//如果来源节点不是已知节点集成员，将其加入
	if !NodeIsKnown(payload.AddrFrom) {
		KnownNodes = append(KnownNodes, payload.AddrFrom)
	}
}

/*向某一地址发送版本信息*/
func SendVersion(addr string, chain *blockchain.Chain) {
	//打包版本数据并发送
	bestHeight, err := chain.GetHeight()
	if err != nil {
		log.Fatal(err)
	}
	payload := utils.GobEncode(Version{version, int(bestHeight), ledger.Config().Addr})

	request := append(utils.CmdToBytes("version", commandLength), payload...)

	SendData(addr, request)
}
