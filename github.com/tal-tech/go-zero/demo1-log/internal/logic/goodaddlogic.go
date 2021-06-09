package logic

import (
	"context"

	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/svc"
	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GoodAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) GoodAddLogic {
	return GoodAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodAddLogic) GoodAdd(req types.GeneralRequest) (*types.GeneralResponse, error) {
	// todo: add your logic here and delete this line

	return &types.GeneralResponse{}, nil
}
