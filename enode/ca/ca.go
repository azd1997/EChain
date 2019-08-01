/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/5 14:50
* @Description: The file is for
***********************************************************************/

package ca

import (
	"EChain/enode/ends"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	log "github.com/sirupsen/logrus"
	"time"
)

type Ca struct {

}

//CA生成EndsId，先简单的根据时间及网关节点信息以及
//流程：命令行输入$ enode -ca -endname DHT11_1 -endusage temperature -endwho eiger,
// param parsed -> createEndId -> 组装End结构体（只含这三项数据及本机节点ip） -> ca.CreateEndId(end)
func (ca *Ca) CreateEndId(end ends.End) (endId []byte) {

	end.EndWhen = time.Now().UnixNano()  //当前时间的纳秒时间戳

	//先序列化
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(end)
	if err != nil {
		log.Error(err)
	}

	hash := sha256.Sum256(result.Bytes())
	endId = hash[:]
	end.EndId = endId

	return
}