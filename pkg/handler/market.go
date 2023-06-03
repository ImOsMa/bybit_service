package handler

import (
	"net/http"
	"strconv"

	"github.com/ImOsMa/bybit_service"
	"github.com/ImOsMa/bybit_service/pkg/client/bybit"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getKline(c *gin.Context) {
	accountID, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tokenID, err := getToken(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var (
		serviceSymbol   bybit.SymbolV5
		serviceInterval bybit.Interval
	)

	serviceSymbol = bybit.SymbolV5(c.Query("symbol"))
	if serviceSymbol == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid symbol")
		return
	}

	serviceInterval = bybit.Interval(c.Query("interval"))
	if serviceInterval == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid symbol")
		return
	}

	limit := c.Query("limit")
	if limit == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid limit")
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid limit")
		return
	}

	kline, err := h.services.GetKline(
		bybit_service.User{AccountId: accountID, Token: tokenID}, serviceSymbol, serviceInterval, limitInt)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, kline)
}

func (h *Handler) instrumentInfo(c *gin.Context) {
	accountID, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tokenID, err := getToken(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var (
		serviceSymbol bybit.SymbolV5
	)

	serviceSymbol = bybit.SymbolV5(c.Query("symbol"))
	if serviceSymbol == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid symbol")
		return
	}

	info, err := h.services.InstrumentInfo(bybit_service.User{AccountId: accountID, Token: tokenID}, serviceSymbol)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, info)
}

func (h *Handler) tickers(c *gin.Context) {
	accountID, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tokenID, err := getToken(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var (
		serviceSymbol bybit.SymbolV5
	)

	serviceSymbol = bybit.SymbolV5(c.Query("symbol"))
	if serviceSymbol == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid symbol")
		return
	}

	tickers, err := h.services.Tickers(bybit_service.User{AccountId: accountID, Token: tokenID}, serviceSymbol)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tickers)
}

func (h *Handler) coinInfo(c *gin.Context) {}

func (h *Handler) positionInfo(c *gin.Context) {
	accountID, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tokenID, err := getToken(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var (
		serviceSymbol bybit.SymbolV5
	)

	serviceSymbol = bybit.SymbolV5(c.Query("symbol"))
	if serviceSymbol == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid symbol")
		return
	}

	positionInfo, err := h.services.PositionInfo(bybit_service.User{AccountId: accountID, Token: tokenID}, serviceSymbol)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, positionInfo)
}
