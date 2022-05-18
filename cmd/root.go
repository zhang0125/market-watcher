/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zhang0125/market-watcher/pkg/log"
	"github.com/zhang0125/market-watcher/pkg/types"
	"github.com/zhang0125/market-watcher/pkg/util"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "binance",
	Short: "binance data fetcher",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// initialize tendermint viper config
		util.InitViperConfig(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String(types.HomeFlag, os.ExpandEnv("$HOME/.binance"), "directory for config and data")
	rootCmd.PersistentFlags().String(types.ConfFlag, os.ExpandEnv("default.toml"), "config file name")

	// bind all flags with viper
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		log.Log.Error("init | BindPFlag | rootCmd.Flags", "Error", err)
	}
}
