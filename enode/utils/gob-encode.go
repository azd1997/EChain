/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/8/1 21:19
* @Description: The file is for
***********************************************************************/

package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

/*将数据进行编码得到字节数组*/
func GobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}