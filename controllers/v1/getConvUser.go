package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/chat/models"
)

func (cV1 *ControllerV1) GetConversation(ctx *gin.Context) {
	_, ok := GetUserFromContext(ctx)
	if !ok {
		return
	}
	convIdStr := ctx.Param("id")
	convId, err := strconv.Atoi(convIdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    0,
			Message: fmt.Sprintf("Failed to parse param %v", err),
		})
		return
	}
	convUsers := []models.ConvUserData{}
	convUser := models.ConversationUser{}
	if err := convUser.GetByConvId(uint(convId), &convUsers, PgDb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    1,
			Message: fmt.Sprintf("Failed to parse param %v", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, convUsers)
}
