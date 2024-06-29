package dto

import (
	"bbdk/domain/entity"
	"gorm.io/gorm"
)

type UserRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (req *UserRequest) ToEntity() *entity.User {
	return &entity.User{
		Model:    gorm.Model{ID: req.ID},
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (resp *UserResponse) FromEntity(user *entity.User) {
	resp.ID = user.ID
	resp.Name = user.Name
	resp.Email = user.Email
}
