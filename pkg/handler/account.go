package handler

import (
	"net/http"

	"github.com/ImOsMa/bybit_service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) spotWalletBalance(c *gin.Context) {
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

	balance, err := h.services.SpotWalletBalance(bybit_service.User{AccountId: accountID, Token: tokenID})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, balance)
}

func (h *Handler) info(c *gin.Context) {
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

	info, err := h.services.AccountInfo(bybit_service.User{AccountId: accountID, Token: tokenID})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, info)
}

func (h *Handler) keyInformation(c *gin.Context) {
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

	keyInfo, err := h.services.KeyInformation(bybit_service.User{AccountId: accountID, Token: tokenID})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, keyInfo)
}

func (h *Handler) coinExchangeRecords(c *gin.Context) {}

func (h *Handler) feeRate(c *gin.Context) {
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

	rate, err := h.services.FeeRate(bybit_service.User{AccountId: accountID, Token: tokenID})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, rate)
}
