package handler

import (
	"net/http"

	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/logic"
	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/svc"
	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func DetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GeneralRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDetailLogic(r.Context(), ctx)
		resp, err := l.Detail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
