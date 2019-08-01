/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/8/1 20:26
* @Description: The file is for
***********************************************************************/

package endmgr

import "github.com/spf13/viper"

type Conf struct {

	manager string
	enodeIpAddress string
}

func Config() *Conf {
	return &Conf{
		manager:                 viper.GetString("endmgrConfig.manager"),
		enodeIpAddress:           viper.GetString("endmgrConfig.enodeIpAddress"),
	}
}