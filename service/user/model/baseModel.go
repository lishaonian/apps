package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type BaseModel struct {
	conn sqlx.SqlConn
}

func NewBaseModel(conn sqlx.SqlConn) *BaseModel {
	return &BaseModel{
		conn: conn,
	}
}

func (b *BaseModel) FindList(ctx context.Context, field string, table string, data interface{}) error {

	query := fmt.Sprintf("select %s from `%s` limit %d,%d", field, table, 0, 10)
	err := b.conn.QueryRowsCtx(ctx, data, query)
	if err != nil {
		return err
	}
	return nil
}
