package service

import (
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"

	"github.com/google/wire"
)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.GreeterUsecase
}

func NewRealWorldService(uc *biz.GreeterUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)
