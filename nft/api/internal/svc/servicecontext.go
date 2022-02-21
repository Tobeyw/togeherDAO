package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"togetherdao/nft/api/internal/config"
	"togetherdao/nft/rpc/nft"
)

type ServiceContext struct {
	Config config.Config
	NFT  nft.Nft
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		NFT: nft.NewNft(zrpc.MustNewClient(c.NFT)),
	}
}
