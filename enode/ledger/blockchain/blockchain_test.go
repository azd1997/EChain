/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/17 20:22
* @Description: The file is for
***********************************************************************/

package blockchain

import (
	"EChain/enode/ledger"
	"EChain/enode/utils"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)



// 初始化Chain测试
func initChainWithCloseAndListFilesForTest() (path string) {

	// 初始化json配置
	ledger.InitConfigForTest()

	var configTest = ledger.ConfigTest()

	chain, err := InitChain(configTest)
	if err != nil {
		log.Fatal(err)
	}
	// defer chain.Db.Close()

	log.Println(chain)

	// 列出测试数据库下所有文件名
	path = fmt.Sprintf(
		configTest.DbPathTemp,
		configTest.DbEngine,
		utils.IpToDir(configTest.Addr))
	log.Printf("列出path:%s下文件：\n", path)
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Print(info.Name(), "  ")
		return nil
	})
	log.Println()
	if err != nil {
		log.Fatal(err)
	}

	// 手动关闭数据库
	if err = chain.Db.Close(); err != nil {
		log.Fatal(err)
	}

	return
}

// 删除测试的区块链数据库文件
func delChainDatabaseForTest(path string) {
	// 删除测试数据库，方便重复测试及以后的压力测试
	// go操作文件/文件夹：https://www.cnblogs.com/jkko123/p/7026117.html
	// 常规的Remove/RemoveAll删除目录只能删除没有文件的目录/多级目录
	// 这里使用调用shell命令的方式进行强制删除
	// rd/s/q 盘符:\某个文件夹 （强制删除文件夹及其内所有文件）
	// del/f/s/q 盘符:\文件名 （强制删除文件，文件名必须加后缀名）

	// 1.获取绝对路径 E:\GO\XProjects\EChain\ENode\ledger\tmp\blocks_test\badger\blocks_127_0_0_1_9797
	absPath, _ := filepath.Abs(path)
	log.Println("数据库绝对路径为：", absPath)
	// 2.获取获取path的不可变部分：../tmp/blocks_test/badger/blocks_127_0_0_1_9797 => ../tmp/blocks_test/
	prefix := utils.GetSamePrefixFromTwoStrings(path, ledger.ConfigTest().DbPathTemp)
	log.Println("相对路径目录前缀：", prefix)
	// 3.获取path的可变部分：../tmp/blocks_test/badger/blocks_127_0_0_1_9797 => /badger/blocks_127_0_0_1_9797
	suffix := strings.TrimPrefix(path, prefix)
	suffix = strings.ReplaceAll(suffix, "/", "\\")  // badger/blocks_127_0_0_1_9797 => badger\blocks_127_0_0_1_9797
	suffix = "\\" + suffix  // badger\blocks_127_0_0_1_9797 => \badger\blocks_127_0_0_1_9797
	log.Println("绝对路径剪切后缀：", suffix)
	// 4.获取待删除的路径 E:\GO\XProjects\EChain\ENode\ledger\tmp\blocks_test
	delPath := strings. TrimSuffix(absPath, suffix)
	log.Println("待删除目录为：", delPath)
	// 5.延时， 再调用shell命令  cmd => C:\\Windows\\System32\\cmd.exe
	time.Sleep(10*time.Second)
	if err := exec.Command("cmd", "/C", "rd", "/S", "/Q", delPath).Start(); err != nil {
		log.Fatal(err)
	}
}

// 测试tx.go
func TestTxOperations(t *testing.T) {

	var (
		tx1, tx2, tx2Copy *Tx
		tx2Bytes, tx2CopyBytes []byte
		err error
	)

	// 初始化json配置
	ledger.InitConfigForTest()

	// 创建coinbase事务
	if tx1, err = CoinbaseTx(); err != nil {
		log.Fatal(err)
	}
	log.Println("tx1: ", tx1)
	// 创建一个普通事务
	if tx2, err = NewTx(Online, "测试一个接入事务"); err != nil {
		log.Fatal(err)
	}
	log.Println("tx2: ", tx2)
	// 测试tx2序列化与反序列化
	if tx2Bytes, err = tx2.Serialize(); err != nil {
		log.Fatal(err)
	}
	if tx2Copy, err = DeserializeTx(tx2Bytes); err != nil {
		log.Fatal(err)
	}
	log.Println("tx2Copy: ", tx2Copy)
	// 判断编解码前后是否一致
	if tx2CopyBytes, err = tx2Copy.Serialize(); err != nil {
		log.Fatal(err)
	}
	if sha256.Sum256(tx2Bytes) != sha256.Sum256(tx2CopyBytes) {
		log.Fatal("Tx编/解码前后不一致")
	}

	// 测试tx1，tx2是不是coinbaseTx
	log.Println("tx1是CoinbaseTx? ", tx1.IsCoinbase())
	log.Println("tx2是CoinbaseTx? ", tx2.IsCoinbase())
}

// 测试merkle.go
func TestMerkleOperations(t *testing.T) {
	// 测试叶节点生成
	leaf1 := NewMerkleNode(nil, nil, []byte("Tx数据1"))
	log.Println("leaf1: ", leaf1)
	leaf2 := NewMerkleNode(nil, nil, []byte("Tx数据2"))
	log.Println("leaf2: ", leaf2)

	// 测试其他节点生成
	node := NewMerkleNode(leaf1, leaf2, nil)
	log.Println("node: ", node)

	//测试树的生成，只有三个节点
	data := [][]byte{
		[]byte("Tx数据1"),
		[]byte("Tx数据2"),
	}
	tree := NewMerkleTree(data)
	log.Println("tree: ", tree)

}

// 测试block.go
func TestBlockOperations(t *testing.T) {

	var (
		cbTx *Tx
		genesis, genesis2 *Block
		genesisBytes []byte
		err error
	)

	// 初始化json配置
	ledger.InitConfigForTest()

	// 创造创世区块
	if cbTx, err = CoinbaseTx(); err != nil {
		log.Fatal(err)
	}
	genesis = genesisBlock(cbTx)
	log.Println(genesis)

	// 区块序列化
	if genesisBytes, err = genesis.Serialize(); err != nil {
		log.Fatal(err)
	}
	log.Println(genesisBytes)

	// 区块反序列化
	if genesis2, err = DeserializeBlock(genesisBytes); err != nil {
		log.Fatal(err)
	}
	log.Println(genesis2)

}

// 测试chain/InitChain()
func TestInitChain(t *testing.T) {

	// 初始化测试区块链，列举目录下文件，并关闭数据库连接
	relRealPath := initChainWithCloseAndListFilesForTest()

	// 删除测试数据库
	delChainDatabaseForTest(relRealPath)

	log.Println()   // 仅作换行分隔
}

// 测试chain/ContinueChain()
func TestContinueChain(t *testing.T) {

	// 初始化测试区块链，列举目录下文件，并关闭数据库连接
	relRealPath := initChainWithCloseAndListFilesForTest()

	// 继续区块链
	chain, err := ContinueChain(ledger.ConfigTest())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(chain)	// NOTICE:注意检查前后两次打印chain是否一致

	// 注意此时数据库没有关闭，但无所谓，反正会被删掉

	// 删除测试数据库
	delChainDatabaseForTest(relRealPath)

	log.Println()   // 仅作换行分隔
}