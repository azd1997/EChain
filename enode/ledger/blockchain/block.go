/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/16 14:28
* @Description: The file is for
***********************************************************************/

package blockchain

import (
	"EChain/enode/ledger"
	"EChain/enode/utils"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

//网关将某个时间段手机的所有终端设备连接请求打包发出，
//所有网关在收到区块后，根据区块链中该终端设备记录生成评分进行投票，过2/3则通过
//通过后，网关开始将该设备数据正式存入设备数据库，不通过则挂掉连接协程，并删除其缓存
//只保留日志记录

type Header struct {
	Timestamp int64
	Height int64
	PrevHash []byte
	Hash []byte
	Recorder string
}

type Block struct {
	Header
	Txs []*Tx
}

// 对区块内事务集，构建MerkleTree，再返回Merkle根哈希
func (b *Block) HashTxs() []byte {
	var (
		txsBytes [][]byte
		txBytes []byte
		err error
	)

	for _, tx := range b.Txs {
		txBytes, err = tx.Serialize()
		if err != nil {
			continue	//如果出错就跳过，影响不大
		}
		txsBytes = append(txsBytes, txBytes)
	}

	tree := NewMerkleTree(txsBytes)

	return tree.RootNode.Data	//对于非叶节点，Data存的是哈希
}

// 生成区块的哈希值
func (b *Block) hash() []byte {
	// 转为[]byte拼接，并以[]byte{}分隔
	data := bytes.Join(
		[][]byte{
			utils.Int64ToBytes(b.Timestamp),
			utils.Int64ToBytes(b.Height),
			b.PrevHash,
			[]byte(b.Recorder),
			b.HashTxs(),
		},
		[]byte{})
	hash := sha256.Sum256(data)
	return hash[:]
}

// 新建区块
func NewBlock(txs []*Tx, prevHash []byte, height int64) *Block {
	header := &Header{
		Timestamp: time.Now().Unix(),
		Height:    height,
		PrevHash:  prevHash,
		Recorder:  ledger.Config().Addr,
	}
	block := &Block{
		*header,
		txs,
	}
	block.Hash = block.hash()

	return block
}

// 创世区块
func genesisBlock(coinbase *Tx) *Block {
	return NewBlock([]*Tx{coinbase}, []byte{}, 0)
}

// 区块序列化
func (b *Block) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 区块反序列化
func DeserializeBlock(data []byte) (*Block, error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

