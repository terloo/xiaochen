package gd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/terloo/xiaochen/client"

	"github.com/go-flac/flacpicture/v2"
	"github.com/go-flac/flacvorbis/v2"
	goflac "github.com/go-flac/go-flac/v2"
)

// SearchMusic 搜索音乐
func SearchMusic(ctx context.Context, source string, search string, count int, page int) ([]Music, error) {
	b, err := client.HttpGet(ctx, gdURL, http.Header{}, url.Values{
		"types":  []string{"search"},
		"source": []string{source},
		"name":   []string{search},
		"count":  []string{strconv.Itoa(count)},
		"pages":  []string{strconv.Itoa(page)},
	})
	if err != nil {
		return nil, err
	}
	var result []Music
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

// GetMusic 通过id，以及歌名、作者、获取对应的music对象
func GetMusic(ctx context.Context, id string, source string, name string, artist string) (*Music, error) {
	if len(name) == 0 && len(artist) == 0 {
		return nil, errors.New("no enough search info")
	}
	search := fmt.Sprintf("%s %s", name, artist)
	for i := 0; i < 20; i++ {
		musics, err := SearchMusic(ctx, source, search, 20, i)
		if err != nil {
			return nil, err
		}
		for _, music := range musics {
			if id == string(music.Id) {
				return &music, nil
			}
		}
	}
	return nil, errors.Errorf("music not found, id: %s, source: %s", id, source)
}

// PersistentMusic 下载并整理歌词元数据，持久化到指定目录中
func PersistentMusic(ctx context.Context, savePath string, music Music) error {
	// 目录判断
	info, err := os.Stat(savePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "读取目录失败")
		}
		err := os.MkdirAll(savePath, os.ModeDir|0755)
		if err != nil {
			return errors.Wrap(err, "创建目录失败")
		}
	}
	info, err = os.Stat(savePath)
	if err != nil {
		return errors.Wrap(err, "读取目录失败")
	}
	if !info.IsDir() {
		return errors.Errorf("%s 不是目录", savePath)
	}

	_, err = os.ReadDir(savePath)
	if err != nil {
		return errors.Wrap(err, "读取目录失败")
	}

	// 下载歌曲
	reader, err := DownloadMusic(ctx, music)
	if err != nil {
		return err
	}

	// 修改歌曲元数据
	flacFile, err := modifyMusicMeta(ctx, reader, music)
	if err != nil {
		return err
	}
	defer flacFile.Close()

	// 保存修改后的歌曲
	musicName := fmt.Sprintf("%s - %s.%s", music.Artist[0], music.Name, "flac")
	musicPath := filepath.Join(savePath, musicName)
	file, err := os.OpenFile(musicPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrapf(err, "打开[%s]失败", musicPath)
	}
	defer file.Close()

	_, err = flacFile.WriteTo(file)
	if err != nil {
		return errors.Wrapf(err, "保存最终文件[%s]失败", musicPath)
	}
	return nil
}

// DownloadMusicPic 下载音乐封面
func DownloadMusicPic(ctx context.Context, music Music) (io.Reader, string, error) {
	b, err := client.HttpGet(ctx, gdURL, http.Header{}, url.Values{
		"types":  []string{"pic"},
		"source": []string{music.Source},
		"id":     []string{string(music.PicId)},
	})
	if err != nil {
		return nil, "", err
	}

	var musicPic MusicPic
	err = json.Unmarshal(b, &musicPic)
	if err != nil {
		return nil, "", errors.WithStack(err)
	}
	if len(musicPic.Url) == 0 {
		return nil, "", errors.New("获取歌曲封面链接失败")
	}

	resp, err := http.Get(musicPic.Url)
	if err != nil {
		return nil, "", errors.WithStack(err)
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, "", errors.WithStack(err)
	}
	return &buf, getExtension(musicPic.Url), nil

}

