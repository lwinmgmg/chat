package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/chat/models"
)

func (cV1 *ControllerV1) GetConversations(ctx *gin.Context) {
	userCode, ok := GetUserFromContext(ctx)
	if !ok {
		return
	}
	conv := models.Conversation{}
	convList := []models.ConversationInfo{}
	if err := conv.GetConvByUserId(userCode, &convList, PgDb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    0,
			Message: "Error on get conv list",
		})
		return
	}
	ctx.JSON(http.StatusOK, convList)
}
