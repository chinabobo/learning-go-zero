package logic

import (
	"context"

	"github.com/chinabobo/learning-go-zero/rpc/internal/svc"
	"github.com/chinabobo/learning-go-zero/rpc/train_proxy"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTaskLogic) CreateTask(in *train_proxy.CreateTaskReq) (*train_proxy.CreateTaskResp, error) {
	// todo: add your logic here and delete this line

	return &train_proxy.CreateTaskResp{}, nil
}
