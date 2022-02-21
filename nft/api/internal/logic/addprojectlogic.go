package logic

import (
	"context"

	"togetherdao/nft/api/internal/svc"
	"togetherdao/nft/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddProjectLogic {
	return AddProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProjectLogic) AddProject(req types.ProjectReq) (resp *types.ProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
