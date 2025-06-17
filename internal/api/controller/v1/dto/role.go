package dto

import "github.com/marifsulaksono/go-echo-boilerplate/internal/model"

type (
	RoleRequest struct {
		Name string `json:"name" validate:"required"`
	}

	GetRoleRequest struct {
		Search string `json:"search"`
	}
)

func (d *RoleRequest) ParseToModel() *model.Role {
	return &model.Role{
		Name: d.Name,
	}
}

func (d *GetRoleRequest) ParseToModel() *model.RoleRequest {
	return &model.RoleRequest{
		Search: d.Search,
	}
}
