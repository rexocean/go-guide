// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"rpc-common/types/user"
	"time"
	"userapi/internal/errorx"
	"userapi/internal/svc"
	"userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	userRequest := &user.UserRequest{
		Name:   req.Name,
		Gender: req.Gender,
	}
	logx.Info("userRequest:", userRequest)
	userResponse, err := l.svcCtx.UserRpc.SaveUser(ctx, userRequest)
	logx.Info("userResponse:", userResponse)
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: "success",
		Data:    userResponse,
	}, nil
}

func (l *UserLogic) GetUser(t *types.IdRequest) (resp *types.Response, err error) {
	// 认证通过后，从token中获取用户id
	userId := l.ctx.Value("userId")
	logx.Infof("获取到的token内容:%s\n", userId)

	if t.Id == "1" {
		//return nil, errors.New("参数不正确")
		return nil, errorx.PramsError
	}

	userRe, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{Id: t.Id})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: "success",
		Data:    userRe,
	}, nil
}

func (l *UserLogic) Login(t *types.LoginRequest) (token string, err error) {
	logx.Infof("invoking login method...")
	userId := 100
	auth := l.svcCtx.Config.Auth
	return l.getToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, int64(userId))
}

func (l *UserLogic) getToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
