package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/chat/models"
)

func (cV1 *ControllerV1) GetConversationByID(ctx *gin.Context) {
	userCode, ok := GetUserFromContext(ctx)
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
	conv := models.Conversation{}
	if err := conv.Get(uint(convId), PgDb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    1,
			Message: fmt.Sprintf("Failed to get conversation %v", err),
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
	ctx.JSON(http.StatusOK, models.ConversationDetail{
		ID:               conv.ID,
		Name:             conv.Name,
		ConversationType: conv.ConType,
		Active:           conv.Active,
		UserID:           conv.UserID,
		ImageURL:         conv.ImageURL,
		ConvUsers:        convUserList,
	})
}
