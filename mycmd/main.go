/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/22 13:07
* @Description: The file is for
***********************************************************************/

package main

import (
	"EChain/mycmd/cmd"
	"fmt"
)



type Commander interface {
	Execute() error
}

func main() {
	//cmd.Execute()
	var RootCmd = &cmd.Command{
		Use:   "test",
		Short: "A brief description of your application",
		Long:  "A longer description",
	}
	RootCmd.SetOptions = func(c *cmd.Command) error {
		fmt.Println("Set Options here")
		c.Flags().StringP("user", "u", "Anonymous", "Help message for username")
		return nil
	}
	RootCmd.Parse = func(c *cmd.Command) error {
		fmt.Println("Parse here")
		c.Flags().Parse(*cmd.GetArgs())
		return nil
	}
	RootCmd.Run = func(c *cmd.Command, a []string) {
		fmt.Println("Do comamnd")
		username, _ := c.Flags().GetString("user")
		fmt.Println("myCommand called by " + username)
	}
	RootCmd.Execute()
}

