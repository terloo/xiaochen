package storage

import (
	"errors"

	"github.com/terloo/xiaochen/storage/models"

	"gorm.io/gorm"
)

type MusicStorage struct {
	db *gorm.DB
}

func newMusicStorage(db *gorm.DB) *MusicStorage {
	return &MusicStorage{db: db}
}

func (r *MusicStorage) Create(music *models.Music) error {
	return r.db.Create(music).Error
}

func (r *MusicStorage) FindByID(id uint) (*models.Music, error) {
	var music models.Music
	if err := r.db.First(&music, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &music, nil
}

func (r *MusicStorage) FindByName(name string) (*models.Music, error) {
	var music models.Music
	if err := r.db.Where("name = ?", name).First(&music).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &music, nil
}

func (r *MusicStorage) Update(music *models.Music) error {
	return r.db.Save(music).Error
}

func (r *MusicStorage) Delete(id uint) error {
	return r.db.Delete(&models.Music{}, id).Error
}
