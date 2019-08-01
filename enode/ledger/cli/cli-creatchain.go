package cli

import (
	"EChain/enode/ledger"
	"EChain/enode/ledger/account"
	"EChain/enode/ledger/blockchain"
	"fmt"
	"log"
)

/*创建区块链，其创世区块coinbase交易地址给定*/
func (cli *CommandLine) createBlockChain() {

	var cc = ledger.Config()

	if !account.ValidateAccount(cc.MinerAddress) {
		log.Panic("Address is not valid")
	}


	chain, err := blockchain.InitChain(cc)
	if err != nil {
		log.Fatal(err)
	}
	defer chain.Db.Close()

	////UTXO
	//UTXOSet := blockchain.UTXOSet{chain}
	//UTXOSet.Reindex()

	fmt.Println("Finished!")
}
