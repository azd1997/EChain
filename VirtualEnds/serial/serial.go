/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/4 23:09
* @Description: The file is for
***********************************************************************/

package main

import (
	"github.com/tarm/serial"
	"log"
)

func main() {
	//选定的COM须与Ends/serial.go选定的串口配对
	c := &serial.Config{Name:"COM2", Baud:115200}

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


	return
}