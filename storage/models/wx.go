package models

import "gorm.io/gorm"

type WxFamily struct {
	gorm.Model
	Wxid     string `gorm:"index"`
	NickName string `gorm:"index"`
	Birthday string `gorm:"index"`
	Lunar    bool   `gorm:"index"`
}

type WxChat struct {
	gorm.Model
	Wxid string `gorm:"index"`
	Name string `gorm:"index"`
}
