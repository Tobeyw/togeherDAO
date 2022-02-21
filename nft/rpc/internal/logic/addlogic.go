package logic

import (
	"context"
	"togetherdao/nft/model"

	"togetherdao/nft/rpc/internal/svc"
	"togetherdao/nft/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *project.ProjectReq) (*project.ProjectResp, error) {
	// todo: add your logic here and delete this line
    _,err := l.svcCtx.Model.Insert(&model.Project{
		ProjectId: in.ProjectId,
		ProjectToken: in.ProjectToken,
	})
	if err != nil {
		return nil, err
	}

	return &project.ProjectResp{Ok: "ok"}, nil
}
