/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 22:55
* @Description: The file is for
***********************************************************************/

package conf

type ledgerConfig struct {
	Addr string `json:"nodeAddr" yaml:"nodeAddr" toml:"nodeAddr"`
	// 也可以匿名嵌入， 即直接BlockChain	`json:"blockChain"`.但这里不采用
	BC blockChain	`json:"blockChain" yaml:"blockChain" toml:"blockChain"`

	Account account `json:"account" yaml:"account" toml:"account"`

	SyncState syncState `json:"syncState" yaml:"syncState" toml:"syncState"`

}

//type ledgerConfig4Toml struct {
//	Ledger map[string]interface{}
//}

type blockChain struct {
	GenesisMsg string `json:"genesisMsg" yaml:"genesisMsg" toml:"genesisMsg"`
	DbEngine string `json:"dbEngine" yaml:"dbEngine" toml:"dbEngine"`
	DbPathTemp string `json:"dbPathTemp" yaml:"dbPathTemp" toml:"dbPathTemp"`
	DbPathTempTest string `json:"dbPathTempTest" yaml:"dbPathTempTest" toml:"dbPathTempTest"`
}

type account struct {
	FilePathTemp string `json:"filePathTemp" yaml:"filePathTemp" toml:"filePathTemp"`
	MinerAddress string `json:"minerAddress" yaml:"minerAddress" toml:"minerAddress"`
}

type syncState struct {
	// 区块链账本状态同步方法： "p2p"-p2p网络最长链同步、"etcd"-etcd同步
	Method string `json:"method" yaml:"method" toml:"method"`
}