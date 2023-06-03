package handler

import (
	"net/http"
	"strconv"

	"github.com/ImOsMa/bybit_service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) postSpotOrder(c *gin.Context) {
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

	var request bybit_service.PostSpotOrderRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	spotOrder, err := h.services.PostSpotOrder(bybit_service.User{AccountId: accountID, Token: tokenID}, request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, spotOrder)
}

func (h *Handler) getSpotOrder(c *gin.Context) {
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

	orderID := c.Query("order_id")
	if orderID == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid symbol")
		return
	}

	spotOrder, err := h.services.GetSpotOrder(bybit_service.User{AccountId: accountID, Token: tokenID}, orderID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, spotOrder)
}

func (h *Handler) deleteSpotOrder(c *gin.Context) {
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

	orderID := c.Query("order_id")
	if orderID == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid symbol")
		return
	}

	spotOrder, err := h.services.DeleteSpotOrder(bybit_service.User{AccountId: accountID, Token: tokenID}, orderID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, spotOrder)
}

func (h *Handler) openSpotOrder(c *gin.Context) {
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

	spotOrders, err := h.services.GetOpenSpotOrders(bybit_service.User{AccountId: accountID, Token: tokenID}, limitInt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, spotOrders)
}

func (h *Handler) changeOrder(c *gin.Context) {
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

	var request bybit_service.ChangeOrderParams

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	changedOrder, err := h.services.ChangeOrder(bybit_service.User{AccountId: accountID, Token: tokenID}, request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, changedOrder)
}

func (h *Handler) history(c *gin.Context) {}
