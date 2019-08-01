module EChain

require (
	github.com/AndreasBriese/bbloom v0.0.0-20190306092124-e2d15f34fcf9 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/dgraph-io/badger v1.5.4
	github.com/dgryski/go-farm v0.0.0-20190423205320-6a90982ecee2 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/mr-tron/base58 v1.1.2
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.2
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/vrecan/death v3.0.1+incompatible
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.2
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace (
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894 => github.com/golang/sys v0.0.0-20190422165155-953cdadca894
)

go 1.12
