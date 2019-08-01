package cli

import (
	"EChain/enode/ledger"
	"EChain/enode/ledger/account"
	"fmt"

)

func (cli *CommandLine) createAccount(cc *ledger.Conf) {

	//创造钱包集对象
	accounts, _ := account.NewAccountsFromFile(cc.Addr)
	//向钱包集新增一个钱包并保存到文件去
	address := accounts.AddAccount()
	accounts.SaveFile(cc.Addr)

	fmt.Printf("New address is: %s\n", address)

}
