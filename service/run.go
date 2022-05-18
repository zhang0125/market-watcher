package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/zhang0125/market-watcher/pkg/util"
	"github.com/zhang0125/market-watcher/service/base"
	"github.com/zhang0125/market-watcher/service/binance"
)

var serviceList []base.Service

func NewServiceList() {
	serviceList = append(serviceList, binance.NewBinanceService())
}

func Run(cmd *cobra.Command, args []string) {
	serviceList = make([]base.Service, 0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	NewServiceList()
	// sync group
	var wg sync.WaitGroup
	wg.Add(len(serviceList))
	for _, service := range serviceList {
		go func(service base.Service) {
			defer wg.Done()
			err := service.Start(ctx)
			if err != nil {
				panic("service start error")
			}
		}(service)

	}

	// wait for all processes
	wg.Wait()

	// catch signal
	catchSignal := make(chan os.Signal, 1)
	signal.Notify(catchSignal, os.Interrupt, syscall.SIGTERM)

	// sig is a ^C, handle it
	for range catchSignal {
		// stop processes
		for _, service := range serviceList {
			fmt.Println("service stop")
			service.Stop()
		}

		fmt.Println("waiting goroutine to stop...")
		time.Sleep(time.Second * 5)

		//stop DB
		util.CloseDB()
		return
	}
}
