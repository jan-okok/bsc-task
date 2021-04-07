package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

var (
	Config *tomlConfig
	once   sync.Once
)

func init() {
	once.Do(func() {
		filePath, err := filepath.Abs("./conf/config.toml")
		if err != nil {
			panic(err)
		}
		if _, err := toml.DecodeFile(filePath, &Config); err != nil {
			panic(fmt.Sprintf("Read conf file error,Err=[%v]", err))
		}
	})
}

type tomlConfig struct {
	BscSignUrl      string `toml:"bscSignUrl"`
	BscSignUser     string `toml:"bscSignUser"`
	BscSignPassword string `toml:"bscSignPassword"`
	Transaction     struct {
		FromAddress     string `toml:"fromAddress"`
		Amount          string `toml:"amount"`
		Fee             int64  `toml:"fee"`
		ToAddrFileName  string `toml:"toAddrFileName"`
		ContractAddress string `toml:"contractAddress"`
	} `toml:"transaction"`
}
