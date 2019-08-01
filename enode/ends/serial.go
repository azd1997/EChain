/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/4 19:49
* @Description: 与终端设备串口通信
***********************************************************************/

package ends

import (
	"github.com/tarm/serial"
)

//go serial: https://github.com/tarm/serial

//串口通信的工作逻辑：
//1.开辟协程，循环检查串口有无新增（有无各种协议连接）
//2.发现串口新增，将新增串口等相关信息通过通道传给新协程
//3.新协程中配置串口连接，读取数据
//4.将数据存入数据库（）
//5.生成身份证件信息，构造评分，今后检查

//serial.go对tarm/serial做一些封装

//每一个串口一个SerialConnector
type SerialConnector struct {
	//SerialConfig *serial.Config
	serialPort *serial.Port  //不需要导出
}


// ex. comName : COM1
func CreateSerialConnector(comName string, baudrate int) (connector *SerialConnector, err error) {
	c := &serial.Config{Name:comName, Baud:baudrate}
	s, err := serial.OpenPort(c)
	if err != nil {
		return
	}
	connector = &SerialConnector{
		serialPort:s,
	}
	return
}

func (connector *SerialConnector) ReadBytes() (result []byte, err error) {

	result = make([]byte, 128)
	n, err := connector.serialPort.Read(result)
	if err != nil {
		return
	}
	result = result[:n]
	return
}

func (connector *SerialConnector) WriteBytes(data []byte) (err error) {

	_, err = connector.serialPort.Write(data)
	if err != nil {
		return
	}
	return
}

func (connector *SerialConnector) Close() (err error) {
	err = connector.serialPort.Close()
	return
}
