package p2p

import (
	"EChain/enode/ledger"
	"EChain/enode/ledger/blockchain"
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
)

type GetData struct {
	AddrFrom string
	Type     string
	ID       []byte
}

//处理获取数据请求
func HandleGetData(request []byte, chain *blockchain.Chain) {

	//获取request中的内容
	var buff bytes.Buffer
	var payload GetData

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}

	//getdata有获取区块和获取交易两种情况

	if payload.Type == "block" {
		block, err := chain.GetBlock([]byte(payload.ID))
		if err != nil {
			log.Fatal(err)
		}

		//向给自己发请求的节点发送单个区块
		SendBlock(payload.AddrFrom, block)
	}

	if payload.Type == "tx" {
		txID := hex.EncodeToString(payload.ID)
		tx := memoryPool[txID]

		SendTx(payload.AddrFrom, &tx)
	}

}

/*发送获取数据的请求*/
func SendGetData(address, kind string, id []byte) {
	payload := GobEncode(GetData{ledger.Config().Addr, kind, id})
	request := append(CmdToBytes("getdata"), payload...)

	SendData(address, request)
}

func SendData(addr string, data []byte) {
	//向addr发起tcp连接
	conn, err := net.Dial(protocol, addr)

	//连接不可用，则更新已知节点集
	if err != nil {
		fmt.Printf("%s is not available\n", addr)
		var updatedNodes []string

		for _, node := range KnownNodes {
			if node != addr {
				updatedNodes = append(updatedNodes, node)
			}
		}

		KnownNodes = updatedNodes

		return
	}

	defer conn.Close()

	//将data []byte复制一份通过conn发给对方
	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
}
