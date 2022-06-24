package logic

import (
	"context"
	"fmt"

	"github.com/lishaonian/apps/common/xerr"
	"github.com/lishaonian/apps/service/user/model"
	"google.golang.org/grpc/status"

	"github.com/lishaonian/apps/service/user/rpc/internal/svc"
	"github.com/lishaonian/apps/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.IdReq) (*pb.UserRep, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	var userlist []model.UserInfo
	fmt.Println(1111111111)

	err = l.svcCtx.BaseModel.FindList(l.ctx, model.UserInfoRows, l.svcCtx.UserModel.TableName(), &userlist)

	fmt.Println(userlist)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("用户不存在")
			//return nil, xerr.NewErrMsg("用户不存在")
			return nil, status.Error(100, "用户不存在")

		}
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	if user.Id == 0 {
		return nil, status.Error(100, "用户不存在")
	}

	return &pb.UserRep{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
	}, nil
}
