/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/26 11:01
* @Description: The file is for
***********************************************************************/

package conf

import (
	"EChain/enode/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// 注意conf包viper.go实现了配置文件的读取。其余文件可以忽略。

// 初始化配置
func InitConfigByViper(configFile string) (err error) {

	// 1.检查文件是否存在
	if exists, err := utils.FileExists(configFile); !exists {
		// 文件不存在时，err可能为空，不管怎样都返回。为空和不为空代表两种出错情况
		// 在这里必须把err为nil处理掉，不然外部不知道怎么办。但是在这里我们只想知道有没有配置文件而已
		// err = nil 的时候意味着错误是文件路径不存在，所以返回一个这个错误
		if err == nil {
			err = os.ErrNotExist
		}
		// 否则是未知错误
		return err
		// TODO:不存在时是否写入一个默认配置文件到该路径？？
	}

	// 2.viper读取配置文件
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		//log.Println("无法读取配置：", err)
		//os.Exit(1)
		return err
	}

	log.Println("初始化配置系统成功！")
	return nil
}


// 初始化配置
//func initConfig() {
//	// Don't forget to read config either from cfgFile or from home directory!
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		if err != nil {
//			log.Println(err)
//			os.Exit(1)
//		}
//
//		// Search config in home directory with name ".cobra" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".cobra")
//	}
//
//	if err := viper.ReadInConfig(); err != nil {
//		log.Println("无法读取配置：", err)
//		os.Exit(1)
//	}
//
//	log.Println("Viper读取配置文件成功！")
//	log.Print(viper.AllKeys())
//	log.Print(viper.AllSettings())
//	log.Println(viper.GetString("loggerConfig.mongodbUri"))
//
//}


// viper包本身内部就实现了一个单例，在我们自己的应用里只需要调用viper.func，实际用到的是v.method。
// 当我们需要多个viper实例时，需要使用viper.New()创造。（一般不会这么做）
