package user

import (
	"net/http"

	"github.com/lishaonian/apps/response"
	"github.com/lishaonian/apps/service/user/api/internal/logic/user"
	"github.com/lishaonian/apps/service/user/api/internal/svc"
	"github.com/lishaonian/apps/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		response.Response(w, resp, err)
	}
}
