/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/4 17:14
* @Description: The file is for
***********************************************************************/

package logger

//日志批次（多条日志，日志缓存）
//logBatch只在本包使用
type logBatch struct {
	logs []interface{}  //多条日志
}

