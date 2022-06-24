package user

import (
	"context"

	"github.com/lishaonian/apps/service/user/api/internal/svc"
	"github.com/lishaonian/apps/service/user/api/internal/types"
	"github.com/lishaonian/apps/service/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserResp, err error) {
	// 使用user rpc
	userinfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: req.Id,
	})
	if err != nil {
		logx.Error("测试的日志")
		return nil, err
	}
	return &types.UserResp{
		Id:   userinfo.Id,
		Name: userinfo.Name,
		Age:  userinfo.Age,
	}, nil
}
