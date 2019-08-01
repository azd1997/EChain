/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/8/1 20:44
* @Description: The file is for
***********************************************************************/

package monitor

import "github.com/spf13/viper"

type Conf struct {

	serverIpAddress string
	serverReadTimeout int
	serverWriteTimeout int
	serverWebRoot string
}

func Config() *Conf {
	return &Conf{
		serverIpAddress:                 viper.GetString("monitorConfig.serverIpAddress"),
		serverReadTimeout:           viper.GetInt("monitorConfig.serverReadTimeout"),
		serverWriteTimeout:             viper.GetInt("monitorConfig.serverWriteTimeout"),
		serverWebRoot:           viper.GetString("monitorConfig.serverWebRoot"),
	}
}