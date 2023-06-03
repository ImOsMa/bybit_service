package service

import (
	"fmt"

	"github.com/ImOsMa/bybit_service"
	"github.com/ImOsMa/bybit_service/pkg/client/bybit"
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) PostSpotOrder(user bybit_service.User, request bybit_service.PostSpotOrderRequest) (bybit_service.PostSpotOrderResult, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	var param bybit.SpotPostOrderParam
	if request.Price == nil {
		param = bybit.SpotPostOrderParam{
			Symbol: request.Symbol,
			Qty:    request.Qty,
			Side:   request.Side,
			Type:   request.Type,
		}

	} else {
		param = bybit.SpotPostOrderParam{
			Symbol: request.Symbol,
			Qty:    request.Qty,
			Side:   request.Side,
			Type:   request.Type,
			Price:  request.Price,
		}
	}

	spotOrder, err := client.Spot().V1().SpotPostOrder(param)
	if err != nil {
		return bybit_service.PostSpotOrderResult{}, err
	}

	if spotOrder.RetMsg != "OK" && spotOrder.RetMsg != "" {
		return bybit_service.PostSpotOrderResult{}, fmt.Errorf("error in post spot order")
	}

	if spotOrder.Result.OrderID == "" {
		return bybit_service.PostSpotOrderResult{}, fmt.Errorf("empty order ID")
	}

	return bybit_service.PostSpotOrderResult{
		OrderID:      spotOrder.Result.OrderID,
		Symbol:       spotOrder.Result.Symbol,
		TransactTime: spotOrder.Result.TransactTime,
		Price:        spotOrder.Result.Price,
		OrigQty:      spotOrder.Result.OrigQty,
		Type:         string(spotOrder.Result.Type),
		Side:         spotOrder.Result.Side,
		Status:       string(spotOrder.Result.Status),
		TimeInForce:  string(spotOrder.Result.TimeInForce),
		AccountID:    spotOrder.Result.AccountID,
		SymbolName:   spotOrder.Result.SymbolName,
		ExecutedQty:  spotOrder.Result.ExecutedQty,
	}, nil
}

func (o *OrderService) GetSpotOrder(user bybit_service.User, orderID string) (bybit_service.SpotGetOrderResult, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	spotOrder, err := client.Spot().V1().SpotGetOrder(bybit.SpotGetOrderParam{
		OrderID: &orderID,
	})

	if err != nil {
		return bybit_service.SpotGetOrderResult{}, err
	}

	if spotOrder.RetMsg != "OK" && spotOrder.RetMsg != "" {
		return bybit_service.SpotGetOrderResult{}, fmt.Errorf("error in get spot order")
	}

	if spotOrder.Result.AccountId == "" || spotOrder.Result.OrderId == "" {
		return bybit_service.SpotGetOrderResult{}, fmt.Errorf("empty order ID or account ID")
	}

	result := spotOrder.Result
	return bybit_service.SpotGetOrderResult{
		AccountId:           result.AccountId,
		ExchangeId:          result.ExchangeId,
		Symbol:              result.Symbol,
		SymbolName:          result.SymbolName,
		OrderLinkId:         result.OrderLinkId,
		OrderId:             result.OrderId,
		Price:               result.Price,
		OrigQty:             result.OrigQty,
		ExecutedQty:         result.ExecutedQty,
		CummulativeQuoteQty: result.CummulativeQuoteQty,
		AvgPrice:            result.AvgPrice,
		Status:              result.Status,
		TimeInForce:         result.TimeInForce,
		Type:                result.Type,
		Side:                result.Side,
		StopPrice:           result.StopPrice,
		IcebergQty:          result.IcebergQty,
		Time:                result.Time,
		UpdateTime:          result.UpdateTime,
		IsWorking:           result.IsWorking,
	}, nil
}

