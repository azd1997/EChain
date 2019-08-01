/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/18 22:05
* @Description: The file is for
***********************************************************************/

package utils

// 从前向后遍历，提取两个字符串的相同前缀
func GetSamePrefixFromTwoStrings(s1, s2 string) string {

	// 记住：直接按下标遍历string其实是按字节（uint8）遍历。这里先转为[]rune
	s1Runes, s2Runes := []rune(s1), []rune(s2)

	// 取两字符串短字符串的字符数，作为遍历终止
	l1, l2 := len(s1Runes), len(s2Runes)
	l := l1
	if l1 > l2 {
		l = l2
	}

	// 从前向后遍历，直至遇不相等字符则退出
	var samePrefix []rune
	for i:=0; i<l; i++ {
		if s1Runes[i] == s2Runes[i] {
			samePrefix = append(samePrefix, s1Runes[i])
		} else {
			break
		}
	}

	return string(samePrefix)
}
