/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/4 20:31
* @Description: The file is for
***********************************************************************/

package ends

import (
	"github.com/tarm/serial"
	"log"
	"testing"
)

//测试说明：使用VSPD创建一对虚拟串口，串口A由本段代码打开，串口B使用sscom等串口调试软件打开
func TestSerial(t *testing.T) {
	c := &serial.Config{Name:"COM1", Baud:115200}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("testwrite"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
	log.Println(buf)

}

func TestSerialConnector(t *testing.T) {

	connector, err := CreateSerialConnector("COM1", 115200)
	defer log.Println("成功关闭连接")
	defer connector.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = connector.WriteBytes([]byte("testwrite"))
	if err != nil {
		log.Fatal(err)
	}

	readData, err := connector.ReadBytes()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", readData)
	log.Println(readData)

}