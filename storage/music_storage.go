package storage

import (
	"errors"

	"github.com/terloo/xiaochen/storage/models"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type musicStorage struct {
	db *gorm.DB
}

func newMusicStorage(db *gorm.DB) *musicStorage {
	return &musicStorage{db: db}
}

func (r *musicStorage) Create(music *models.Music) error {
	return r.db.Create(music).Error
}

func (r *musicStorage) Upsert(music *models.Music) error {
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "music_id"}, {Name: "source"}},
		DoUpdates: []clause.Assignment{{
			Column: clause.Column{Name: "downloaded"},
			Value:  music.Downloaded,
		}},
	}).Create(music).Error
}

func (r *musicStorage) FindByID(id uint) (*models.Music, error) {
	var music models.Music
	if err := r.db.First(&music, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &music, nil
}

func (r *musicStorage) FindByMusicIdAndSource(musicId string, source models.MusicSource) (*models.Music, error) {
	var music models.Music
	if err := r.db.Where("music_id = ? and source = ?", musicId, source).First(&music).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &music, nil
}

func (r *musicStorage) Update(music *models.Music) error {
	return r.db.Save(music).Error
}

func (r *musicStorage) Delete(id uint) error {
	return r.db.Delete(&models.Music{}, id).Error
}
