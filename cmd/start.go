package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zhang0125/market-watcher/pkg/util"
	"github.com/zhang0125/market-watcher/service"
)

// GetStartCmd returns the start command to start bridge
func GetStartCmd() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start watcher server",
		Run: func(cmd *cobra.Command, args []string) {
			util.InitBinanceConfig()
			util.InitDB()
			service.Run(cmd, args)
		}}
	return startCmd
}

func init() {
	rootCmd.AddCommand(GetStartCmd())
}
