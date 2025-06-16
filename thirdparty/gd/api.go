package gd

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	neturl "net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/terloo/xiaochen/client"
	"github.com/terloo/xiaochen/config"
)

var tickerDuration = config.NewLoader("thirdparty.gd.httpDuration")

var httpGetTicker = time.NewTicker(time.Duration(tickerDuration.GetInt()) * time.Second)

func tickerHttpGet(ctx context.Context, url string, header http.Header, param neturl.Values) ([]byte, error) {
	<-httpGetTicker.C
	return client.HttpGet(ctx, url, header, param)
}

// SearchMusic 搜索音乐
func SearchMusic(ctx context.Context, source string, search string, count int, page int) ([]*Music, error) {
	b, err := tickerHttpGet(ctx, gdURL.Get(), http.Header{}, neturl.Values{
		"types":  []string{"search"},
		"source": []string{source},
		"name":   []string{search},
		"count":  []string{strconv.Itoa(count)},
		"pages":  []string{strconv.Itoa(page)},
	})
	if err != nil {
		return nil, err
	}
	var result []*Music
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
	// 遍历20页
	for i := 0; i < 20; i++ {
		musics, err := SearchMusic(ctx, source, search, 20, i+1)
		if err != nil {
			return nil, err
		}
		if len(musics) == 1 {
			break
		}
		for _, music := range musics {
			if id == string(music.Id) {
				return music, nil
			}
		}
	}
	return nil, errors.Errorf("music not found, id: %s, source: %s, name: %s, artist %s", id, source, name, artist)
}

// PersistentMusic 下载并整理歌词元数据，持久化到指定目录中
func PersistentMusic(ctx context.Context, savePath string, music Music) (string, string, error) {
	// 目录判断
	info, err := os.Stat(savePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", "", errors.Wrap(err, "读取目录失败")
		}
		err := os.MkdirAll(savePath, os.ModeDir|0755)
		if err != nil {
			return "", "", errors.Wrap(err, "创建目录失败")
		}
	}
	info, err = os.Stat(savePath)
	if err != nil {
		return "", "", errors.Wrap(err, "读取目录失败")
	}
	if !info.IsDir() {
		return "", "", errors.Errorf("%s 不是目录", savePath)
	}

	_, err = os.ReadDir(savePath)
	if err != nil {
		return "", "", errors.Wrap(err, "读取目录失败")
	}

	// 下载歌曲
	reader, extension, err := DownloadMusic(ctx, music)
	if err != nil {
		return "", "", err
	}

	// 修改歌曲元数据
	musicReader, err := modifyMusicMeta(ctx, reader, music, extension)
	if err != nil {
		return "", "", err
	}

	// 保存修改后的歌曲
	musicName := fmt.Sprintf("%s - %s.%s", music.Artist[0], music.Name, extension)
	musicPath := filepath.Join(savePath, musicName)
	file, err := os.OpenFile(musicPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", "", errors.Wrapf(err, "打开[%s]失败", musicPath)
	}
	defer file.Close()

	_, err = io.Copy(file, musicReader)
	if err != nil {
		return "", "", errors.Wrapf(err, "保存最终文件[%s]失败", musicPath)
	}
	fileMD5, err := calculateFileMD5(musicPath)
	if err != nil {
		log.Printf("calculate file MD5 error: %+v", err)
		fileMD5 = ""
	}
	return musicName, fileMD5, nil
}

