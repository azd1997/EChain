/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/25 21:42
* @Description: The file is for
***********************************************************************/

package logger

import "github.com/spf13/viper"

// import "EChain/enode/conf"

type conf struct {
	mongodbUri string
	mongodbConnectTimeout int
	logBatchSize int
	logCommitTimeout int
}

// 配置文件内容
//# mongoDB地址: 采用mongoDB uri
//mongodbUri: mongodb://127.0.0.1:27017
//# mongoDB连接超时: 单位毫秒
//mongodbConnectTimeout: 5000
//# "日志批次大小": "为了减少mongodb网络往返耗时，打包成一批写入",
//logBatchSize: 100
//# "日志自动提交超时": "在batch未达到批次阈值之前，超时会自动提交日志，单位ms",
//logCommitTimeout: 5000


// 使用自建conf包来读配置
//func Config() *Conf {
//	return &Conf{
//		MongodbUri:conf.E_config.Logger.MongodbUri,
//		MongodbConnectTimeout:conf.E_config.Logger.MongodbConnectTimeout,
//		LogBatchSize:conf.E_config.Logger.LogBatchSize,
//		LogCommitTimeout:conf.E_config.Logger.LogCommitTimeout,
//	}
//}

// 使用Viper配置
func config() *conf {
	return &conf{
		mongodbUri:viper.GetString("loggerConfig.mongodbUri"),
		mongodbConnectTimeout:viper.GetInt("loggerConfig.mongodbConnectTimeout"),
		logBatchSize:viper.GetInt("loggerConfig.logBatchSize"),
		logCommitTimeout:viper.GetInt("loggerConfig.logCommitTimeout"),
	}
}

