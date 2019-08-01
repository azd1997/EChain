/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/18 19:29
* @Description: The file is for
***********************************************************************/

package utils

import (
	"testing"
)

func TestIpToDir(t *testing.T) {
	tests := []struct {
		ip string
		dir string
	}{
		{"127.0.0.1:8080", "127_0_0_1_8080"},
		{"255.255.255.255:9999", "255_255_255_255_9999"},
	}

	for _, tt := range tests {
		temp := IpToDir(tt.ip)
		if temp != tt.dir {
			t.Errorf("转换 %s 为目录失败：应该为 %s， 转换为 %s ", tt.ip, tt.dir, temp)
		}

	}
}

func TestGetSamePrefixFromTwoStrings(t *testing.T) {
	tests := []struct {
		s1, s2, samePrefix string
	}{
		{"", "", ""},
		{"你好啊", "你好啊", "你好啊"},
		{"你好啊，世界", "你好啊，世界真美", "你好啊，世界"},
		{"你好啊，世界", "你好啊，艾老师", "你好啊，"},
		{"你好啊，世界", "世界，你好啊", ""},
		{"你好啊，世界", "不知道什么", ""},
		{"/azd/1997/xxx/yyy", "/azd/1997/yyy/zzz", "/azd/1997/"},
	}

	for _, tt := range tests {
		same := GetSamePrefixFromTwoStrings(tt.s1, tt.s2)
		if same != tt.samePrefix {
			t.Errorf("提取字符串(%s,%s)相同前缀失败\n本应是%s，实际得到%s", tt.s1, tt.s2, tt.samePrefix, same)
		}
	}
}