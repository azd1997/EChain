/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/22 16:35
* @Description: The file is for
***********************************************************************/

package cmd

import (
	"fmt"
	"os"
	"sync"
)

var args *[]string
var mu sync.Mutex

// GetArgs *
func GetArgs() *[]string {
	if args == nil {
		mu.Lock()
		defer mu.Unlock()
		if args == nil {
			args = &os.Args
			fmt.Println("init logic here...", *args)
		}
	}
	return args
}

func init() {
	fmt.Println("use cmd.GetArgs() anywhere...", GetArgs())
}
