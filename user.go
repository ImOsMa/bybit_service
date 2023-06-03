package bybit_service

import "time"

type User struct {
	AccountId string `json:"account_id" binding:"required"`
	Token     string `json:"token" binding:"required"`
}

type Permissions struct {
	ContractTrade []string `json:"contract_trade"`
	Spot          []string `json:"spot"`
	Wallet        []string `json:"wallet"`
	Options       []string `json:"options"`
	Derivatives   []string `json:"derivatives"`
	CopyTrading   []string `json:"copy_trading"`
	BlockTrade    []string `json:"block_trade"`
	Exchange      []string `json:"exchange"`
	Nft           []string `json:"nft"`
}

type KeyInformation struct {
	UserID      int         `json:"user_id"`
	VipLevel    string      `json:"vip_level"`
	CreatedAt   time.Time   `json:"created_at"`
	Permissions Permissions `json:"permissions"`
}
