package logic

import (
	"context"
	"rpc-common/types/user"
	"strconv"
	"user/internal/model"

	"user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	id, _ := strconv.ParseInt(in.Id, 10, 64)
	data, err := l.svcCtx.UserRepo.FindById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:     in.GetId(),
		Name:   data.Name,
		Gender: data.Gender,
	}, nil
}

func (l *UserLogic) SaveUser(in *user.UserRequest) (*user.UserResponse, error) {
	data := &model.User{
		Name:   in.Name,
		Gender: in.Gender,
	}
	err := l.svcCtx.UserRepo.Save(context.Background(), data)
	if err != nil {
		return nil, err
	}
	//userId := fmt.Sprintf("%d", data.Id)
	return &user.UserResponse{
		Id:     strconv.FormatInt(data.Id, 10),
		Name:   data.Name,
		Gender: data.Gender,
	}, err
}
