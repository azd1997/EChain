/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/19 22:56
* @Description: The file is for
***********************************************************************/

package conf

type loggerConfig struct {
	MongodbUri string `json:"mongodbUri" yaml:"mongodbUri" toml:"mongodbUri"`
	MongodbConnectTimeout int `json:"mongodbConnectTimeout" yaml:"mongodbConnectTimeout" toml:"mongodbConnectTimeout"`
	LogBatchSize int `json:"logBatchSize" yaml:"logBatchSize" toml:"logBatchSize"`
	LogCommitTimeout int `json:"logCommitTimeout" yaml:"logCommitTimeout" toml:"logCommitTimeout"`
}
