package music

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"github.com/terloo/xiaochen/config"
	"github.com/terloo/xiaochen/storage"
	"github.com/terloo/xiaochen/storage/models"
	"github.com/terloo/xiaochen/thirdparty/gd"
	"github.com/terloo/xiaochen/thirdparty/netease"
)

var uid = config.NewLoader("thirdparty.netease.uid")

func StartPeriodNetease(ctx context.Context) {
	printfLogger := cron.VerbosePrintfLogger(log.New(log.Writer(), "[period_netease]  ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(printfLogger),
		cron.WithChain(cron.Recover(printfLogger), cron.SkipIfStillRunning(printfLogger)),
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
	if err := ctx.Err(); err != nil {
		return
	}

	// 获取网易云like列表
	playList, err := netease.GetUserPlayList(ctx, uid.GetInt())
	if err != nil {
		log.Printf("get netease playList error: %+v\n", err)
	}

	for _, pl := range playList {
		if err := ctx.Err(); err != nil {
			return
		}
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
			if err := ctx.Err(); err != nil {
				return
			}

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
				if err := ctx.Err(); err != nil {
					return
				}

				record, err := storage.MusicRepo.FindByMusicIdAndSource(strconv.Itoa(music.Id), models.MusicSourceNetease)
				if err != nil {
					log.Printf("get db data error: %+v\n", err)
				}
				if record != nil && (record.Downloaded || record.DownloadRetryCount >= 10) {
					// log.Printf("music [%v %s] has downloaded or retch max retry count\n", music.Artist, music.Name)
					continue
				}

				// 下载歌曲
				gdMusic, fileName, md5, err := persistentLikeMusicInternal(ctx, music)
				if err != nil {
					log.Printf("persistent music [%v %s] error: %+v\n", music.Artist, music.Name, err)
				}

				var musicRecord *models.Music
				if err == nil {
					musicRecord = &models.Music{
						MusicId:             string(gdMusic.Id),
						Name:                gdMusic.Name,
						Artist:              strings.Join(gdMusic.Artist, ","),
						Album:               gdMusic.Album,
						PicId:               string(gdMusic.PicId),
						LyricId:             string(gdMusic.LyricId),
						Source:              models.MusicSource(gdMusic.Source),
						FileName:            fileName,
						MD5:                 md5,
						DownloadChannel:     models.MusicDownloadChannelNeteaseLike,
						Downloaded:          true,
						DownloadRetryCount:  0,
						DownloadErrorReason: "",
					}
				} else {
					retryCount := 0
					if record != nil {
						retryCount = record.DownloadRetryCount + 1
					}
					musicRecord = &models.Music{
						MusicId:             strconv.Itoa(music.Id),
						Name:                music.Name,
						Artist:              strings.Join(music.Artist, ","),
						Album:               music.Album,
						PicId:               "",
						LyricId:             "",
						Source:              models.MusicSourceNetease,
						FileName:            strings.Replace(uuid.New().String(), "-", "", -1),
						MD5:                 "",
						DownloadChannel:     models.MusicDownloadChannelNeteaseLike,
						Downloaded:          false,
						DownloadRetryCount:  retryCount,
						DownloadErrorReason: err.Error(),
					}
				}

				// 保存数据库记录
				err = storage.MusicRepo.Upsert(musicRecord)
				if err != nil {
					log.Printf("persistent dg music db data error: %+v\n", err)
					continue
				}
			}
		}
	}
}

func persistentLikeMusicInternal(ctx context.Context, music netease.Music) (*gd.Music, string, string, error) {
	if err := ctx.Err(); err != nil {
		return nil, "", "", err
	}

	log.Printf("seeking music [%v %s] gd info", music.Artist, music.Name)
	gdMusic, err := gd.GetMusic(ctx, strconv.Itoa(music.Id), models.MusicSourceNetease.String(), music.Name, music.Artist, "")
	if err != nil {
		log.Printf("get dg music info error: %+v\n", err)
		return nil, "", "", err
	}
	fileName, md5, err := gd.PersistentMusic(ctx, *gdMusic)
	if err != nil {
		log.Printf("persistent dg music file error: %+v\n", err)
		return nil, "", "", err
	}
	return gdMusic, fileName, md5, nil
}
