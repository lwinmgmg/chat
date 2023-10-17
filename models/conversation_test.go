package models_test

import (
	"errors"
	"testing"

	"github.com/lwinmgmg/chat/models"
	"gorm.io/gorm"
)

func TestFindNormalConversation(t *testing.T) {
	models.PgDb.Transaction(func(tx *gorm.DB) error {
		var uid1 = "abc"
		var uid2 = "def"

		convId, err := models.FindNormalConversation(uid1, uid2, tx)
		if convId.ID != 0 || err == nil {
			t.Error("Expect convId as 0 but getting :", convId, err)
		}
		conv, _, _, _ := models.CreateNewNormalConv(uid1, uid2, tx)
		convId1, err := models.FindNormalConversation(uid1, uid2, tx)
		if convId1.ID != conv.ID || err != nil {
			t.Error("Expected :", conv.ID, "Getting :", convId1)
		}
		return errors.New("To Roll back")
	})
}

func TestGetNormalConversation(t *testing.T) {
	models.PgDb.Transaction(func(tx *gorm.DB) error {
		var uid1 = "abc"
		var uid2 = "def"
		conv, _, _, _ := models.CreateNewNormalConv(uid1, uid2, tx)
		convId, err := models.GetNormalConversation(uid1, uid2, tx)
		if err != nil {
			t.Error("Error on getting :", err)
		}
		if conv.ID != convId.ID {
			t.Error("Expected equal manual created ID and convId", conv.ID, convId)
		}
		uid3 := "efg"
		convId1, err := models.GetNormalConversation(uid1, uid3, tx)
		if err != nil {
			t.Error("Error on getting :", err)
		}
		if convId1.ID == 0 {
			t.Error("Expected greater than 0, getting :", convId1)
		}

		return errors.New("To Roll back")
	})
}
