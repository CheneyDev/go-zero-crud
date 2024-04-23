package logic

import (
	"context"
	"demo/internal/svc"
	"demo/internal/types"
	"demo/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line

	user := &model.Users{
		Name:  req.Name,
		Email: req.Email,
	}
	res, err := model.UsersModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{Id: res.Id}, nil

}
