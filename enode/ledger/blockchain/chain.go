/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/16 14:28
* @Description: The file is for
***********************************************************************/

package blockchain

import (
	"EChain/enode/database"
	"EChain/enode/ledger"
	"EChain/enode/utils"
	"fmt"
	"log"
	"os"
	"runtime"
)

// Chain 存储最后一个区块哈希及数据库指针
type Chain struct {
	LastHash []byte
	Db database.Database	// TODO:该用指针吗？
}

// 初始化带有创世区块的区块链。创建者设置为假账户
// TODO:作为外部调用API
// 正常启动时，传入ledger.Config()；测试用ledger.ConfigTest()
func InitChain(cc *ledger.Conf) (*Chain, error) {
	var (
		path string
		db database.Database
		cbTx *Tx
		genesis *Block
		genesisBlockBytes []byte
		err error
	)

	// 检查区块链是否存在
	path = fmt.Sprintf(
		cc.DbPathTemp,
		cc.DbEngine,
		utils.IpToDir(cc.Addr))
	log.Println(path)

	// 创建该目录
	// 因为路径模板dbPathTemp有两个可变参数，所以必须通过MkdirAll来创建目录（其会创建所有必须的上级目录）
	if err := os.MkdirAll(path, os.ModeDir); err != nil {
		log.Fatal(err)
	}

	// 检查数据库文件是否存在
	if database.DbExists(cc.DbEngine, path) {
		// TODO:错误处理优化
		log.Println("区块链已存在...")
		// runtime.Goexit() 退出所在位置的go程
		// 扩展: os.exit() 退出进程，给进程分配的空间也会销毁
		runtime.Goexit()
	}

	// 打开数据库
	if db, err = database.OpenDatabaseWithRetry(cc.DbEngine, path); err != nil {
		return nil, err
	}
	// TODO：思考这里需不需要defer关闭
	// defer db.Close()

	// 生成创世区块
	if cbTx, err = CoinbaseTx(); err != nil {
		return nil, err
	}
	genesis = genesisBlock(cbTx)
	if genesisBlockBytes, err = genesis.Serialize(); err != nil {
		return nil, err
	}

	// 存入创世区块及lastHash
	if err = db.Set(genesis.Hash, genesisBlockBytes); err != nil {
		return nil, err
	}
	if err = db.Set([]byte("lh"), genesis.Hash); err != nil {
		return nil, err
	}

	// 返回Chain对象
	return &Chain{genesis.Hash, db}, nil
}

// 已有区块链的基础上继续区块链
// 正常启动时，传入ledger.Config()；测试用ledger.ConfigTest()
func ContinueChain(cc *ledger.Conf) (*Chain, error) {
	var (
		path string
		db database.Database
		lastHash []byte
		err error
	)

	// 检查区块链是否存在
	path = fmt.Sprintf(
		cc.DbPathTemp,
		cc.DbEngine,
		utils.IpToDir(cc.Addr))
	if !database.DbExists(cc.DbEngine, path) {
		fmt.Println("区块链不存在...")
		runtime.Goexit()
	}

	// 打开数据库
	if db, err = database.OpenDatabaseWithRetry(cc.DbEngine, path); err != nil {
		return nil, err
	}
	// TODO：思考这里需不需要defer关闭
	// 区块链部分在运行时一直需要进行数据库连接？只有需要存块和取块时才需要。
	// 要是选择每次存取再打开数据库，需要在database接口再定义一个open方法
	// TODO:暂时使用长连接，了解下连接池
	// defer db.Close()

	if lastHash, err = db.Get([]byte("lh")); err != nil {
		return nil, err
	}

	// 返回Chain对象
	return &Chain{lastHash, db}, nil
}

