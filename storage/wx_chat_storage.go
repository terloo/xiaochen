package storage

import (
	"errors"

	"github.com/terloo/xiaochen/storage/models"

	"gorm.io/gorm"
)

type WxChatStorage struct {
	db *gorm.DB
}

func newWxChatStorage(db *gorm.DB) *WxChatStorage {
	return &WxChatStorage{db: db}
}

func (r *WxChatStorage) Create(chat *models.WxChat) error {
	return r.db.Create(chat).Error
}

func (r *WxChatStorage) FindByID(id uint) (*models.WxChat, error) {
	var chat models.WxChat
	if err := r.db.First(&chat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &chat, nil
}

func (r *WxChatStorage) FindByName(name string) (*models.WxChat, error) {
	var chat models.WxChat
	if err := r.db.Where("name = ?", name).First(&chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &chat, nil
}

func (r *WxChatStorage) Update(chat *models.WxChat) error {
	return r.db.Save(chat).Error
}

func (r *WxChatStorage) Delete(id uint) error {
	return r.db.Delete(&models.WxChat{}, id).Error
}
