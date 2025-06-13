package gd

import (
	"encoding/json"
	"strconv"
)

type Music struct {
	Id      StringOrInt
	Name    string
	Artist  []string
	Album   string
	PicId   StringOrInt
	LyricId StringOrInt
	Source  string
}

type MusicURL struct {
	Url string
	// 音质
	Br float64
	// 大小
	Size int
}

type MusicPic struct {
	Url string
}

type MusicLyric struct {
	Lyric string
}

type StringOrInt string

func (s *StringOrInt) UnmarshalJSON(data []byte) error {
	var intValue int
	if err := json.Unmarshal(data, &intValue); err == nil {
		*s = StringOrInt(strconv.Itoa(intValue))
		return nil
	}

	var stringValue string
	if err := json.Unmarshal(data, &stringValue); err != nil {
		return err
	}

	*s = StringOrInt(stringValue)
	return nil
}
