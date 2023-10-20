package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/chat/grpc/client"
	"github.com/lwinmgmg/chat/models"
)

func ParseToken(keyString, tokenType string) (string, error) {
	if keyString == "" {
		return "", errors.New("empty token")
	}
	inputTokenType := keyString[0:len(tokenType)]
	inputTokenString := keyString[len(tokenType):]
	if inputTokenType != tokenType {
		return "", errors.New("invalid token")
	}
	return strings.TrimSpace(inputTokenString), nil
}

func JwtAuthMiddleware(tokenType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyString := ctx.Request.Header.Get("Authorization")
		inputTokenString, err := ParseToken(keyString, tokenType)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.DefaultResponse{
				Code:    2,
				Message: fmt.Sprintf("Authorization Required! [%v]", keyString),
			})
			return
		}
		user, err := client.GetUserByToken(inputTokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.DefaultResponse{
				Code:    1,
				Message: "Authorization Required!",
			})
			return
		}
		ctx.Set("userCode", user.Code)
	}
}
