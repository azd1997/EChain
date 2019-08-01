package cli

import (
	"EChain/enode/ledger/account"
	"EChain/enode/ledger/blockchain"
	"EChain/enode/utils"
	"fmt"

	"log"
)

/*获取账户余额*/
func (cli *CommandLine) getBalance(address, nodeID string) {

	if !account.ValidateAccount(address) {
		log.Panic("Address is not valid")
	}

	chain, err := blockchain.ContinueChain(cc)

	//UTXOSet := blockchain.UTXOSet{chain}
	defer chain.Db.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	//UTXOs := UTXOSet.FindUnspentTransactions(pubKeyHash)

	//for _, out := range UTXOs {
	//	balance += out.Value
	//}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}
