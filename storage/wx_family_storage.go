package storage

import (
	"errors"

	"github.com/terloo/xiaochen/storage/models"

	"gorm.io/gorm"
)

type wxFamilyStorage struct {
	db *gorm.DB
}

func newWxFamilyStorage(db *gorm.DB) *wxFamilyStorage {
	return &wxFamilyStorage{db: db}
}

func (r *wxFamilyStorage) Create(family *models.WxFamily) error {
	return r.db.Create(family).Error
}

func (r *wxFamilyStorage) FindAll() ([]*models.WxFamily, error) {
	var families []*models.WxFamily
	if err := r.db.Find(&families).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return families, nil
}

func (r *wxFamilyStorage) FindByID(id uint) (*models.WxFamily, error) {
	var family models.WxFamily
	if err := r.db.First(&family, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &family, nil
}

func (r *wxFamilyStorage) FindByName(name string) (*models.WxFamily, error) {
	var family models.WxFamily
	if err := r.db.Where("name = ?", name).First(&family).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &family, nil
}

func (r *wxFamilyStorage) Update(family *models.WxFamily) error {
	return r.db.Save(family).Error
}

func (r *wxFamilyStorage) Delete(id uint) error {
	return r.db.Delete(&models.WxFamily{}, id).Error
}
