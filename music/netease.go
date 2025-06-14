package music

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/robfig/cron/v3"
	"github.com/terloo/xiaochen/config"
	"github.com/terloo/xiaochen/storage"
	"github.com/terloo/xiaochen/storage/models"
	"github.com/terloo/xiaochen/thirdparty/gd"
	"github.com/terloo/xiaochen/thirdparty/netease"
)

var uid = config.NewLoader("thirdparty.netease.uid")
var savePath = config.NewLoader("thirdparty.gd.savePath")

func StartPeriodNetease(ctx context.Context) {
	printfLogger := cron.VerbosePrintfLogger(log.New(log.Writer(), "[period_netease]  ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(printfLogger),
		cron.WithChain(cron.Recover(printfLogger)),
	)

	c.AddFunc("0 */5 * * * *", func() {
		PersistentLikeMusic(ctx)
	})

	c.Start()
	<-ctx.Done()
	stop := c.Stop()
	<-stop.Done()
}

func PersistentLikeMusic(ctx context.Context) {
	// 获取网易云like列表
	playList, err := netease.GetUserPlayList(ctx, uid.GetInt())
	if err != nil {
		log.Printf("get netease playList error: %+v\n", err)
	}

	for _, pl := range playList {
		if pl.SpecialType != 5 {
			// 仅获取“我喜欢”
			continue
		}
		trackIds, err := netease.GetUserPlayListTrackIds(ctx, pl.Id)
		if err != nil {
			log.Printf("get netease playList track id error: %+v\n", err)
		}
		if len(trackIds) < 1 {
			log.Printf("no music in playLsit %s\n", pl.Name)
			continue
		}
		chunksSize := 20
		for i := 0; i < len(trackIds); i += chunksSize {
			end := i + chunksSize
			if end > len(trackIds) {
				end = len(trackIds)
			}
			trackIdChunk := trackIds[i:end]
			if len(trackIdChunk) == 0 {
				break
			}
			log.Printf("seeking trackIdChunk %v\n", trackIdChunk)
			musics, err := netease.GetSongDetails(ctx, trackIdChunk...)
			if err != nil {
				log.Printf("get netease song detail error: %+v\n", err)
				continue
			}
			for _, music := range musics {
				record, err := storage.MusicRepo.FindByMusicIdAndSource(strconv.Itoa(music.Id), models.MusicSourceNetease)
				if err != nil {
					log.Printf("get db data error: %+v\n", err)
				}
				if record != nil && record.Downloaded {
					log.Printf("music [%v %s] has downloaded\n", music.Artist, music.Name)
					continue
				}

				log.Printf("seeking music [%v %s] gd info", music.Artist, music.Name)
				gdMusic, err := gd.GetMusic(ctx, strconv.Itoa(music.Id), models.MusicSourceNetease.String(), music.Name, music.Artist[0])
				if err != nil {
					log.Printf("get dg music info error: %+v\n", err)
					continue
				}
				err = gd.PersistentMusic(ctx, savePath.Get(), *gdMusic)
				if err != nil {
					log.Printf("persistent dg music file error: %+v\n", err)
					continue
				}
				// 保存数据库记录
				err = storage.MusicRepo.Upsert(&models.Music{
					MusicId:         string(gdMusic.Id),
					Name:            gdMusic.Name,
					Artist:          strings.Join(gdMusic.Artist, ","),
					Album:           gdMusic.Album,
					PicId:           string(gdMusic.PicId),
					LyricId:         string(gdMusic.LyricId),
					Source:          models.MusicSource(gdMusic.Source),
					DownloadChannel: models.MusicDownloadChannelNeteaseLike,
					Downloaded:      true,
				})
				if err != nil {
					log.Printf("persistent dg music db data error: %+v\n", err)
					continue
				}
			}
		}
	}
}
