package database

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DBConn struct {
	Conn      sqlx.SqlConn
	ConnCache sqlc.CachedConn
}

func Connect(datasource string, conf cache.CacheConf) *DBConn {
	sqlConn := sqlx.NewMysql(datasource)
	d := &DBConn{
		Conn: sqlConn,
	}
	if conf != nil {
		cacheConn := sqlc.NewConn(sqlConn, conf)
		d.ConnCache = cacheConn
	}
	return d
}
