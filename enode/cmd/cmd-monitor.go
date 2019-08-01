/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/8/1 20:27
* @Description: The file is for
***********************************************************************/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "监控服务器",
	Long:  `监控服务器，监测网关节点及终端设备状态`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("监控系统启动")
		// 启动服务器
	},
}