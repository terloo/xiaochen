package storage

import (
	"errors"

	"github.com/terloo/xiaochen/storage/models"

	"gorm.io/gorm"
)

type wxChatStorage struct {
	db *gorm.DB
}

func newWxChatStorage(db *gorm.DB) *wxChatStorage {
	return &wxChatStorage{db: db}
}

func (r *wxChatStorage) Create(chat *models.WxChat) error {
	return r.db.Create(chat).Error
}

func (r *wxChatStorage) FindByID(id uint) (*models.WxChat, error) {
	var chat models.WxChat
	if err := r.db.First(&chat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &chat, nil
}

func (r *wxChatStorage) FindByName(name string) (*models.WxChat, error) {
	var chat models.WxChat
	if err := r.db.Where("name = ?", name).First(&chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &chat, nil
}

func (r *wxChatStorage) Update(chat *models.WxChat) error {
	return r.db.Save(chat).Error
}

func (r *wxChatStorage) Delete(id uint) error {
	return r.db.Delete(&models.WxChat{}, id).Error
}
