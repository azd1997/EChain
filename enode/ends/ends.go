/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/4 18:53
* @Description: 处理终端设备连接，收集终端设备数据
***********************************************************************/

package ends

type EndsConnector interface {
	//Connector()
	ReadBytes() (result []byte, err error)
	WriteBytes(data []byte) (err error)
	Close() (err error)
}

type End struct {
	EndId []byte  // 由CA生成，先用sha256顶着
	EndName string // "DHT11_1"
	EndUsage string //"temperature"
	EndWho string //接入的工作人员姓名工号 "eiger"
	EndWhen int64 //最开始分发Id的时间戳
	EndWhere string  //终端设备部署位置
	EndWhichNode string //由哪个网关节点签发的Id "127.0.0.1:9717"
}