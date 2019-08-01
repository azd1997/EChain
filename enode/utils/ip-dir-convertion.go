/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/18 19:23
* @Description: The file is for
***********************************************************************/

package utils

import "strings"

// 将ip字符串转为用于目录的字符串
// 		e.g. 127.0.0.1:8080 => 127_0_0_1_8080
func IpToDir(ip string) (dir string) {
	temp := strings.ReplaceAll(ip, ".", "_")
	return strings.ReplaceAll(temp, ":", "_")
}
