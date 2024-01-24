package handler

import (
	"github.com/chinabobo/learning-go-zero/internal/logic"
	"github.com/chinabobo/learning-go-zero/internal/svc"
	"github.com/chinabobo/learning-go-zero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func DemoPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewDemoPostLogic(r.Context(), svcCtx)
		resp, err := l.DemoPost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
