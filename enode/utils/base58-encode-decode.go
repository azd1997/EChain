/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 15:42
* @Description: The file is for
***********************************************************************/

package utils

import (
	"github.com/mr-tron/base58"
	"log"
)

/*基于第三方base58库实现的Base58编码*/
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

/*基于第三方base58库实现的Base58解码*/
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Fatal(err)
	}

	return decode
}