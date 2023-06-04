package service

import (
	"github.com/ImOsMa/bybit_service"
	"github.com/ImOsMa/bybit_service/pkg/client/bybit"
)

type Account interface {
	AccountInfo(user bybit_service.User) (bybit_service.AccountInfo, error)
	SpotWalletBalance(user bybit_service.User) ([]bybit_service.AccountWalletBalance, error)
	FeeRate(user bybit_service.User) (string, error)
	KeyInformation(user bybit_service.User) (bybit_service.KeyInformation, error)
	CoinExchangeRecords(user bybit_service.User)
}

type Market interface {
	GetKline(user bybit_service.User, symbol bybit.SymbolV5, interval bybit.Interval, limit int) (bybit_service.GetKline, error)
	InstrumentInfo(user bybit_service.User, symbol bybit.SymbolV5) (bybit_service.GetInstrumentInfo, error)
	Tickers(user bybit_service.User, symbol bybit.SymbolV5) (bybit_service.TickersSportResult, error)
	PositionInfo(user bybit_service.User, symbol bybit.SymbolV5) (bybit_service.PositionInfo, error)
	CoinInfo(user bybit_service.User, symbol bybit.SymbolV5) (bybit_service.CoinInfo, error)
}

type Order interface {
	PostSpotOrder(user bybit_service.User, request bybit_service.PostSpotOrderRequest) (bybit_service.PostSpotOrderResult, error)
	GetSpotOrder(user bybit_service.User, orderID string) (bybit_service.SpotGetOrderResult, error)
	DeleteSpotOrder(user bybit_service.User, orderID string) (bybit_service.SpotDeleteResult, error)
	GetOpenSpotOrders(user bybit_service.User, limit int) ([]bybit_service.SpotOpenOrdersResult, error)
	ChangeOrder(user bybit_service.User, params bybit_service.ChangeOrderParams) (string, error)
	History(user bybit_service.User)
}

type Service struct {
	Account
	Order
	Market
}

func NewService() *Service {
	return &Service{
		Account: NewAccountService(),
		Market:  NewMarketService(),
		Order:   NewOrderService(),
	}
}
