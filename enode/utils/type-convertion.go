package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Int64转为[]byte （大端）
func Int64ToBytes(i int64) []byte {
	var myBytes = make([]byte, 8)
	binary.BigEndian.PutUint64(myBytes, uint64(i))
	return myBytes
}

// []byte转为Int64  （大端）
func BytesToInt64(bytes []byte) int64 {
	return int64(binary.BigEndian.Uint64(bytes))
}


// 整形转换成字节
func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian,tmp)
	return bytesBuffer.Bytes()
}

// 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	_ = binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

// 整型数据转换成十六进制的字节数组
func Int64ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Fatal(err)
	}
	return buff.Bytes()
}