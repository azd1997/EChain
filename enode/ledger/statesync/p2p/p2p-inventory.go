package p2p

import (
	"EChain/enode/ledger"
	"EChain/enode/ledger/blockchain"
	"EChain/enode/utils"
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

type Inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

/*向某节点发送存证*/
func SendInv(address, kind string, items [][]byte) {
	inventory := Inv{ledger.Config().Addr, kind, items}
	payload := utils.GobEncode(inventory)
	request := append(utils.CmdToBytes("inv", commandLength), payload...)

	SendData(address, request)
}

//处理节点接收到来自其他节点的存证，存证有区块存证和交易存证两种
func HandleInv(request []byte, chain *blockchain.Chain) {
	//获取request中的内容
	var buff bytes.Buffer
	var payload Inv

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received inventory with %d %s\n", len(payload.Items), payload.Type)

	if payload.Type == "block" {
		//收到区块存证，则向对方请求这个区块的数据
		blockInTransit = payload.Items

		blockHash := payload.Items[0]
		SendGetData(payload.AddrFrom, "block", blockHash)

		//将blockInTransit中不是payload中那块的区块哈希加入newInTransit
		//再用newInTransit更新blockInTransit
		var newInTransit [][]byte
		for _, b := range blockInTransit {
			if bytes.Compare(b, blockHash) != 0 {
				newInTransit = append(newInTransit, b)
			}
		}
		blockInTransit = newInTransit
	}

	if payload.Type == "tx" {
		txID := payload.Items[0]

		//如果本地内存池中没有对方发来存证的交易，则向对方请求交易数据
		if memoryPool[hex.EncodeToString(txID)].Id == nil {
			SendGetData(payload.AddrFrom, "tx", txID)
		}
	}

}
