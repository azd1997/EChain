/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/22 16:35
* @Description: The file is for
***********************************************************************/

package cmd

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

// Command .
type Command struct {
	// Use is the one-line usage message.
	Use string
	// Short is the short description shown in the 'help' output.
	Short string
	// Long is the long message shown in the 'help <this-command>' output.
	Long string
	// SetOptions:
	SetOptions func(c *Command) error
	// Parse:
	Parse func(c *Command) error
	// Run: Typically the actual work function. Most commands will only implement this.
	Run func(cmd *Command, args []string)
}

// Execute .
func (c *Command) Execute() error {
	if ok := c.SetOptions(c); ok != nil {
		fmt.Println("Error in SetOptions!")
		return ok
	}
	if ok := c.Parse(c); ok != nil {
		fmt.Println("Error in Parsing!")
		return ok
	}
	c.Run(c, *GetArgs())
	return nil
}

// Flags returns the complete FlagSet that applies
// to this command (local and persistent declared here and by all parents).
func (c *Command) Flags() *flag.FlagSet {
	if c.flags == nil {
		c.flags = flag.NewFlagSet(c.Use, flag.ContinueOnError)
	}
	return c.flags
}

