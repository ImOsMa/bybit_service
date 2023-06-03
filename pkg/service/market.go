package service

import (
	"fmt"

	"github.com/ImOsMa/bybit_service"
	"github.com/ImOsMa/bybit_service/pkg/client/bybit"
)

type MarketService struct {
}

func NewMarketService() *MarketService {
	return &MarketService{}
}

func (m *MarketService) GetKline(user bybit_service.User, symbol bybit.SymbolV5, interval bybit.Interval, limit int) (bybit_service.GetKline, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	klineData, err := client.V5().Market().GetKline(bybit.V5GetKlineParam{
		Category: bybit.CategoryV5Spot,
		Symbol:   symbol,
		Interval: interval,
		Limit:    &limit,
	})

	if err != nil {
		return bybit_service.GetKline{}, err
	}

	if klineData.RetMsg != "OK" && klineData.RetMsg != "" {
		return bybit_service.GetKline{}, fmt.Errorf("error in getting kline info")
	}

	klineItems := make([]bybit_service.GetKlineItem, 0)
	for _, item := range klineData.Result.List {
		klineItems = append(klineItems, bybit_service.GetKlineItem{
			StartTime: item.StartTime,
			Open:      item.Open,
			High:      item.High,
			Low:       item.Low,
			Close:     item.Close,
			Volume:    item.Volume,
			Turnover:  item.Turnover,
		})
	}

	return bybit_service.GetKline{
		Category: string(klineData.Result.Category),
		Symbol:   string(klineData.Result.Symbol),
		List:     klineItems,
	}, nil
}

func (m *MarketService) InstrumentInfo(user bybit_service.User, symbol bybit.SymbolV5) (bybit_service.GetInstrumentInfo, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	instrumentInfo, err := client.V5().Market().GetInstrumentsInfo(bybit.V5GetInstrumentsInfoParam{
		Category: bybit.CategoryV5Spot,
		Symbol:   &symbol,
	})

	if err != nil {
		return bybit_service.GetInstrumentInfo{}, err
	}

	if instrumentInfo.RetMsg != "OK" && instrumentInfo.RetMsg != "" {
		return bybit_service.GetInstrumentInfo{}, fmt.Errorf("error in getting instrument info")
	}

	if len(instrumentInfo.Result.Spot.List) == 0 {
		return bybit_service.GetInstrumentInfo{}, fmt.Errorf("empty response of instrument info")
	}
	return bybit_service.GetInstrumentInfo{
		Category:   string(instrumentInfo.Result.Spot.Category),
		Symbol:     string(instrumentInfo.Result.Spot.List[0].Symbol),
		BaseCoin:   string(instrumentInfo.Result.Spot.List[0].BaseCoin),
		QuoteCoin:  string(instrumentInfo.Result.Spot.List[0].QuoteCoin),
		Innovation: string(instrumentInfo.Result.Spot.List[0].Innovation),
		Status:     string(instrumentInfo.Result.Spot.List[0].Status),
		LotSizeFilter: bybit_service.LotSizeFilter{
			BasePrecision:  instrumentInfo.Result.Spot.List[0].LotSizeFilter.BasePrecision,
			QuotePrecision: instrumentInfo.Result.Spot.List[0].LotSizeFilter.QuotePrecision,
			MaxOrderQty:    instrumentInfo.Result.Spot.List[0].LotSizeFilter.MaxOrderQty,
			MinOrderQty:    instrumentInfo.Result.Spot.List[0].LotSizeFilter.MinOrderQty,
			MinOrderAmt:    instrumentInfo.Result.Spot.List[0].LotSizeFilter.MinOrderAmt,
			MaxOrderAmt:    instrumentInfo.Result.Spot.List[0].LotSizeFilter.MaxOrderQty,
		},
	}, nil
}

func (m *MarketService) Tickers(user bybit_service.User, symbol bybit.SymbolV5) (bybit_service.TickersSportResult, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	tickers, err := client.V5().Market().GetTickers(bybit.V5GetTickersParam{
		Category: bybit.CategoryV5Spot,
		Symbol:   &symbol,
	})

	if err != nil {
		return bybit_service.TickersSportResult{}, err
	}

	if tickers.RetMsg != "OK" && tickers.RetMsg != "" {
		return bybit_service.TickersSportResult{}, fmt.Errorf("error in getting instrument info")
	}

	if len(tickers.Result.Spot.List) == 0 {
		return bybit_service.TickersSportResult{}, fmt.Errorf("empty response of instrument info")
	}

	mainElement := tickers.Result.Spot.List[0]
	return bybit_service.TickersSportResult{
		Category:      string(tickers.Result.Spot.Category),
		Symbol:        string(mainElement.Symbol),
		Bid1Price:     mainElement.Bid1Price,
		Bid1Size:      mainElement.Bid1Size,
		Ask1Price:     mainElement.Ask1Price,
		Ask1Size:      mainElement.Ask1Size,
		LastPrice:     mainElement.LastPrice,
		PrevPrice24H:  mainElement.PrevPrice24H,
		Price24HPcnt:  mainElement.Price24HPcnt,
		HighPrice24H:  mainElement.HighPrice24H,
		LowPrice24H:   mainElement.LowPrice24H,
		Turnover24H:   mainElement.Turnover24H,
		Volume24H:     mainElement.Volume24H,
		UsdIndexPrice: mainElement.UsdIndexPrice,
	}, nil
}

func (m *MarketService) PositionInfo(user bybit_service.User, symbol bybit.SymbolV5) (bybit_service.PositionInfo, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	positionInfo, err := client.V5().Position().GetPositionInfo(bybit.V5GetPositionInfoParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   &symbol,
	})

	if err != nil {
		return bybit_service.PositionInfo{}, err
	}

	if positionInfo.RetMsg != "OK" && positionInfo.RetMsg != "" {
		return bybit_service.PositionInfo{}, fmt.Errorf("error in getting instrument info")
	}

	if len(positionInfo.Result.List) == 0 {
		return bybit_service.PositionInfo{}, fmt.Errorf("empty response of instrument info")
	}

	mainElement := positionInfo.Result.List[0]
	return bybit_service.PositionInfo{
		Symbol:         string(mainElement.Symbol),
		Leverage:       mainElement.Leverage,
		AvgPrice:       mainElement.AvgPrice,
		LiqPrice:       mainElement.LiqPrice,
		RiskLimitValue: mainElement.RiskLimitValue,
		TakeProfit:     mainElement.TakeProfit,
		PositionValue:  mainElement.PositionValue,
		TpSlMode:       string(mainElement.TpSlMode),
		RiskID:         mainElement.RiskID,
		UnrealisedPnl:  mainElement.UnrealisedPnl,
		MarkPrice:      mainElement.MarkPrice,
		CumRealisedPnl: mainElement.CumRealisedPnl,
		PositionMM:     mainElement.PositionMM,
		CreatedTime:    mainElement.CreatedTime,
		PositionIM:     mainElement.PositionIM,
		UpdatedTime:    mainElement.UpdatedTime,
		Side:           string(mainElement.Side),
		BustPrice:      mainElement.BustPrice,
		Size:           mainElement.Size,
		PositionStatus: mainElement.PositionStatus,
		StopLoss:       mainElement.StopLoss,
		TradeMode:      mainElement.TradeMode,
	}, nil
}

func (m *MarketService) CoinInfo(user bybit_service.User, coin string) {}
