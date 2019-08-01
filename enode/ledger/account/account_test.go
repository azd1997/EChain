/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 16:56
* @Description: The file is for
***********************************************************************/

package account

import (
	"EChain/enode/ledger"
	"log"
	"testing"
)

// 测试account.go。测试生成一个账户及账户地址及验证地址。
func TestAccount(t *testing.T) {
	acc := NewAccount()
	accAddr := acc.Address()
	if !ValidateAccount(string(accAddr)) {
		log.Fatal("账户地址验证不通过！")
	}
}

// 测试Accounts.go。 测试将多个账户存入文件
func TestAccountsToFile(t *testing.T) {

	// 加载总配置
	ledger.InitConfigForTest()

	var accounts Accounts
	accounts.Map = make(map[string]*Account)	// 记得给map初始化，不然它是空指针

	// 添加三个账户
	accounts.AddAccount()
	accounts.AddAccount()
	accounts.AddAccount()

	// 存入文件
	accounts.SaveFile(ledger.Config().Addr)
}

// 测试Accounts.go。 测试将从文件读取账户及读取账户文件内账户的地址
func TestAccountsFromFile(t *testing.T) {

	TestAccountsToFile(t)

	// 加载总配置
	ledger.InitConfigForTest()

	// 加载账户文件，得到Accounts对象
	accounts, err := NewAccountsFromFile(ledger.Config().Addr)
	if err != nil {
		log.Fatal(err)
	}

	allAddresses := accounts.GetAllAddress()
	for _, a := range allAddresses {
		log.Println(a)
	}
}