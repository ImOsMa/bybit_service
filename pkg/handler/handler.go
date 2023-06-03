package handler

import (
	"github.com/ImOsMa/bybit_service/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	account := router.Group("/account")
	{
		account.GET("/spot_wallet_balance", h.spotWalletBalance)
		account.GET("/info", h.info)
		account.GET("/fee_rate", h.feeRate)
		account.GET("/key_information", h.keyInformation)
		account.GET("/coin_exchange_records", h.coinExchangeRecords)
	}

	order := router.Group("/order")
	{
		order.POST("/post_spot_order", h.postSpotOrder)
		order.GET("/get_spot_order", h.getSpotOrder)
		order.DELETE("/delete_spot_order", h.deleteSpotOrder)
		order.GET("/open_spot_order", h.openSpotOrder)
		order.PUT("/change_order", h.changeOrder)
		order.GET("/history", h.history)
	}

	market := router.Group("/market")
	{
		market.GET("/get_kline", h.getKline)
		market.GET("/instrument_info", h.instrumentInfo)
		market.GET("/tickers", h.tickers)
		market.GET("/position_info", h.positionInfo)
		market.GET("/coin_info", h.coinInfo)
	}

	return router
}
