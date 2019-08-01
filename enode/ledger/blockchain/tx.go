package blockchain

import (
	"EChain/enode/ends"
	"EChain/enode/ledger"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

//这里的tx用于接收来自外部数据采集模块传输的数据。双方通过管道或者其他的结构体等发现
//tx 事务  （不使用区块链的交易名称）

const (
	Online int = iota	//终端设备接入事件
	Offline 			//终端设备
	Coinbase
)

type Tx struct {
	Id []byte	// hash
	Type int
	Score int
	Timestamp int64
	End *ends.End
	Manager string
	Msg []byte	//存放消息
	//TODO:增加该终端该次事务的具体数据信息，比如本次接入的认证时间内的请求数据
}

// 对Tx序列化为[]byte
func (tx *Tx) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(tx)

	return buf.Bytes(), err
}

// 对Tx[]byte反序列化为Tx
func DeserializeTx(data []byte) (*Tx, error) {
	var tx Tx
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&tx)
	if err != nil {
		return &Tx{}, err
	}

	return &tx, nil
}

// 创造一个交易事务Tx
func NewTx(typ int, msg string) (*Tx, error) {
	tx :=  &Tx{
		Type:      typ,
		Timestamp: time.Now().Unix(),
		Manager:   ledger.Config().Addr,
		Msg:       []byte(msg),
	}
	id, err := tx.hash()
	if err != nil {
		return nil, err
	}
	tx.Id = id
	return tx, nil
}

// 创世区块的唯一交易，指定一些信息
func CoinbaseTx() (*Tx, error) {
	tx, err := NewTx(Coinbase, ledger.Config().GenesisMsg)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// 判断事务是否是Coinbase事务
func (tx *Tx) IsCoinbase() bool {
	return tx.Type == Coinbase
}

// 为未生成Id的事务生成事务哈希作为Id
func (tx *Tx) hash() ([]byte, error) {
	txBytes, err := tx.Serialize()
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(txBytes)
	return hash[:], nil
}
