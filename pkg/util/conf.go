package util

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zhang0125/market-watcher/pkg/types"
)

type Serve struct {
	Interval string
	Enable   bool
}

type Configuration struct {
	Symbols []string
	Dsn     string
	Serve   map[string]Serve
}

var conf Configuration

func GetConfig() *Configuration {
	return &conf
}

func InitBinanceConfig() {
	fmt.Println("start init binance config")
	_, err := toml.DecodeFile(
		viper.GetString(types.HomeFlag)+"/conf/"+viper.GetString(types.ConfFlag),
		&conf,
	)
	if err != nil {
		fmt.Println("init binance fail ", err)
	}
	fmt.Println("end init binance config")
}

// InitViperConfig sets global viper configuration
func InitViperConfig(cmd *cobra.Command) {
	homeValue, _ := cmd.Flags().GetString(types.HomeFlag)
	viper.Set(types.HomeFlag, homeValue)
	confValue, _ := cmd.Flags().GetString(types.ConfFlag)
	viper.Set(types.ConfFlag, confValue)
}
