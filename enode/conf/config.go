/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/4 16:00
* @Description: The file is for
***********************************************************************/

package conf

var (
	// 单例
	E_config *Config
)

type Config struct {
	Logger loggerConfig	`json:"loggerConfig" yaml:"loggerConfig" toml:"loggerConfig"`
	Ledger ledgerConfig	`json:"ledgerConfig" yaml:"ledgerConfig" toml:"ledgerConfig"`
	End endConfig	`json:"endConfig" yaml:"endConfig" toml:"endConfig"`
}

// 为什么这些类型不导出？因为不希望被外部创建这些结构体，外部只需要能够访问到我配置文件的数据就行了

// TODO:整个工程在后续优化错误处理！

// TODO:配置文件热重载
// https://www.cnblogs.com/CraryPrimitiveMan/p/7928647.html




