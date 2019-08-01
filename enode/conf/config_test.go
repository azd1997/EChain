/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/17 17:39
* @Description: The file is for
***********************************************************************/

package conf

import (
	"log"
	"testing"
)

func TestInitConfigByJson(t *testing.T) {
	var configFile = "enode.json"
	if err := InitConfigByJson(configFile); err != nil {
		log.Fatal(err)
	}
	log.Println(E_config)
}

func TestInitConfigByYaml(t *testing.T) {
	var configFile = "enode.yaml"
	if err := InitConfigByYaml(configFile); err != nil {
		log.Fatal(err)
	}
	log.Println(E_config)
}

func TestInitConfigByToml(t *testing.T) {
	var configFile = "enode.toml"
	if err := InitConfigByToml(configFile); err != nil {
		log.Fatal(err)
	}
	log.Println(E_config)
}
