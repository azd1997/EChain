/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/21 10:55
* @Description: ledger部分的配置参数
***********************************************************************/

package ledger

import (
	"EChain/enode/conf"
	"github.com/spf13/viper"
	"log"
)

type Conf struct {

	Addr string

	GenesisMsg string
	DbEngine string
	DbPathTemp string

	AccountsFilePathTemp string
	MinerAddress string

	SyncStateMethod string
}

func Config() *Conf {
	return &Conf{
		Addr:                 viper.GetString("ledgerConfig.nodeAddr"),
		GenesisMsg:           viper.GetString("ledgerConfig.blockChain.genesisMsg"),
		DbEngine:             viper.GetString("ledgerConfig.blockChain.dbEngine"),
		DbPathTemp:           viper.GetString("ledgerConfig.blockChain.dbPathTemp"),
		AccountsFilePathTemp: viper.GetString("ledgerConfig.account.filePathTemp"),
		MinerAddress:         viper.GetString("ledgerConfig.account.minerAddress"),
		SyncStateMethod:      viper.GetString("ledgerConfig.syncState.method"),
	}
}



//func Config() *Conf {
//	return &Conf{
//		Addr:                 conf.E_config.Ledger.Addr,
//		GenesisMsg:           conf.E_config.Ledger.BC.GenesisMsg,
//		DbEngine:             conf.E_config.Ledger.BC.DbEngine,
//		DbPathTemp:           conf.E_config.Ledger.BC.DbPathTemp,
//		AccountsFilePathTemp: conf.E_config.Ledger.Account.FilePathTemp,
//		MinerAddress:         conf.E_config.Ledger.Account.MinerAddress,
//		SyncStateMethod:      conf.E_config.Ledger.SyncState.Method,
//	}
//}




func ConfigTest() *Conf {
	return &Conf{
		Addr:                 conf.E_config.Ledger.Addr,
		GenesisMsg:           conf.E_config.Ledger.BC.GenesisMsg,
		DbEngine:             conf.E_config.Ledger.BC.DbEngine,
		DbPathTemp:           conf.E_config.Ledger.BC.DbPathTempTest,
		AccountsFilePathTemp: conf.E_config.Ledger.Account.FilePathTemp,
		MinerAddress:         conf.E_config.Ledger.Account.MinerAddress,
		SyncStateMethod:      conf.E_config.Ledger.SyncState.Method,
	}
}

// 初始化json配置 config.E_config
func InitConfigForTest() {
	const configFile = "../../config/enode.yaml"

	if err := conf.InitConfigByYaml(configFile); err != nil {
		log.Panic(err)
	}
}