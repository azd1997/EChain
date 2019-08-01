package p2p

import (
	"EChain/enode/ledger"
	"EChain/enode/ledger/blockchain"
	"EChain/enode/utils"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Block struct {
	AddrFrom string
	Block    []byte
}

// 处理接收到区块时
func HandleBlock(request []byte, chain *blockchain.Chain) {
	//获取request内容
	var buff bytes.Buffer
	var payload Block

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}

	//将接收到的区块添加到区块链中
	blockData := payload.Block
	block, err := blockchain.DeserializeBlock(blockData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Received a new block!")
	if err = chain.AddBlock(block); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added block %x\n", block.Hash)

	//若blockInTransit中还有内容，那么据此继续向对方请求区块数据
	//这表示只要blockInTransit非空，就会不断请求，对方不断返回区块，自己不断处理区块
	if len(blockInTransit) > 0 {
		blockHash := blockInTransit[0]
		SendGetData(payload.AddrFrom, "block", blockHash)

		blockInTransit = blockInTransit[1:]
	} else {
		//更新未花费输出集
		/*UTXOSet := blockchain.UTXOSet{chain}
		UTXOSet.Reindex()*/
	}
}

// 向某地址发送区块
func SendBlock(addr string, b *blockchain.Block) {
	bBytes, err := b.Serialize()
	if err != nil {
		log.Fatal(err)
	}
	data := Block{ledger.Config().Addr, bBytes}
	payload := utils.GobEncode(data)
	request := append(utils.CmdToBytes("block", commandLength), payload...)

	SendData(addr, request)
}
