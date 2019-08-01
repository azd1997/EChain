/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/26 18:38
* @Description: The file is for
***********************************************************************/

package cmd

import (
	"EChain/enode/ledger"
	"EChain/enode/ledger/account"
	"EChain/enode/ledger/blockchain"
	"EChain/enode/ledger/statesync/p2p"
	"EChain/enode/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ledgerCmd = &cobra.Command{
	Use:   "ledger",
	Short: "区块链账本",
	Long:  `Echain节点间区块链账本，用以记录终端设备及网关节点行为记录`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("区块链账本启动")
		// 不带参数。若执行ledger，则检查本地节点IP是否对应有区块链数据库存在
		// （不论是不是自己创建的），如有，则继续区块链，如无，创建区块链

	},
}

var createChainCmd = &cobra.Command{
	Use:   "createchain",
	Short: "创建区块链账本",
	Long:  `网关节点第一次启动时创建区块链账本，这是由某一台指定节点执行的`,
	Args:cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始区块链账本创建")
		// 创建区块链
		creatChain()
	},

}

var createAccountCmd = &cobra.Command{
	Use:   "createaccount",
	Short: "创建账户地址",
	Long:  `创建一个新的账户地址，并写入账户文件`,
	Args:cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("生成新账户...")
		createAccount()
	},
}

var getBalanceCmd = &cobra.Command{
	Use:   "getbalance",
	Short: "查询账户地址余额",
	Long:  `查询给定账户地址链上总余额，并打印输出`,
	Args:cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("查询账户余额...")
		getBalance(args[0])
	},
}


var listAddressCmd = &cobra.Command{
	Use:   "listaddress",
	Short: "列出所有账户地址",
	Long:  `在控制台打印所有账户地址`,
	Args:cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("列举所有账户地址...")
		listAddress()
	},
}

var printChainCmd = &cobra.Command{
	Use:   "printchain",
	Short: "打印区块链",
	Long:  `打印当前区块链全部区块信息，或根据查询条件打印`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("打印区块链信息...")
		// printchain 打印所有
		// printchain --hash= xxxx 根据区块哈希打印单个区块信息
		// printchain --start= 2 --interval= 1(default) --end= 4 根据区块高度查询并打印信息
		// TODO:实现以上

		printChain()

	},
}

var startNodeCmd = &cobra.Command{
	Use:   "startnode",
	Short: "启动账本节点服务器",
	Long:  `根据配置文件启动账本节点，监听p2p网络请求并响应`,
	Args:cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("启动区块链账本...")

		startNode()
	},
}



func creatChain() {

	// 获取配置
	var cc = ledger.Config()

	// 检查账户地址有效性
	if !account.ValidateAccount(cc.MinerAddress) {
		log.Fatalln(ledger.ErrInvalidAccountAddress)
	}

	// 创建区块链
	chain, err := blockchain.InitChain(cc)	// 在InitChain中会检查区块链是否已存在，存在会失败。
	if err != nil {
		log.Panic(err)
	}
	defer chain.Db.Close()  // 注册及时关闭

	log.Println("创建区块链成功！")
}

func continueChain() {

	// 用指定账户地址替换配置文件中账户地址
	var cc = ledger.Config()

	// 检查账户地址有效性
	if !account.ValidateAccount(cc.MinerAddress) {
		log.Fatalln(ledger.ErrInvalidAccountAddress)
	}

	// 创建区块链
	chain, err := blockchain.ContinueChain(cc)	// 在InitChain中会检查区块链是否已存在，存在会失败。
	if err != nil {
		log.Panic(err)
	}
	defer chain.Db.Close()  // 注册及时关闭

	log.Println("继续区块链成功！")
}

func createAccount() {

	var cc = ledger.Config()

	//创造钱包集对象
	accounts, _ := account.NewAccountsFromFile(cc.Addr)
	//向钱包集新增一个钱包并保存到文件去
	address := accounts.AddAccount()
	accounts.SaveFile(cc.Addr)

	log.Printf("新生成账户地址: %s\n", address)
}

func getBalance(address string) {

	var cc = ledger.Config()

	if !account.ValidateAccount(address) {
		log.Fatalln(ledger.ErrInvalidAccountAddress)
	}

	chain, err := blockchain.ContinueChain(cc)
	if err != nil {
		log.Panic(err) //panic才会执行defer
	}
	defer chain.Db.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	// TODO:基于账户模型建立余额制，查找账户地址余额

	log.Printf("%s 账户余额: %d\n", address, balance)

}

func listAddress() {

	var cc = ledger.Config()

	accounts, _ := account.NewAccountsFromFile(cc.Addr)
	addresses := accounts.GetAllAddress()

	for _, address := range addresses {
		fmt.Println(address)
	}
}

func printChain() {

	var cc = ledger.Config()

	chain, err := blockchain.ContinueChain(cc)
	if err != nil {
		log.Panic(err)
	}
	defer chain.Db.Close()

	iter := chain.Iterator()

	for {
		block, err := iter.Next()
		if err != nil {
			continue  // 跳过就好
		}

		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("TxsHash: %x\n", block.HashTxs())
		fmt.Printf("Hash: %x\n", block.Hash)

		// TODO:区块的检验
		//pow := blockchain.NewProof(block)
		//fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))

		// 打印区块内事务信息
		for _, tx := range block.Txs {
			fmt.Println(tx)
		}
		fmt.Println()

		if len(block.PrevHash) == 0 { //创世区块PrevHash设为0
			break
		}
	}
}

func startNode() {

	var cc = ledger.Config()

	fmt.Printf("启动区块链账本节点 %s\n", cc.Addr)

	if len(cc.MinerAddress) > 0 {
		if account.ValidateAccount(cc.MinerAddress) {
			log.Println("正在挖矿记账，挖矿账户为： ", cc.MinerAddress)
		} else {
			log.Fatalln("错误的挖矿地址!")
		}
	}
	p2p.StartServer(cc)
}