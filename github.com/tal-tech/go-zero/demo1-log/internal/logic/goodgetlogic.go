package logic

import (
	"context"

	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/svc"
	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GoodGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GoodGetLogic {
	return GoodGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodGetLogic) GoodGet(req types.GeneralRequest) (*types.GeneralResponse, error) {
	// todo: add your logic here and delete this line

	return &types.GeneralResponse{}, nil
}
