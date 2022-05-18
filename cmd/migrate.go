package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zhang0125/market-watcher/model/binancemodel"
	"github.com/zhang0125/market-watcher/pkg/util"
)

// GetMigrateCmd returns the start command to start bridge
func GetMigrateCmd() *cobra.Command {
	//var logger = log.Log.With("cmd", "binance")
	binanceCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate binance db",
		Run: func(cmd *cobra.Command, args []string) {
			util.InitBinanceConfig()
			util.InitDB()
			err := util.DB.AutoMigrate(&binancemodel.WsKlineModel{})
			if err != nil {
				return
			}
		}}
	return binanceCmd
}

func init() {
	rootCmd.AddCommand(GetMigrateCmd())
}
