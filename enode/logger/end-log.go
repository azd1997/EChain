/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/4 17:05
* @Description: The file is for
***********************************************************************/

package logger


//Endlog需要被其他包使用，所以大写
type EndLog struct {
	endId []byte
	endUsage string
	inTime int64  //接入时间
	inNums int64  //接入第几次
	inSuccess bool  //接入成功？
}
