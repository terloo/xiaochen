package models

import (
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	MusicId         string `gorm:"index"`
	Name            string `gorm:"index"`
	Artist          string
	Album           string               `gorm:"index"`
	PicId           string               `gorm:"index"`
	LyricId         string               `gorm:"index"`
	Source          MusicSource          `gorm:"index,type:varchar(30)"`
	DownloadChannel MusicDownloadChannel `gorm:"index,type:varchar(30)"`
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
	MusicDownloadChannelWx      MusicDownloadChannel = "wx"
	MusicDownloadChannelManual      MusicDownloadChannel = "manual"
)
