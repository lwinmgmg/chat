package models_test

import (
	"testing"

	"github.com/lwinmgmg/chat/models"
)

func TestGetMessage(t *testing.T) {
	var convId uint = 1000000
	var mesg1 = models.Message{
		ConversationID: convId,
	}
	if err := mesg1.Create(models.MongoDb); err != nil {
		t.Error(err)
	}
	defer mesg1.Delete(models.MongoDb)
	var mesg2 = models.Message{
		ConversationID: convId,
	}
	if err := mesg2.Create(models.MongoDb); err != nil {
		t.Error(err)
	}
	defer mesg2.Delete(models.MongoDb)
	res1 := []map[string]any{}
	err := models.GetMessages(convId, mesg1.ID.Hex(), 10, &res1, models.MongoDb)
	if err != nil {
		t.Error()
	}
	if len(res1) != 1 || err != nil {
		t.Error("Fetch Test 2 : Expected one record, getting :", res1, err, mesg1)
	}
	res2 := []map[string]any{}
	err = models.GetMessages(convId, mesg2.ID.Hex(), 10, &res2, models.MongoDb)
	if err != nil {
		t.Error()
	}
	if len(res2) != 2 || err != nil {
		t.Error("Fetch Test 1 : Expected two record, getting :", res2, err, mesg2)
	}
	res3 := []map[string]any{}
	err = models.GetMessages(convId, mesg1.ID.Hex(), 1, &res3, models.MongoDb)
	if err != nil {
		t.Error()
	}
	if len(res3) != 1 || err != nil {
		t.Error("Limit test : Expected one record, getting :", res3, err, mesg2)
	}
}
