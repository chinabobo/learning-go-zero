package logic

import (
	"context"
	"fmt"
	"github.com/chinabobo/learning-go-zero/internal/svc"
	"github.com/chinabobo/learning-go-zero/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type DemoPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoPostLogic {
	return &DemoPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoPostLogic) DemoPost(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	resp.Message = fmt.Sprintf("give %s", req.Name)
	return
}
