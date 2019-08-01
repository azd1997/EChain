/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/26 11:37
* @Description: The file is for
***********************************************************************/

package cmd

import (
	log "github.com/sirupsen/logrus"
	//flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	configFile     string
	projectBase string
	userLicense string
)

// flags定义及解析
func initFlags() {
	//rootCmd.
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "github.com/azd1997/", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "a", "enode Eiger", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "apache", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")

	_ = viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	_ = viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	_ = viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	viper.SetDefault("author", "Eiger <374192922@qq.com>")
	viper.SetDefault("license", "apache")

	// ledger
	getBalanceCmd.Flags().StringP("Address", "a", "", "填入账户地址，根据账户地址查询其余额 (参数必需)")



	log.Println("初始化命令行参数集成功！")


}