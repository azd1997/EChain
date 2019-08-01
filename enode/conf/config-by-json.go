/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 21:56
* @Description: The file is for
***********************************************************************/

package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// 读取配置
func InitConfigByJson(configFile string) (err error) {

	var (
		content []byte
		config Config
	)

	// 读取配置文件，得到[]byte内容
	if content, err = ioutil.ReadFile(configFile); err != nil {
		log.Println("读取配置失败")
		return
	}

	// 反序列化
	if err = json.Unmarshal(content, &config); err != nil {
		return
	}

	// 赋值单例
	E_config = &config

	//log.Print(E_config)

	return

}