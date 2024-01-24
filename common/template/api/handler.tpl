package {{.PkgName}}

import (
	"net/http"
    "lightleap-c/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"lightleap-c/common/xerr"
    "lightleap-c/common/xerr/code"
    "lightleap-c/common/xerr/msg"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpResult(r, w, nil, xerr.Wrap(err, code.CODE_PARAMS_INVALID, msg.GetMsg(code.CODE_PARAMS_INVALID)))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		resp, err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		result.HttpResult(r, w, resp, err)
	}
}
