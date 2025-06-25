package netease

import (
	"context"
	"encoding/json"
	"log"
	neturl "net/url"
	"strconv"
	"strings"

	"github.com/terloo/xiaochen/client"

	"github.com/pkg/errors"
)

func CheckLoginStatus(ctx context.Context) (bool, error) {
	b, err := client.HttpGet(ctx, baseUrl.Get()+"/login/status", nil, nil)
	if err != nil {
		return false, err
	}
	log.Println(string(b))
	return true, nil
}

func GetUserPlayList(ctx context.Context, uid int) ([]PlayList, error) {
	b, err := client.HttpGet(ctx, baseUrl.Get()+"/user/playlist", nil, neturl.Values{
		"uid": []string{strconv.Itoa(uid)},
	})
	if err != nil {
		return nil, err
	}
	var result PlayListResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if result.Code != 200 {
		return nil, errors.Errorf("result code is not 200, resp is %s", string(b))
	}
	return result.PlayList, nil
}

func GetSongDetails(ctx context.Context, songIds ...int) ([]Music, error) {
	if len(songIds) < 1 {
		return []Music{}, nil
	}

	var songIdsStr []string
	for _, songId := range songIds {
		songIdsStr = append(songIdsStr, strconv.Itoa(songId))
	}

	b, err := client.HttpGet(ctx, baseUrl.Get()+"/song/detail", nil, neturl.Values{
		"ids": []string{strings.Join(songIdsStr, ",")},
	})
	if err != nil {
		return nil, err
	}

	var result SongDetailResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if result.Code != 200 {
		return nil, errors.Errorf("result code is not 200, resp is %s", string(b))
	}

	var musics []Music
	for _, song := range result.Songs {
		var arNames []string
		for _, e := range song.Ar {
			arNames = append(arNames, e.Name)
		}
		musics = append(musics, Music{
			IdName: IdName{
				Id:   song.Id,
				Name: song.Name,
			},
			Artist: arNames,
			Album:  song.Al.Name,
		})
	}
	return musics, nil
}

func GetUserPlayListTrackIds(ctx context.Context, playListId int) ([]int, error) {
	b, err := client.HttpGet(ctx, baseUrl.Get()+"/playlist/detail", nil, neturl.Values{
		"id": []string{strconv.Itoa(playListId)},
	})
	if err != nil {
		return nil, err
	}
	var result PlayListDetailsResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if result.Code != 200 {
		return nil, errors.Errorf("result code is not 200, resp is %s", string(b))
	}

	var trackIds []int
	for _, track := range result.PlayList.TrackIds {
		trackIds = append(trackIds, track.Id)
	}

	return trackIds, nil
}
