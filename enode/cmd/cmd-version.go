/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/21 17:42
* @Description: 软件版本命令
***********************************************************************/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出当前软件版本信息",
	Long:  `输出当前软件版本信息，包括软件名称，版本及可执行文件存放位置`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func version() {
	absPath, _ := filepath.Abs("../entry/enode.exe")
	fmt.Println("enode v0.1.0 ", absPath)
}
