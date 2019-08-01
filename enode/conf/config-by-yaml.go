/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 21:57
* @Description: The file is for
***********************************************************************/

package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 读取配置
func InitConfigByYaml(configFile string) (err error) {
	var (
		content []byte
		config Config
	)

	// 读取配置文件，得到[]byte内容
	if content, err = ioutil.ReadFile(configFile); err != nil {
		log.Println("读取配置失败")
		return
	}

	// 反序列化到config结构体
	if err = yaml.Unmarshal(content, &config); err != nil {
		return err
	}

	// 赋值单例
	E_config = &config

	return
}
