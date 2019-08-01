/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/31 17:25
* @Description: The file is for
***********************************************************************/

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	// 验证命令行参数合法性
	validateArgs()

	// 定义两个命令
	cmd1 := flag.NewFlagSet("memver", flag.ExitOnError)
	nBits1 := cmd1.Int("n", 0, "输入比特位数")
	m1 := cmd1.Int("m", 1, "m = 0 或 1")  // 默认为1

	cmd2 := flag.NewFlagSet("memver", flag.ExitOnError)
	nBits2 := cmd2.Int("n", 0, "输入比特位数")
	m2 := cmd2.Int("m", 1, "m = 0 或 1")  // 默认为1

	switch os.Args[1] {
	case "memver":
		err := cmd1.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "datver":
		err := cmd2.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()

	}

	if cmd1.Parsed() {
		list := calcDataSet(*nBits1)
		fmt.Println("总数据集长度: ", len(list))
		fmt.Println("总数据集: ", list)

		list0, list2 := ClassifyDataSetByM(list, uint8(*m1))
		fmt.Println("list0 (计算系数为0): ")
		fmt.Println(list0)
		fmt.Println("list2 (计算系数为2): ")
		fmt.Println(list2)
	}

	if cmd2.Parsed() {
		fmt.Println("开始计算...")
		calcDataSetDatVer(*nBits2, uint8(*m2))
		fmt.Println("计算完成，数据集文件生成于当前目录下")
	}

	return
}

func validateArgs() {
	// m可以省略以使用默认值，但n不可以。因此要么是三个args要么是4个args
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("clac memver -n [n] -m [m]  —— 使用基于内存的版本，所得数据集暂存于内存中，n不宜过大，建议20以下，所得结果打印与命令行终端")
	fmt.Println("clac datver -n [n] -m [m]  —— 使用基于dat文件存储的版本，所得数据集存于dat文本中")
	fmt.Println("所得数据集每行最后一个元素为求得的汉明距离")
}


type DataSet [][]uint8

// n尽量小于20，不然内存不够用报错。  另一种解决办法是每一次得到序列结果就存入文本而不是驻留在内存
func calcDataSet(n int) (result DataSet) {

	max := int64(math.Pow(2, float64(n)))

	result = make([][]uint8, max)

	var i int64
	for i=0;i<max;i++ {

		// 得到全0向量
		var vec []uint8
		vec = make([]uint8, n+1)
		for j:=0;j<n+1;j++ {
			vec[j] = 0
		}

		// 得到数字的二进制字符串
		binaryStr := fmt.Sprintf("%b", i)
		// 等价于binaryStr := strconv.FormatUint(uint64(i), 2)
		runeB := []rune(binaryStr)
		lenB := len(runeB)

		// 转为无符号整型数组
		for j:=0;j<lenB;j++ {
			num, _ := strconv.ParseUint(string(runeB[j]), 0, n)
			vec[n-lenB+j] = uint8(num)
		}

		// 计算每个vec序列(前n位)中有多少个1
		var numOf1 uint8 = 0
		for i:=0;i<n;i++ {
			numOf1 += vec[i]
		}
		vec[n] = numOf1  // 汉明距离值

		// 将该序列结果存储
		result[i] = vec
	}


	return result
}

func ClassifyDataSetByM(data DataSet, m uint8) (data0,data1 DataSet) {
	// m = 0 or 1
	// data

	// 检查m
	if m != 0 && m != 1 {
		return nil, nil
	}

	// 内存分配
	data0, data1 = make([][]uint8, 0), make([][]uint8, 0)

	n := len(data[0]) - 1  // 还原序列长度

	// 循环遍历
	for _, a := range data {
		numOf1 := a[n]
		// class = 0 or 2
		class := uint8(1 + math.Pow(-1, float64(m + numOf1)))
		switch class {
		case 0:
			data0 = append(data0, a)
		case 2:
			data1 = append(data1, a)
		default:
			log.Panic("异常！")
		}
	}

	return data0, data1
}

