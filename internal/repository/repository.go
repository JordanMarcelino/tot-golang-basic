package repository

import "tot_golang/internal/entity"

type Repository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(user *entity.User) error
	FindAll() ([]entity.User, error)
	FindById(user *entity.User) ([]*entity.User, error)
}
