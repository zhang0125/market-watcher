package binance

import (
	"context"
	"fmt"
	"sync"

	"github.com/adshao/go-binance/v2"
	"github.com/zhang0125/market-watcher/model/binancemodel"
	"github.com/zhang0125/market-watcher/pkg/log"
	"github.com/zhang0125/market-watcher/pkg/util"
	"github.com/zhang0125/market-watcher/service/base"
)

var _ base.Service = (*Service)(nil)

type Service struct {
	base.BaseService
	stopped bool
}

func NewBinanceService() *Service {
	ts := &Service{}
	ts.BaseService = *base.NewBaseService(log.Log, "BinanceService", ts)
	return ts
}
func (t *Service) OnStop() {
	t.stopped = true
}

func (t *Service) OnStart(ctx context.Context) error {
	go t.StartFetch(ctx)
	return nil
}

func (t *Service) StartFetch(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	go klineFetch(&wg)
	wg.Wait()
}

var errHandler = func(err error) {
	fmt.Println(err)
}

func klineFetch(wg *sync.WaitGroup) {
	defer wg.Done()
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		if !event.Kline.IsFinal {
			// Waiting for data completion
			return
		}
		if err := binancemodel.CreateKline(event); err != nil {
			log.Log.Error("save kline failed", err)
		}
		log.Log.Info("event", event)
	}
	klineServe, ok := util.GetConfig().Serve["kline"]
	if !ok || !klineServe.Enable {
		return
	}
	symbols := util.GetConfig().Symbols
	interval := klineServe.Interval
	symbolIntervalPair := make(map[string]string, len(symbols))
	for _, symbol := range symbols {
		symbolIntervalPair[symbol] = interval
	}
	doneC, _, err := binance.WsCombinedKlineServe(symbolIntervalPair, wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
