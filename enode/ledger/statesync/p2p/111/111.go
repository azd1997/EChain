/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 17:25
* @Description: The file is for
***********************************************************************/

package _11

import "os"

func main() {

	defer os.Exit(0)
	Cli := cli.CommandLine{}
	Cli.Run()

}
