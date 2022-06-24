package user

import (
	"context"

	"github.com/lishaonian/apps/service/user/api/internal/svc"
	"github.com/lishaonian/apps/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserUpdateReq) (resp bool, err error) {

	/* 	user := &model.UserInfo{
	   		Id:   req.Id,
	   		Name: req.Name,
	   	}
	   	err = nil
	   	if err != nil {
	   		return false, err
	   	} */
	return true, nil
}
