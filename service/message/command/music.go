package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/terloo/xiaochen/config"
	"github.com/terloo/xiaochen/storage/models"
	"github.com/terloo/xiaochen/thirdparty/gd"
	"github.com/terloo/xiaochen/thirdparty/wxbot"
)

var savePath = config.NewLoader("thirdparty.gd.savePath")

type Music struct {
}

var _ Handler = (*Music)(nil)

func (m *Music) CommandName() string {
	return "下歌"
}

func (m *Music) Exec(ctx context.Context, caller string, args []string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if len(args) < 1 {
		return errors.Errorf("not valid args: %v", args)
	}

	musics, err := gd.SearchMusic(ctx, string(models.MusicSourceKuwo), strings.Join(args, " "), 1, 1)
	if err != nil {
		return err
	}
	if len(musics) < 1 {
		return errors.Errorf("未找到歌曲 %v", args)
	}

	for _, music := range musics {
		_ = wxbot.SendMsg(ctx, fmt.Sprintf("正在下载 [%s - %s]", music.Artist[0], music.Name), caller)
		musicName, md5, err := gd.PersistentMusic(ctx, savePath.Get(), *music)
		if err != nil {
			return errors.Wrapf(err, "persisten music [%s - %s] error\n", music.Artist[0], music.Name)
		}
		_ = wxbot.SendMsg(ctx, fmt.Sprintf("已成功下载 [%s]，md5: %s", musicName, md5), caller)
	}
	return nil
}
