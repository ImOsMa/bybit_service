package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	accountId = "AccountId"
	token     = "Access-Token"
)

func getAccountId(c *gin.Context) (string, error) {
	id := c.GetHeader(accountId)
	if id == "" {
		return "", errors.New("accountId is not exists")
	}

	if len(id) == 0 {
		return "", errors.New("empty accountId")
	}

	return id, nil
}

func getToken(c *gin.Context) (string, error) {
	id := c.GetHeader(token)
	if id == "" {
		return "", errors.New("token is not exists")
	}

	if len(id) == 0 {
		return "", errors.New("empty token")
	}

	return id, nil
}
