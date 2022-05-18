package binancemodel

import (
	"github.com/adshao/go-binance/v2"
	"github.com/zhang0125/market-watcher/pkg/util"
)

type WsKlineModel struct {
	ID     uint            `gorm:"primaryKey"`
	Event  string          `json:"e"`
	Time   int64           `json:"E"`
	Symbol string          `json:"s"`
	Kline  binance.WsKline `gorm:"embedded" json:"k"`
}

func (WsKlineModel) TableName() string {
	return "kline"
}

func CreateKline(e *binance.WsKlineEvent) error {
	tx := util.DB.Create(&WsKlineModel{
		Event:  e.Event,
		Time:   e.Time,
		Symbol: e.Symbol,
		Kline:  e.Kline,
	})
	return tx.Error
}
