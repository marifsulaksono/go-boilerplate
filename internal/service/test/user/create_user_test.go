package test

import (
	"context"
	"fmt"
	"testing"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service"
	testconfig "github.com/marifsulaksono/go-echo-boilerplate/internal/service/test"
	"github.com/stretchr/testify/assert"
)

var (
	ctx         = context.Background()
	errExpected = fmt.Errorf("expected error")
	data        = model.User{
		Name:     "Muhammad Arif Sulaksono",
		Email:    "marifsulaksono@gmail.com",
		Password: "password123",
	}
)

func TestCreateUser(t *testing.T) {
	mock := testconfig.SetupAppTest(t)
	svc := service.NewUserService(mock.ContractRepo)
	defer mock.Ctrl.Finish()

	t.Run("Should be error generate hashed password", func(t *testing.T) {
		gen := monkey.Patch(helper.GenerateHashedPassword, func(password string) (string, error) {
			return "", errExpected
		})
		defer gen.Unpatch()

		_, err := svc.Create(ctx, &data)
		assert.Error(t, err)
		assert.ErrorIs(t, err, errExpected)
	})

	t.Run("Should be error create user", func(t *testing.T) {
		gen := monkey.Patch(helper.GenerateHashedPassword, func(password string) (string, error) {
			return "hashedPassword", nil
		})
		defer gen.Unpatch()

		mock.UserRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return("", errExpected).Times(1)

		_, err := svc.Create(ctx, &data)
		assert.Error(t, err)
		assert.ErrorIs(t, err, errExpected)
	})
}
