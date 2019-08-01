/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/26 18:37
* @Description: The file is for
***********************************************************************/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var loggerCmd = &cobra.Command{
	Use:   "logger",
	Short: "日志系统",
	Long:  `日志系统用来记录程序运行日志，存入mongodb`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("日志系统启动")
	},
}