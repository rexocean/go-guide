// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"userapi/internal/config"
	"userapi/internal/errorx"
	"userapi/internal/handler"
	"userapi/internal/svc"
	"userapi/zapx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userapi-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	// 设置全局的错误处理器
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (httpStatus int, errorData interface{}) {
		switch e := err.(type) {
		case *errorx.BizErr:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	// 设置全局的错误处理器
	httpx.SetErrorHandler(func(err error) (httpStatus int, errorData interface{}) {
		switch e := err.(type) {
		case *errorx.BizErr:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	// 设置log的writer
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
	fmt.Sprintf("starting server at %s:%d", c.Host, c.Port)
	server.Start()
}
