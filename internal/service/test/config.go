package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_repo_contract "github.com/marifsulaksono/go-echo-boilerplate/shared/mock/contract/repository"
	mock_repository "github.com/marifsulaksono/go-echo-boilerplate/shared/mock/repository"
)

type SetupApp struct {
	Ctrl *gomock.Controller

	ContractRepo *mock_repo_contract.MockRepositoryContract

	UserRepo *mock_repository.MockUserRepository
}

func SetupAppTest(t *testing.T) *SetupApp {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_repository.NewMockUserRepository(ctrl)

	contractRepo := mock_repo_contract.NewMockRepositoryContract(ctrl)

	contractRepo.EXPECT().GetUser().Return(userRepo).AnyTimes()

	return &SetupApp{
		Ctrl:         ctrl,
		ContractRepo: contractRepo,
		UserRepo:     userRepo,
	}
}
