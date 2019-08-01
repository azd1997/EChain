/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/21 17:11
* @Description: The file is for
***********************************************************************/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "enode",
	Short: "enode is a blockchain-based iot platform",
	Long: `A Fast and Flexible Static Site Generator built with
               love by spf13 and friends in Go.
               Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("测试Cobra...")
		fmt.Println(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