// DownloadMusicLyric 下载音乐歌词
func DownloadMusicLyric(ctx context.Context, music Music) (io.Reader, error) {
	b, err := client.HttpGet(ctx, gdURL, http.Header{}, url.Values{
		"types":  []string{"lyric"},
		"source": []string{music.Source},
		"id":     []string{string(music.LyricId)},
	})
	if err != nil {
		return nil, err
	}

	var musicLyric MusicLyric
	err = json.Unmarshal(b, &musicLyric)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(musicLyric.Lyric) == 0 {
		return nil, errors.New("获取歌曲歌词失败")
	}
	return strings.NewReader(musicLyric.Lyric), nil
}

// DownloadMusic 下载音乐
func DownloadMusic(ctx context.Context, music Music) (io.Reader, error) {
	b, err := client.HttpGet(ctx, gdURL, http.Header{}, url.Values{
		"types":  []string{"url"},
		"source": []string{music.Source},
		"id":     []string{string(music.Id)},
	})
	if err != nil {
		return nil, err
	}

	var musicURL MusicURL
	err = json.Unmarshal(b, &musicURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(musicURL.Url) == 0 {
		return nil, errors.New("获取歌曲下载链接失败")
	}

	extension := getExtension(musicURL.Url)
	if extension != "flac" {
		return nil, errors.New("后缀不是flac")
	}

	// 下载歌曲
	resp, err := http.Get(musicURL.Url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &buf, nil
}

func getExtension(filename string) string {
	lastDot := strings.LastIndex(filename, ".")
	if lastDot == -1 || lastDot == len(filename)-1 {
		return ""
	}
	return filename[lastDot+1:]
}

func modifyMusicMeta(ctx context.Context, reader io.Reader, music Music) (*goflac.File, error) {
	flacFile, err := goflac.ParseBytes(reader)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var vorbisComment *flacvorbis.MetaDataBlockVorbisComment
	var cmtIdx int
	for idx, block := range flacFile.Meta {
		if block.Type == goflac.VorbisComment {
			vorbisComment, err = flacvorbis.ParseFromMetaDataBlock(*block)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			cmtIdx = idx
		}
	}
	vorbisComment = flacvorbis.New()
	vorbisComment.Add(flacvorbis.FIELD_TITLE, music.Name)
	vorbisComment.Add(flacvorbis.FIELD_ALBUM, music.Album)
	for _, artist := range music.Artist {
		vorbisComment.Add(flacvorbis.FIELD_ARTIST, artist)
	}
	vorbisComment.Add("SOURCE", music.Source)
	vorbisComment.Add("MUSIC_ID", string(music.Id))

	// 添加歌词
	lyricReader, err := DownloadMusicLyric(ctx, music)
	if err != nil {
		return nil, err
	}
	lyric, err := io.ReadAll(lyricReader)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	vorbisComment.Add("LYRICS", string(lyric))
	picReader, picExtension, err := DownloadMusicPic(ctx, music)
	if err != nil {
		return nil, err
	}
	pic, err := io.ReadAll(picReader)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 创建图片元数据块
	if idx := strings.LastIndex(picExtension, "?"); idx != -1 {
		picExtension = picExtension[:idx]
	}
	if picExtension == "jpg" {
		picExtension = "jpeg"
	}
	if picExtension == "jpeg" || picExtension == "png" {

		picture, err := flacpicture.NewFromImageData(
			flacpicture.PictureTypeFrontCover,
			"Front cover",
			pic,
			"image/"+picExtension,
		)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		picturemeta := picture.Marshal()
		flacFile.Meta = append(flacFile.Meta, &picturemeta)
	}

	cmtsmeta := vorbisComment.Marshal()
	if cmtIdx > 0 {
		flacFile.Meta[cmtIdx] = &cmtsmeta
	} else {
		flacFile.Meta = append(flacFile.Meta, &cmtsmeta)
	}
	return flacFile, nil
}
