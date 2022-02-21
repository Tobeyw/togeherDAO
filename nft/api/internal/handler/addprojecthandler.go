package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"togetherdao/nft/api/internal/logic"
	"togetherdao/nft/api/internal/svc"
	"togetherdao/nft/api/internal/types"
)

func AddProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddProjectLogic(r.Context(), svcCtx)
		resp, err := l.AddProject(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
