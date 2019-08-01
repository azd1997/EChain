/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 14:53
* @Description: The file is for
***********************************************************************/

package account

import (
	"EChain/enode/ledger"
	"EChain/enode/utils"
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// TODO: accounts.go负责处理账户的文件存储编码解码、加密解密、哈西验证等等


//const accountFile = "../tmp/accounts/accounts_%s.data" //根据不同node ip生成不同的accounts.data

// 以表的形式存储账户信息
type Accounts struct {
	Map map[string]*Account
}

// 将Accounts字典维护的内容编码之后写进文本
//		注意每次保存都是使用新的Accounts对象取刷新原先的文本内容
func (as *Accounts) SaveFile(nodeAddr string) {
	var content bytes.Buffer

	// 检查路径存在与否，不存在则创建
	pathTemp := ledger.Config().AccountsFilePathTemp
	log.Println(pathTemp)
	dir := strings.TrimSuffix(pathTemp, "/accounts_%s.data")
	log.Println(dir)
	if exists, _ := utils.DirExists(dir); !exists {
		if err := os.MkdirAll(dir, os.ModeDir); err != nil {
			log.Fatal(err)
		}
	}

	accountFile := fmt.Sprintf(pathTemp, utils.IpToDir(nodeAddr))

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(as)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(accountFile, content.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// 从文本文件加载钱包文件，解码后还原出钱包字典
func (as *Accounts) LoadFile(nodeAddr string) error {

	accountFile := fmt.Sprintf(ledger.Config().AccountsFilePathTemp, utils.IpToDir(nodeAddr))
	if _, err := os.Stat(accountFile); os.IsNotExist(err) {
		return err
	}

	var accounts Accounts

	fileContent, err := ioutil.ReadFile(accountFile)
	if err != nil {
		log.Fatal(err)
	}

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&accounts)
	if err != nil {
		log.Fatal(err)
	}

	as.Map = accounts.Map

	return nil
}

// 从账户文件读取账户（可能有多个）信息，组装Accounts对象
func NewAccountsFromFile(nodeAddr string) (*Accounts, error) {
	accounts := Accounts{}
	accounts.Map = make(map[string]*Account)

	err := accounts.LoadFile(nodeAddr)

	return &accounts, err
}

// 将账户地址作为键，从账户字典中查找对应账户
func (as *Accounts) GetAccount(address string) Account {
	return *as.Map[address]
}

// 从账户字典获取所有账户地址，并存入账户地址的切片数组中
func (as *Accounts) GetAllAddress() []string {
	var addresses []string

	for address := range as.Map {
		addresses = append(addresses, address)
	}

	return addresses
}

// 生成新账户并加入账户字典，返回账户地址
func (as *Accounts) AddAccount() string {
	account := NewAccount()
	address := fmt.Sprintf("%s", account.Address())

	as.Map[address] = account

	return address
}