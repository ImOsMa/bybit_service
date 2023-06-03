package bybit_service

type GetKlineItem struct {
	StartTime string `json:"start_time"`
	Open      string `json:"open"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
	Turnover  string `json:"turnover"`
}

type GetKline struct {
	Category string         `json:"category"`
	Symbol   string         `json:"symbol"`
	List     []GetKlineItem `json:"list"`
}

type LotSizeFilter struct {
	BasePrecision  string `json:"base_precision"`
	QuotePrecision string `json:"quote_precision"`
	MaxOrderQty    string `json:"max_order_qty"`
	MinOrderQty    string `json:"min_order_qty"`
	MinOrderAmt    string `json:"min_order_amt"`
	MaxOrderAmt    string `json:"max_order_amt"`
}

type GetInstrumentInfo struct {
	Category      string        `json:"category"`
	Symbol        string        `json:"symbol"`
	BaseCoin      string        `json:"base_coin"`
	QuoteCoin     string        `json:"quote_coin"`
	Innovation    string        `json:"innovation"`
	Status        string        `json:"status"`
	LotSizeFilter LotSizeFilter `json:"lot_size_filter"`
}

type TickersSportResult struct {
	Category      string `json:"category"`
	Symbol        string `json:"symbol"`
	Bid1Price     string `json:"bid1_price"`
	Bid1Size      string `json:"bid1_size"`
	Ask1Price     string `json:"ask1_price"`
	Ask1Size      string `json:"ask1_size"`
	LastPrice     string `json:"last_price"`
	PrevPrice24H  string `json:"prev_price24h"`
	Price24HPcnt  string `json:"price24h_pcnt"`
	HighPrice24H  string `json:"high_price24h"`
	LowPrice24H   string `json:"low_price24h"`
	Turnover24H   string `json:"turnover24h"`
	Volume24H     string `json:"volume24h"`
	UsdIndexPrice string `json:"usd_index_price"`
}

type PositionInfo struct {
	Symbol         string `json:"symbol"`
	Leverage       string `json:"leverage"`
	AvgPrice       string `json:"avg_price"`
	LiqPrice       string `json:"liq_price"`
	RiskLimitValue string `json:"risk_limit_value"`
	TakeProfit     string `json:"take_profit"`
	PositionValue  string `json:"position_value"`
	TpSlMode       string `json:"tpsl_mode"`
	RiskID         int    `json:"risk_id"`
	UnrealisedPnl  string `json:"unrealised_pnl"`
	MarkPrice      string `json:"mark_price"`
	CumRealisedPnl string `json:"cum_realised_pnl"`
	PositionMM     string `json:"position_mm"`
	CreatedTime    string `json:"created_time"`
	PositionIM     string `json:"position_im"`
	UpdatedTime    string `json:"updated_time"`
	Side           string `json:"side"`
	BustPrice      string `json:"bust_price"`
	Size           string `json:"size"`
	PositionStatus string `json:"position_status"`
	StopLoss       string `json:"stop_loss"`
	TradeMode      int    `json:"trade_mode"`
}
