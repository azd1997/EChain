/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/8/1 20:26
* @Description: The file is for
***********************************************************************/

package cmd

import (
	"EChain/enode/ca"
	"EChain/enode/database"
	"EChain/enode/ends"
	"EChain/enode/log"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/spf13/cobra"
)

var endmgrCmd = &cobra.Command{
	Use:   "endmgr",
	Short: "终端设备管理器",
	Long:  `终端设备管理器，用来对终端设备接入等进行管理`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("终端设备管理器")

		// 打印用法
	},
}

var assignEndIdCmd = &cobra.Command{
	Use:   "assignid",
	Short: "分配Id",
	Long:  `对终端设备分配Id`,
	Run: func(cmd *cobra.Command, args []string) {

		_ = cmd.ParseFlags(args)
		assignEndId(
			cmd.Flag("name").Value.String(),
			cmd.Flag("usage").Value.String(),
			cmd.Flag("location").Value.String(),
		)
	},
}


func assignEndId(endName, endUsage, endWhere string) {
	//要做的事：
	//1.根据参数生成EndId
	//2.打印在终端
	//3.保存到BadgerDB的endids表中 <k,v> = <id,end>

	end := &ends.End{
		EndName:endName,
		EndUsage:endUsage,
		EndWhere:endWhere,
		EndWho:endWho,
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