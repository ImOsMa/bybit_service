package service

import (
	"fmt"

	"github.com/ImOsMa/bybit_service"
	"github.com/ImOsMa/bybit_service/pkg/client/bybit"
)

type AccountService struct {
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (a *AccountService) AccountInfo(user bybit_service.User) (bybit_service.AccountInfo, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	accountInfo, err := client.V5().Account().GetAccountInfo()
	if err != nil {
		return bybit_service.AccountInfo{}, err
	}

	if accountInfo.RetMsg != "OK" && accountInfo.RetMsg != "" {
		return bybit_service.AccountInfo{}, fmt.Errorf("error in getting account info")
	}

	return bybit_service.AccountInfo{
		MarginMode:          string(accountInfo.Result.MarginMode),
		UpdatedTime:         accountInfo.Result.UpdatedTime,
		UnifiedMarginStatus: "UnifiedMarginStatusRegular",
	}, nil
}

func (a *AccountService) SpotWalletBalance(user bybit_service.User) ([]bybit_service.AccountWalletBalance, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	walletBalance, err := client.Spot().V1().SpotGetWalletBalance()

	if err != nil {
		return nil, err
	}

	if walletBalance.RetMsg != "OK" && walletBalance.RetMsg != "" {
		return nil, fmt.Errorf("error in getting account wallet balance")
	}

	spotBalances := make([]bybit_service.AccountWalletBalance, 0)
	for _, balance := range walletBalance.Result.Balances {
		if balance.Coin == "" {
			continue
		}
		spotBalances = append(spotBalances, bybit_service.AccountWalletBalance{
			Coin:     balance.Coin,
			CoinID:   balance.CoinID,
			CoinName: balance.CoinName,
			Total:    balance.Total,
			Free:     balance.Free,
			Locked:   balance.Locked,
		})
	}

	return spotBalances, nil
}

func (a *AccountService) FeeRate(user bybit_service.User) (string, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	accountInfo, err := client.V5().Account().GetAccountInfo()
	if err != nil {
		return "", err
	}

	if accountInfo.RetMsg != "OK" && accountInfo.RetMsg != "" {
		return "", fmt.Errorf("error in getting account info")
	}

	return "0.001", nil
}

func (a *AccountService) KeyInformation(user bybit_service.User) (bybit_service.KeyInformation, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	keyInformation, err := client.V5().User().GetAPIKey()

	if err != nil {
		return bybit_service.KeyInformation{}, err
	}

	if keyInformation.RetMsg != "OK" && keyInformation.RetMsg != "" {
		return bybit_service.KeyInformation{}, fmt.Errorf("error in getting account key information")
	}

	return bybit_service.KeyInformation{
		UserID:    keyInformation.Result.UserID,
		VipLevel:  keyInformation.Result.VipLevel,
		CreatedAt: keyInformation.Result.CreatedAt,
		Permissions: bybit_service.Permissions{
			ContractTrade: keyInformation.Result.Permissions.ContractTrade,
			Spot:          keyInformation.Result.Permissions.Spot,
			Wallet:        keyInformation.Result.Permissions.Wallet,
			Options:       keyInformation.Result.Permissions.Options,
			Derivatives:   keyInformation.Result.Permissions.Derivatives,
			CopyTrading:   keyInformation.Result.Permissions.CopyTrading,
			BlockTrade:    keyInformation.Result.Permissions.BlockTrade,
			Exchange:      keyInformation.Result.Permissions.Exchange,
			Nft:           keyInformation.Result.Permissions.Nft,
		},
	}, nil
}

func (a *AccountService) CoinExchangeRecords(user bybit_service.User) {}
