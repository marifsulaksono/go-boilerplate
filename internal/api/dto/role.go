package dto

import "github.com/marifsulaksono/go-echo-boilerplate/internal/model"

type RoleRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r *RoleRequest) ParseToModel() *model.Role {
	return &model.Role{
		Name: r.Name,
	}
}
