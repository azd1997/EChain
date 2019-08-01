/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/16 21:05
* @Description: The file is for
***********************************************************************/

package blockchain

import "EChain/enode/database"

type ChainIterator struct {
	CurrentHash []byte
	Db database.Database
}

// 从后向前遍历，获取当前区块，并更新迭代器CurrentHash
func (iter *ChainIterator) Next() (*Block, error) {
	var (
		block *Block
		blockBytes []byte
		err error
	)

	// 显然，这里iter。Db需要传入一个连接状态的Db。由Chain.Db赋值
	if blockBytes, err = iter.Db.Get(iter.CurrentHash); err != nil {
		return nil, err
	}
	if block, err = DeserializeBlock(blockBytes); err != nil {
		return nil, err
	}
	iter.CurrentHash = block.PrevHash

	return block, nil
}