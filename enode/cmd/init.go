/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/26 13:15
* @Description: 初始化
***********************************************************************/

package cmd

import (
	log "github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
	"runtime"
)



func init() {
	// 初始化
	// cobra.OnInitialize(initENode)

	// initENode()

	initEnv()

	initFlags()

	rootCmd.AddCommand(versionCmd)

	// loggerCmd
	rootCmd.AddCommand(loggerCmd)

	// ledgerCmd
	rootCmd.AddCommand(ledgerCmd)
	ledgerCmd.AddCommand(createChainCmd)
	ledgerCmd.AddCommand(createAccountCmd)
	ledgerCmd.AddCommand(getBalanceCmd)
	ledgerCmd.AddCommand(listAddressCmd)
	ledgerCmd.AddCommand(printChainCmd)
	ledgerCmd.AddCommand(startNodeCmd)

	// monitorCmd
	rootCmd.AddCommand(monitorCmd)

	// endmgrCmd
	rootCmd.AddCommand(endmgrCmd)
	endmgrCmd.AddCommand(assignEndIdCmd)


}

// 初始化顺序：import -> 包级const/var -> init() -> main()


// 多线程
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("初始化多线程环境成功！")
}




