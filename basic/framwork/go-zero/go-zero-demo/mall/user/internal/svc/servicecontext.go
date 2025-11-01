package svc

import (
	"user/database"
	"user/internal/config"
	"user/internal/dao"
	"user/internal/repo"
)

// serviceContext logic 依赖的资源池
type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	connect := database.Connect(c.Mysql.DataSource, c.CacheRedis)
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(connect),
	}
}
