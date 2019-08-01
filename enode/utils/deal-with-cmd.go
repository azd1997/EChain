/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/8/1 21:20
* @Description: The file is for
***********************************************************************/

package utils

import "fmt"

/*从request信息中抽取前12字节作为命令*/
func ExtractCmd(request []byte, commandLength int) []byte {
	return request[:commandLength]
}

// 将命令（字符串）转为字节数组
func CmdToBytes(cmd string, commandLength int) []byte {

	var cmdBytes = make([]byte, commandLength)

	for i, c := range cmd {
		cmdBytes[i] = byte(c)
	}

	return cmdBytes[:]
}

/*将命令（字节数组）转为字符串并输出*/
func BytesToCmd(cmdBytes []byte) string {
	var cmd []byte

	for _, b := range cmdBytes {
		if b != 0x0 {
			cmd = append(cmd, b)
		}
	}

	return fmt.Sprintf("%s", cmd)
}