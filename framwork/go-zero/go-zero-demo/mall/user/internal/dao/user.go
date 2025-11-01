package dao

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/database"
	"user/internal/model"
)

var cacheUserIdPrefix = "cache:user:id:"

type UserDao struct {
	*database.DBConn
}

func (d *UserDao) FindById(ctx context.Context, id int64) (user *model.User, err error) {
	user = &model.User{}
	query := fmt.Sprintf("select * from %s where id=?", user.TableName())
	logx.Info("[UserDao] [FindById]", query, id)
	userIdKey := fmt.Sprintf("%s:%d", cacheUserIdPrefix, id)
	logx.Info("[UserDao] [userIdKey]", userIdKey)
	err = d.ConnCache.QueryRowCtx(ctx, user, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	return
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		conn,
	}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	logx.Info("UserDao Save usesr is ", user)
	sql := fmt.Sprintf("insert into %s(name, gender) values(?,?)", user.TableName())
	result, err := d.Conn.ExecCtx(ctx, sql, user.Name, user.Gender)
	if err != nil {
		logx.Error(err.Error())
		return nil
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}
