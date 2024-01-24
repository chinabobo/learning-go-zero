package {{.packageName}}

import (
	"context"
    "lightleap-c/common/xlog"
	{{.imports}}

	"github.com/zeromicro/go-zero/core/logx"
)

type {{.logicName}} struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func New{{.logicName}}(ctx context.Context,svcCtx *svc.ServiceContext) *{{.logicName}} {
	return &{{.logicName}}{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: xlog.FromContext(ctx),
	}
}
{{.functions}}
