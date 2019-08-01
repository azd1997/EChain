package p2p

import (
	"EChain/enode/ledger"
	"EChain/enode/ledger/blockchain"
	"EChain/enode/utils"
	"bytes"
	"encoding/gob"
	"log"
)

type GetBlocks struct {
	AddrFrom string
}

// 处理获取全部区块（哈希）存证请求
func HandleGetBlocks(request []byte, chain *blockchain.Chain) {
	//获取request内容
	var buff bytes.Buffer
	var payload GetBlocks

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}

	//向发请求的节点发送存证，说自己存了所有的区块
	blocks, err := chain.GetAllBlockHashes()
	SendInv(payload.AddrFrom, "block", blocks)
}

// 发送获取区块请求
func SendGetBlocks(address string) {
	payload := utils.GobEncode(GetBlocks{ledger.Config().Addr})
	request := append(utils.CmdToBytes("getblocks", commandLength), payload...)

	SendData(address, request)
}

/*向已知节点集和中的所有节点发送GetBlocks的请求*/
func RequestBlocks() {
	for _, node := range KnownNodes {
		SendGetBlocks(node)
	}
}
