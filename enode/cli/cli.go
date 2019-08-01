/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/5 15:35
* @Description: The file is for
***********************************************************************/

package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	log "github.com/sirupsen/logrus"
)

type CommandLine struct {

}

/*打印命令行工具用法*/
func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" ca -endname NAME -endusage USAGE -endwho WHO -输入终端设备信息生成ID")
}

/*检查命令行输入参数是否至少有两个*/
func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
		//os.Exit(1)
	}
}

/*运行命令行程序*/
func (cli *CommandLine) Run() {
	cli.validateArgs()

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Printf("请设置NODE_ID环境变量!")
		runtime.Goexit()
	}

	//设置命令行参数
	caCmd := flag.NewFlagSet("ca", flag.ExitOnError)
	//createChainCmd

	caEndName := caCmd.String("endname", "DHT11_test", "终端设备名称编号")
	caEndUsage := caCmd.String("endusage", "temperature", "终端设备用途")
	caEndWhere := caCmd.String("endwhere", "三号楼521", "终端设备部署位置")
	caEndWho := caCmd.String("endwho", "eiger", "终端设备ID分配操作人")

	//解析命令行参数
	switch os.Args[1] {
	case "ca":
		err := caCmd.Parse(os.Args[2:])
		if err != nil {
			log.Error(err)
		}
	}

	//根据参数处理任务
	if caCmd.Parsed() {
		//由于我都设置了默认值，所以这里不做参数的空检查，直接执行任务
		cli.assignEndId(*caEndName, *caEndUsage, *caEndWhere, *caEndWho, nodeID)
	}

}

