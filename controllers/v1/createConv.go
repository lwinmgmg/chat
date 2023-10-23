package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/chat/grpc/client"
	"github.com/lwinmgmg/chat/models"
	"github.com/lwinmgmg/chat/services"
)

func checkUserExist(uids []string) error {
	for _, uid := range uids {
		if _, err := client.GetUserByCode(uid); err != nil {
			return err
		}
	}
	return nil
}

func (cV1 *ControllerV1) PostConversation(ctx *gin.Context) {
	userCode, ok := GetUserFromContext(ctx)
	if !ok {
		return
	}
	convData := models.ConversationCreateData{}
	if err := ctx.ShouldBindJSON(&convData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    0,
			Message: fmt.Sprintf("Request must be json format [%v]", err.Error()),
		})
		return
	}
	if err := checkUserExist(convData.UserList); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.DefaultResponse{
			Code:    0,
			Message: fmt.Sprintf("Record not found [%v]", err.Error()),
		})
		return
	}
	var conv *models.Conversation
	var err error
	switch convData.ConvType {
	case models.NormalCon:
		if len(convData.UserList) != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, models.DefaultResponse{
				Code:    0,
				Message: "Request does not have receiver",
			})
			return
		}
		conv, err = models.GetNormalConversation(userCode, convData.UserList[0], services.PgDb)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
				Code:    1,
				Message: "Can't create conversation " + err.Error(),
			})
			return
		}
	case models.GroupCon:
		conv, err = models.CreateNewGroupConv(userCode, convData.UserList, services.PgDb)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
				Code:    1,
				Message: "Can't create conversation " + err.Error(),
			})
			return
		}
	}
	convUserList := []models.ConvUserDetail{}
	convUser := models.ConversationUser{}
	if err := convUser.GetByConvId(conv.ID, &convUserList, PgDb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    2,
			Message: fmt.Sprintf("Failed to get conversation user %v", err),
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
		LastMesgId:       conv.LastMesgID,
		ConvUsers:        convUserList,
	})
}
