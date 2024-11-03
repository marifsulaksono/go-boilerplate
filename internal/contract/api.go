package contract

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/common"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/service"
)

type Contract struct {
	Service    *service.Contract
	Repository *repository.Contract
	Common     *common.Contract
}

/*
This prt is the place to prepare dependency injection on internal projects.

more info contact me @marifsulaksono
*/
func NewContract(ctx context.Context) (*Contract, error) {
	common, err := common.NewCommon(ctx)
	if err != nil {
		return nil, err
	}

	repository, err := repository.NewRepository(ctx, common)
	if err != nil {
		return nil, err
	}

	service, err := service.NewService(ctx, repository)
	if err != nil {
		return nil, err
	}

	return &Contract{
		Service:    service,
		Repository: repository,
		Common:     common,
	}, nil
}
