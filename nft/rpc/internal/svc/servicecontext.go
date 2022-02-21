package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"togetherdao/nft/model"
	"togetherdao/nft/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model model.ProjectModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewProjectModel(sqlx.NewMysql(c.DataSource)),
	}
}