// 只存储list0及list2.
func calcDataSetDatVer(n int, m uint8) {

	// 打开或创建数据文件
	file0, err := os.OpenFile("list0.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}
	defer file0.Close()
	// 向其写入开头内容
	str := fmt.Sprintf("\nList0 (n=%d,m=%d) :\n", n, uint8(m))
	if _, err := file0.WriteString(str); err != nil {
		log.Panic(err)
	}

	file2, err := os.OpenFile("list2.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}
	defer file2.Close()
	// 向其写入开头内容
	str = fmt.Sprintf("\nList2 (n=%d,m=%d) :\n", n, uint8(m))
	if _, err := file2.WriteString(str); err != nil {
		log.Panic(err)
	}

	max := int64(math.Pow(2, float64(n)))

	var i int64
	for i=0;i<max;i++ {

		// 得到全0向量
		var vec []uint8
		vec = make([]uint8, n+1)
		for j:=0;j<n+1;j++ {
			vec[j] = 0
		}

		// 得到数字的二进制字符串
		binaryStr := fmt.Sprintf("%b", i)
		// 等价于binaryStr := strconv.FormatUint(uint64(i), 2)
		runeB := []rune(binaryStr)
		lenB := len(runeB)

		// 转为无符号整型数组
		for j:=0;j<lenB;j++ {
			num, _ := strconv.ParseUint(string(runeB[j]), 0, n)
			vec[n-lenB+j] = uint8(num)
		}

		// 计算每个vec序列(前n位)中有多少个1
		var numOf1 uint8 = 0
		for i:=0;i<n;i++ {
			numOf1 += vec[i]
		}
		vec[n] = numOf1  // 汉明距离值

		// 注意！[]uint8。 []byte其实就是[]uint8，所以可以直接用于file.Write([]byte)。但注意写入文本时byte是uint8代表的字符，而我们的uint8其实代表数字大小，因此需要转为字符
		//var vecBytes = make([]string, n+1)
		//for i:=0;i<=n;i++ {
		//	vecBytes[i] = strconv.FormatUint(uint64(vec[i]), 10)
		//}
		vecStr := fmt.Sprintf("%d", vec)

		// 将该序列结果存储入文件
		class := uint8(1 + math.Pow(-1, float64(m + numOf1)))
		switch class {
		case 0:
			file0.WriteString(vecStr)
			file0.WriteString("\n")
		case 2:
			file2.WriteString(vecStr)
			file2.WriteString("\n")
		default:
			// 出问题就啥也不干，等于跳过
		}
	}
}





func getList(n int) (list [][]int) {

	x := math.Pow(2, float64(n))
	fmt.Println(x)
	counter := 0


	for float64(counter) < x {
		// 每次循环时随机置位
		var vector []int
		for i:=0;i<n;i++ {
			vector = append(vector, rand.Intn(2))
			fmt.Println(vector)
		}
		//fmt.Println("未检查", vector)

		// 检查vector是否重复，不重复添加到list
		for _, a := range list {
			fmt.Println("a: ", a)
			for i:=0;i<n;i++ {
				if a[i] != vector[i] {
					//说明a!=vector
					list[counter] = vector
					counter++
					//fmt.Println(vector)
				}
			}
		}



		for j:=0;j<=counter;j++ {
			for i:=0;i<n;i++ {
				if list[j][i] != vector[i] {
					//说明a!=vector
					list[counter] = vector
					counter++
					//fmt.Println(vector)
				}
			}
		}

	}

	return list
}


func getList2(n int) (list [][]int) {

	max := int(math.Pow(2, float64(n)))
	fmt.Println(max)

	for i:=0;i<max;i++ {
		//binaryStr := biu.ByteToBinaryString(byte(i))
		//fmt.Println(byte(i))
		//fmt.Println(binaryStr)

		b := fmt.Sprintf("%b", i)
		fmt.Println("binaryStr: ", b)
		//b1 := strconv.FormatUint(uint64(i), 2)
		//fmt.Println("b1: ", b1)
		//n, _ := strconv.ParseUint(b, 0, n)
		//fmt.Println("n: ", n)

		runeB := []rune(b)

		var num []uint64
		num = make([]uint64, n)
		for i:=0;i<n;i++ {
			num[i] = 0
		}

		len1 := len(runeB)
		for j:=0;j<len1;j++ {
			nu, _ := strconv.ParseUint(string(runeB[j]), 0, n)
			fmt.Println(nu)
			//num = append(num, nu)
			num[n-len1+j] = nu

		}
		fmt.Println("num: ", num)

	}
	return nil

}
