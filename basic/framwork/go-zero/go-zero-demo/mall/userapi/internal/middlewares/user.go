package middlewares

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (*UserMiddleware) LoginAndReg(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("[LoginAndReg] invoke reg and login before do sth, req: ", r)
		next(w, r)
		logx.Info("[LoginAndReg] invoke reg and login after do sth,, req: ", r)
	}
}

func (*UserMiddleware) Global(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("global all  before do sth, req: ", r)
		next(w, r)
		logx.Info("global after after do sth,, req: ", r)
	}
}
