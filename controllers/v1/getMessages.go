package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/chat/models"
)

func (cV1 *ControllerV1) GetMessages(ctx *gin.Context) {
	userCode, ok := GetUserFromContext(ctx)
	if !ok {
		return
	}
	convIdStr := ctx.Param("convId")
	lastMesgStr := ctx.Param("lastMesg")
	limitStr := ctx.Query("limit")
	limit := 10
	if len(limitStr) != 0 {
		nLimit, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = nLimit
		}
	}
	convId, err := strconv.Atoi(convIdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    0,
			Message: fmt.Sprintf("Failed to parse param %v", err),
		})
		return
	}
	convUser := models.ConversationUser{}
	convUserList := []models.ConvUserDetail{}
	if err := convUser.GetByConvId(uint(convId), &convUserList, PgDb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    2,
			Message: fmt.Sprintf("Failed to get conversation user %v", err),
		})
		return
	}
	isAllowed := false
	for _, v := range convUserList {
		if v.UserID == userCode {
			isAllowed = true
		}
	}
	if !isAllowed {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    3,
			Message: "You are not allow to access this conversation",
		})
		return
	}
	mesgs := []models.Message{}
	if err := models.GetMessages(uint(convId), lastMesgStr, int64(limit), &mesgs, MongoDb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    3,
			Message: "Error on fetching messages" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, mesgs)
}
