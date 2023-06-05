package bybit_service

import "github.com/ImOsMa/bybit_service/pkg/client/bybit"

type PostSpotOrderRequest struct {
	Symbol bybit.SymbolSpot    `json:"symbol"`
	Qty    float64             `json:"qty"`
	Side   bybit.Side          `json:"side"`
	Type   bybit.OrderTypeSpot `json:"type"`
	Price  *float64            `json:"price"`
}

type PostSpotOrderResult struct {
	OrderID      string `json:"order_id"`
	Symbol       string `json:"symbol"`
	TransactTime string `json:"transact_time"`
	Price        string `json:"price"`
	OrigQty      string `json:"orig_qty"`
	Type         string `json:"type"`
	Side         string `json:"side"`
	Status       string `json:"status"`
	TimeInForce  string `json:"time_in_force"`
	AccountID    string `json:"account_id"`
	SymbolName   string `json:"symbol_name"`
	ExecutedQty  string `json:"executed_qty"`
}

type SpotGetOrderResult struct {
	AccountId           string `json:"account_id"`
	ExchangeId          string `json:"exchange_id"`
	Symbol              string `json:"symbol"`
	SymbolName          string `json:"symbol_name"`
	OrderLinkId         string `json:"order_link_id"`
	OrderId             string `json:"order_id"`
	Price               string `json:"price"`
	OrigQty             string `json:"orig_qty"`
	ExecutedQty         string `json:"executed_qty"`
	CummulativeQuoteQty string `json:"cummulative_quote_qty"`
	AvgPrice            string `json:"avg_price"`
	Status              string `json:"status"`
	TimeInForce         string `json:"time_in_force"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stop_price"`
	IcebergQty          string `json:"iceberg_qty"`
	Time                string `json:"time"`
	UpdateTime          string `json:"update_time"`
	IsWorking           bool   `json:"is_working"`
}

type SpotDeleteResult struct {
	OrderId      string `json:"order_id"`
	Symbol       string `json:"symbol"`
	Status       string `json:"status"`
	AccountId    string `json:"account_id"`
	TransactTime string `json:"transact_time"`
	Price        string `json:"price"`
	OrigQty      string `json:"orig_qty"`
	ExecutedQty  string `json:"executed_qty"`
	TimeInForce  string `json:"time_in_force"`
	Type         string `json:"type"`
	Side         string `json:"side"`
}

type SpotOpenOrdersResult struct {
	AccountID           string `json:"account_id"`
	ExchangeID          string `json:"exchange_id"`
	Symbol              string `json:"symbol"`
	SymbolName          string `json:"symbol_name"`
	OrderID             string `json:"order_id"`
	Price               string `json:"price"`
	OrigQty             string `json:"orig_qty"`
	ExecutedQty         string `json:"executed_qty"`
	CummulativeQuoteQty string `json:"cummulative_quote_qty"`
	AvgPrice            string `json:"avg_price"`
	Status              string `json:"status"`
	TimeInForce         string `json:"time_in_force"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stop_price"`
	IcebergQty          string `json:"iceberg_qty"`
	Time                string `json:"time"`
	UpdateTime          string `json:"update_time"`
	IsWorking           bool   `json:"isWorking"`
}

type ChangeOrderParams struct {
	OrderID      string  `json:"order_id"`
	Qty          *string `json:"qty,omitempty"`
	Price        *string `json:"price,omitempty"`
	StopLoss     *string `json:"stop_loss,omitempty"`
	TriggerPrice *string `json:"trigger_price,omitempty"`
}
