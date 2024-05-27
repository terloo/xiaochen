package command

import (
	"context"
	"strings"

	"github.com/terloo/xiaochen/thirdparty/juhe"
	"github.com/terloo/xiaochen/wxbot"
)

type Zhougong struct {
}

var _ Handler = (*Zhougong)(nil)

func (z *Zhougong) CommandName() string {
	return "解梦"
}

func (z *Zhougong) Exec(ctx context.Context, caller string, args []string) error {
	zhouGongResult, err := juhe.GetZhouGong(ctx, strings.Join(args, " "))
	if err != nil {
		_ = wxbot.SendMsg(ctx, err.Error(), caller)
		return err
	}
	respMsg := strings.Join(zhouGongResult.List, "\n")
	_ = wxbot.SendMsg(ctx, respMsg, caller)
	return nil
}
