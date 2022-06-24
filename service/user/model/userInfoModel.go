package model

import (
	"github.com/jinzhu/gorm"
)

var _ UserInfoModel = (*customUserInfoModel)(nil)

type (
	// UserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInfoModel.
	UserInfoModel interface {
		userInfoModel
	}

	customUserInfoModel struct {
		*defaultUserInfoModel
	}
)

// NewUserInfoModel returns a model for the database table.
func NewUserInfoModel(conn *gorm.DB) UserInfoModel {
	return &customUserInfoModel{
		defaultUserInfoModel: newUserInfoModel(conn),
	}
}
