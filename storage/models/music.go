package models

import (
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	MusicId         string               `gorm:"index;type:varchar(100);uniqueIndex:uniq_id_source;not null"`
	Name            string               `gorm:"index;type:varchar(100)"`
	Artist          string               `gorm:"index;type:varchar(300)"`
	Album           string               `gorm:"index;type:varchar(100)"`
	PicId           string               `gorm:"index;type:varchar(100)"`
	LyricId         string               `gorm:"index;type:varchar(100)"`
	Source          MusicSource          `gorm:"index;type:varchar(30);uniqueIndex:uniq_id_source;not null"`
	DownloadChannel MusicDownloadChannel `gorm:"index;type:varchar(30)"`
	Downloaded      bool                 `gorm:"index;default:0"`
}

type MusicSource string

func (s MusicSource) String() string {
	return string(s)
}

const (
	MusicSourceKuwo    MusicSource = "kuwo"
	MusicSourceNetease MusicSource = "netease"
)

type MusicDownloadChannel string

func (s MusicDownloadChannel) String() string {
	return string(s)
}

const (
	MusicDownloadChannelNeteaseLike MusicDownloadChannel = "netease_like"
	MusicDownloadChannelWx          MusicDownloadChannel = "wx"
	MusicDownloadChannelManual      MusicDownloadChannel = "manual"
)
