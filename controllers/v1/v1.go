package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/chat/models"
)

type ControllerV1 struct {
	Router *gin.RouterGroup
}

func (cV1 *ControllerV1) Serve() {
	cV1.Router.POST("/conversations", cV1.PostConversation)
}

func GetUserFromContext(ctx *gin.Context) (string, bool) {
	userCode, ok := ctx.Get("userCode")
	userCodeStr, ok1 := userCode.(string)
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.DefaultResponse{
			Code:    1,
			Message: "Authorization Required!",
		})
		return "", false
	}
	return userCodeStr, true
}
