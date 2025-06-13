package storage

import (
	"log"

	"github.com/terloo/xiaochen/storage/db"
	"github.com/terloo/xiaochen/storage/models"
)

var MusicRepo = newMusicStorage(db.DB)
var WxFamilyRepo = newWxFamilyStorage(db.DB)
var WxChatRepo = newWxChatStorage(db.DB)

func init() {

	var err error
	// 迁移model
	err = db.DB.AutoMigrate(&models.Music{})
	if err != nil {
		log.Fatalf("failed to migrate model: %+v", err)
	}
	err = db.DB.AutoMigrate(&models.WxFamily{})
	if err != nil {
		log.Fatalf("failed to migrate model: %+v", err)
	}
	err = db.DB.AutoMigrate(&models.WxChat{})
	if err != nil {
		log.Fatalf("failed to migrate model: %+v", err)
	}
}
