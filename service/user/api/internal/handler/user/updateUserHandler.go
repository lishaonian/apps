package user

import (
	"net/http"

	"github.com/lishaonian/apps/service/user/api/internal/logic/user"
	"github.com/lishaonian/apps/service/user/api/internal/svc"
	"github.com/lishaonian/apps/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUpdateUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