func (o *OrderService) DeleteSpotOrder(user bybit_service.User, orderID string) (bybit_service.SpotDeleteResult, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	spotOrder, err := client.Spot().V1().SpotDeleteOrder(bybit.SpotDeleteOrderParam{
		OrderID: &orderID,
	})

	if err != nil {
		return bybit_service.SpotDeleteResult{}, err
	}

	if spotOrder.RetMsg != "OK" && spotOrder.RetMsg != "" {
		return bybit_service.SpotDeleteResult{}, fmt.Errorf("error in delete spot order")
	}

	if spotOrder.Result.AccountId == "" || spotOrder.Result.OrderId == "" {
		return bybit_service.SpotDeleteResult{}, fmt.Errorf("empty order ID or account ID")
	}

	result := spotOrder.Result
	return bybit_service.SpotDeleteResult{
		OrderId:      result.OrderId,
		Symbol:       result.Symbol,
		Status:       result.Status,
		AccountId:    result.AccountId,
		TransactTime: result.TransactTime,
		Price:        result.Price,
		OrigQty:      result.OrigQty,
		ExecutedQty:  result.ExecutedQty,
		TimeInForce:  result.TimeInForce,
		Type:         result.Type,
		Side:         result.Side,
	}, nil
}

func (o *OrderService) GetOpenSpotOrders(user bybit_service.User, limit int) ([]bybit_service.SpotOpenOrdersResult, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	openOrders, err := client.Spot().V1().SpotOpenOrders(bybit.SpotOpenOrdersParam{Limit: &limit})

	if err != nil {
		return nil, err
	}

	if openOrders.RetMsg != "OK" && openOrders.RetMsg != "" {
		return nil, fmt.Errorf("error get open spot orders")
	}

	if len(openOrders.Result) == 0 {
		return nil, nil
	}

	resultOrders := make([]bybit_service.SpotOpenOrdersResult, 0)
	for _, order := range openOrders.Result {
		resultOrders = append(resultOrders, bybit_service.SpotOpenOrdersResult{
			AccountID:           order.AccountID,
			ExchangeID:          order.ExchangeID,
			Symbol:              order.Symbol,
			SymbolName:          order.SymbolName,
			OrderID:             order.OrderID,
			Price:               order.Price,
			OrigQty:             order.OrigQty,
			ExecutedQty:         order.ExecutedQty,
			CummulativeQuoteQty: order.CummulativeQuoteQty,
			AvgPrice:            order.AvgPrice,
			Status:              order.Status,
			TimeInForce:         order.TimeInForce,
			Type:                order.Type,
			Side:                order.Side,
			StopPrice:           order.StopPrice,
			IcebergQty:          order.IcebergQty,
			Time:                order.Time,
			UpdateTime:          order.UpdateTime,
			IsWorking:           order.IsWorking,
		})
	}

	return resultOrders, nil
}

func (o *OrderService) ChangeOrder(user bybit_service.User, params bybit_service.ChangeOrderParams) (string, error) {
	client := bybit.NewTestClient().WithAuth(user.AccountId, user.Token)
	spotOrder, err := client.Spot().V1().SpotGetOrder(bybit.SpotGetOrderParam{
		OrderID: &params.OrderID,
	})

	if err != nil {
		return "", err
	}

	if spotOrder.RetMsg != "OK" && spotOrder.RetMsg != "" {
		return "", fmt.Errorf("error in get spot order")
	}

	if spotOrder.Result.AccountId == "" || spotOrder.Result.OrderId == "" {
		return "", fmt.Errorf("empty order ID or account ID")
	}

	changedOrder, err := client.V5().Order().AmendOrder(bybit.V5AmendOrderParam{
		Category:     bybit.CategoryV5Spot,
		Symbol:       bybit.SymbolV5(spotOrder.Result.Symbol),
		OrderID:      &spotOrder.Result.OrderId,
		Qty:          params.Qty,
		Price:        params.Price,
		StopLoss:     params.StopLoss,
		TriggerPrice: params.TriggerPrice,
	})

	if err != nil {
		return "", err
	}

	if changedOrder.RetMsg != "OK" && changedOrder.RetMsg != "" {
		return "", fmt.Errorf("error in changing spot order")
	}

	if changedOrder.Result.OrderID == "" {
		return "", fmt.Errorf("empty order id in response")
	}

	return changedOrder.Result.OrderID, nil
}

func (o *OrderService) History(user bybit_service.User) {}