// DownloadMusicPic 下载音乐封面
func DownloadMusicPic(ctx context.Context, music Music) (io.Reader, string, error) {
	b, err := tickerHttpGet(ctx, gdURL.Get(), http.Header{}, neturl.Values{
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
		return nil, "", errors.Errorf("获取歌曲封面链接失败，resp: %s", string(b))
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
	picFormat, err := guessPicFormat(buf.Bytes())
	if err != nil {
		return nil, "", errors.WithStack(err)
	}
	return &buf, picFormat, nil

}

// DownloadMusicLyric 下载音乐歌词
func DownloadMusicLyric(ctx context.Context, music Music) (string, error) {
	b, err := tickerHttpGet(ctx, gdURL.Get(), http.Header{}, neturl.Values{
		"types":  []string{"lyric"},
		"source": []string{music.Source},
		"id":     []string{string(music.LyricId)},
	})
	if err != nil {
		return "", err
	}

	var musicLyric MusicLyric
	err = json.Unmarshal(b, &musicLyric)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if len(musicLyric.Lyric) == 0 {
		return "", errors.Errorf("获取歌曲歌词失败，resp: %s", string(b))
	}
	return musicLyric.Lyric, nil
}

// DownloadMusic 下载音乐
func DownloadMusic(ctx context.Context, music Music) (io.Reader, string, error) {
	b, err := tickerHttpGet(ctx, gdURL.Get(), http.Header{}, neturl.Values{
		"types":  []string{"url"},
		"source": []string{music.Source},
		"id":     []string{string(music.Id)},
	})
	if err != nil {
		return nil, "", err
	}

	var musicURL MusicURL
	err = json.Unmarshal(b, &musicURL)
	if err != nil {
		return nil, "", errors.WithStack(err)
	}
	if len(musicURL.Url) == 0 {
		return nil, "", errors.Errorf("获取歌曲下载链接失败，resp: %s", string(b))
	}

	extension := getExtension(musicURL.Url)
	if extension != "flac" && extension != "mp3" {
		return nil, "", errors.Errorf("后缀不是flac或mp3，resp: %s", string(b))
	}

	// 下载歌曲
	resp, err := http.Get(musicURL.Url)
	if err != nil {
		return nil, "", errors.WithMessage(err, "下载歌曲失败")
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, "", errors.WithStack(err)
	}
	return &buf, extension, nil
}

func getExtension(filename string) string {
	lastDot := strings.LastIndex(filename, ".")
	if lastDot == -1 || lastDot == len(filename)-1 {
		return ""
	}
	return filename[lastDot+1:]
}

func guessPicFormat(data []byte) (string, error) {
	if len(data) < 3 {
		return "", errors.New("no valid pic data")
	}
	if data[0] == 0xFF && data[1] == 0xD8 {
		return "jpeg", nil
	}
	if data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E {
		return "png", nil
	}
	return "", errors.Errorf("not valid pic format, header: %v", data[:3])
}

func calculateFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer file.Close()

	hash := md5.New()

	buf := make([]byte, 65536) // 64KB缓冲区
	for {
		n, err := file.Read(buf)
		if n > 0 {
			hash.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", errors.WithStack(err)
		}
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func modifyMusicMeta(ctx context.Context, reader io.Reader, music Music, extension string) (io.Reader, error) {
	var handler MusicMetaDataHandler
	var err error
	if extension == "flac" {
		handler, err = NewFlacMusicMetaDataHandler(music, reader)
		defer handler.Close()
		if err != nil {
			return nil, err
		}
	} else if extension == "mp3" {
		handler, err = NewMp3MusicMetaDataHandler(music, reader)
		defer handler.Close()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.Errorf("unsuppor file format: %s", extension)
	}
	err = handler.AddCommentMetaData()
	if err != nil {
		return nil, err
	}
	// 添加歌词
	lyric, err := DownloadMusicLyric(ctx, music)
	if err != nil {
		return nil, err
	}
	err = handler.AddLyric(lyric)
	if err != nil {
		return nil, err
	}
	// 添加图片
	picReader, picExtension, err := DownloadMusicPic(ctx, music)
	if err != nil {
		return nil, err
	}
	if picExtension != "jpeg" && picExtension != "png" {
		return nil, errors.Errorf("not valid pic format: %s\n", picExtension)
	}
	err = handler.AddPic(picReader, picExtension)
	if err != nil {
		return nil, err
	}
	musicReader, err := handler.toReader()
	if err != nil {
		return nil, err
	}
	return musicReader, nil
}
