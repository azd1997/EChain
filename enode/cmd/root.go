/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/21 17:11
* @Description: The file is for
***********************************************************************/

package cmd

import (
	"EChain/enode/conf"
	"EChain/enode/ledger"
	"EChain/enode/logger"
	"fmt"
	//"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "enode",
	Short: "enode is a blockchain-based iot platform",
	Long: `A Fast and Flexible Static Site Generator built with
               love by spf13 and friends in Go.
               Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		//fmt.Println(configFile)
		//
		////cmd.ParseFlags(args)
		//fmt.Println("解析参数")
		//fl := cmd.PersistentFlags().Lookup("author")
		//fmt.Println(fl.Value)
		//fmt.Println(configFile)

		// Run流程
		// 1. 根据--config（configFile）读取配置
		// 2. 初始化各项自定义服务

		var err error
		// 2. 配置
		if err = conf.InitConfigByViper(configFile); err != nil {
			goto ERR
		}

		if err = logger.InitLogger(); err != nil {
			goto ERR
		}

		if err = ledger.InitLedger(); err != nil {
			goto ERR
		}

		return
	ERR:
		log.Error(err)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

