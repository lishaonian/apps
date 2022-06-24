package svc

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jinzhu/gorm"
	"github.com/lishaonian/apps/service/user/model"
	"github.com/lishaonian/apps/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserInfoModel
	BaseModel *model.BaseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	db, err := gorm.Open("mysql", c.DB.DataSource)
	fmt.Printf("%v", err)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserInfoModel(db),
		BaseModel: model.NewBaseModel(sqlConn),
	}
}
