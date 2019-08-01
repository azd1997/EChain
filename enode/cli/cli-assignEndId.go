/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/5 15:57
* @Description: The file is for
***********************************************************************/

package cli

import (
	"EChain/enode/ca"
	"EChain/enode/database"
	"EChain/enode/ends"
	"bytes"
	"encoding/gob"
	log "github.com/sirupsen/logrus"
)

func (cli *CommandLine) assignEndId(endName, endUsage, endWhere, endWho, nodeId string) {
	//要做的事：
	//1.根据参数生成EndId
	//2.打印在终端
	//3.保存到BadgerDB的endids表中 <k,v> = <id,end>

	end := &ends.End{
		EndName:endName,
		EndUsage:endUsage,
		EndWho:endWho,
		EndWhere:endWhere,
		EndWhichNode:nodeId,
	}

	//TODO:优化、决定到底由哪个包实现

	localCa := &ca.Ca{}
	id := localCa.CreateEndId(*end)  //之后end被填上time和Id

	//存储End及Id
	var endBuffer bytes.Buffer
	encoder := gob.NewEncoder(&endBuffer)
	err := encoder.Encode(end)
	if err != nil {
		log.Error(err)
	}
	endBytes := endBuffer.Bytes()

	badger, err := database.OpenDatabase("badger", "tmp/badgerdb/endids")
	if err != nil {
		log.Error(err)
	}
	err = badger.Set(id, endBytes)
	if err != nil {
		log.Error(err)
	}

	//打印
	log.Info("生成Id成功！", id)

}