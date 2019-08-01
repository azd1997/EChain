/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/21 17:42
* @Description: 软件版本命令
***********************************************************************/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().String("ver", "321", "版本")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ENode --blockchain-based iot platform v0.1 -- HEAD")
		fmt.Println(args)
	},
}
