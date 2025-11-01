package handler

import (
	"userapi/internal/svc"
)

type UserHandler struct {
	serverCtx *svc.ServiceContext
}

func NewUserHandler(serverCtx *svc.ServiceContext) *UserHandler {
	return &UserHandler{
		serverCtx: serverCtx,
	}
}
