package bybit_service

type AccountInfo struct {
	MarginMode          string `json:"margin_mode" binding:"required"`
	UpdatedTime         string `json:"updated_time" binding:"required"`
	UnifiedMarginStatus string `json:"unified_margin_status" binding:"required"`
}

type AccountWalletBalance struct {
	Coin     string `json:"coin"`
	CoinID   string `json:"coin_id"`
	CoinName string `json:"coin_name"`
	Total    string `json:"total"`
	Free     string `json:"free"`
	Locked   string `json:"locked"`
}