// “挖矿”
func (c *Chain) MineBlock(txs []*Tx) (*Block, error) {

	var (
		tx *Tx
		lastHash, lastBlockBytes, newBlockBytes []byte
		lastBlock, newBlock *Block
		err error
	)

	// 本地验证事务的有效性
	for _, tx = range txs {
		if !c.VerifyTx(tx) {
			log.Panic("无效的请求连接事务")
		}
	}

	// 获取区块链中最新区块
	if lastHash, err = c.Db.Get([]byte("lh")); err != nil {
		return nil, err
	}
	if lastBlockBytes, err = c.Db.Get(lastHash); err != nil {
		return nil, err
	}
	if lastBlock, err = DeserializeBlock(lastBlockBytes); err != nil {
		return nil, err
	}

	// 打包新区块
	newBlock = NewBlock(txs, lastHash, lastBlock.Height + 1)
	if newBlockBytes, err = newBlock.Serialize(); err != nil {
		return nil, err
	}

	// 存入数据库并更新
	if err = c.Db.Set(newBlock.Hash, newBlockBytes); err != nil {
		return nil, err
	}
	if err = c.Db.Set([]byte("lh"), newBlock.Hash); err != nil {
		return nil, err
	}

	return newBlock, nil
}

// 将区块链网络的区块更新入本地区块链，或者说同步区块链
func (c *Chain) AddBlock(b *Block) error {

	var (
		blockBytes, lastBlockBytes, lastHash []byte
		lastBlock *Block
		err error
	)

	// 检查该区块是否已经存在本地，若存在，啥也不做
	if _, err = c.Db.Get(b.Hash); err == nil {
		return nil
	}

	// 如果区块在本地区块链中找不到，则先直接存进来
	if blockBytes, err = b.Serialize(); err != nil {
		return err
	}
	if err = c.Db.Set(b.Hash, blockBytes); err != nil {
		return err
	}

	// 取出区块链中当前“lh”对应的区块
	if lastHash, err = c.Db.Get([]byte("lh")); err != nil {
		return err
	}
	if lastBlockBytes, err = c.Db.Get(lastHash); err != nil {
		return err
	}
	if lastBlock, err = DeserializeBlock(lastBlockBytes); err != nil {
		return err
	}

	// 比较新加入的区块与所取出的最后区块高度
	if b.Height > lastBlock.Height {
		// 更新“lh”
		// TODO:思考为什么直接更新了，要是新加的区块比最后区块高好几个区块呢？
		if err = c.Db.Set([]byte("lh"), b.Hash); err != nil {
			return err
		}
		c.LastHash = b.Hash
	}

	return nil
}

// 从区块链中获取（查询）区块
func (c *Chain) GetBlock(blockHash []byte) (*Block, error) {

	var (
		block *Block
		blockBytes []byte
		err error
	)

	if blockBytes, err = c.Db.Get(blockHash); err != nil {
		return nil, err
	}
	if block, err = DeserializeBlock(blockBytes); err != nil {
		return nil, err
	}

	return block, nil
}

// 获取区块链所有区块哈希集合，用以快速验证不同节点间区块链的一致性
func (c *Chain) GetAllBlockHashes() ([][]byte, error) {
	var (
		blockHashes [][]byte
		block *Block
		iter *ChainIterator
		err error
	)

	iter = c.Iterator()
	for {
		if block, err = iter.Next(); err != nil {
			return nil, err
		}
		blockHashes = append(blockHashes, block.Hash)
		if len(block.PrevHash) == 0 {
			break
		}
	}

	return blockHashes, nil
}

// 获取区块链高度
func (c *Chain) GetHeight() (int64, error) {

	var (
		lastBlock *Block
		lastHash, lastBlockBytes []byte
		err error
	)

	if lastHash, err = c.Db.Get([]byte("lh")); err != nil {
		return -1, err
	}
	if lastBlockBytes, err = c.Db.Get(lastHash); err != nil {
		return -2, err
	}
	if lastBlock, err = DeserializeBlock(lastBlockBytes); err != nil {
		return -3, err
	}
	return lastBlock.Height, nil
}

// 获取某个网关节点账户的“余额”
func (c *Chain) GetNodeBalance(nodeAccount string) int64 {

	// TODO
	return 0
}

// 获取某个终端设备账户的“余额”
func (c *Chain) GetEndBalance(endAccount string) int64 {

	// TODO
	return 0
}

// 获取区块链迭代器ChainIterator
func (c *Chain) Iterator() *ChainIterator {
	return &ChainIterator{c.LastHash, c.Db}
}

// 验证接入事务的合法性
func (c *Chain) VerifyTx(tx *Tx) bool {

	// TODO:如何验证接入事务

	return true
}



