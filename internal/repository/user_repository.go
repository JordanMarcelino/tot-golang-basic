package repository

import (
	"tot_golang/internal/entity"
	"tot_golang/internal/util"
)

var users = []*entity.User{
	{ID: 1, Name: "Arthur Hitam", Division: "Engineering"},
	{ID: 2, Name: "Sukma Adhi Wijawa", Division: "Engineering"},
	{ID: 3, Name: "Jric Erianto", Division: "CEO"},
	{ID: 4, Name: "Srian Bangapta", Division: "RnD"},
	{ID: 5, Name: "Sizky Riregar", Division: "Engineering"},
}

type UserRepository struct {
	Users []*entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{Users: users}
}

func (u *UserRepository) Create(user *entity.User) error {
	lastId := u.Users[len(u.Users)-1].ID

	user.ID = lastId + 1

	u.Users = append(u.Users, user)

	return nil
}

func (u *UserRepository) Update(user *entity.User) error {
	if err := u.FindById(&entity.User{ID: user.ID}); err != nil {
		return err
	}

	for _, us := range u.Users {
		if us.ID == user.ID {

			if user.Name != "" {
				us.Name = user.Name
			}

			if user.Division != "" {
				us.Division = user.Division
			}

		}
	}

	return nil
}

func (u *UserRepository) Delete(user *entity.User) error {
	if err := u.FindById(user); err != nil {
		return err
	}

	var newUsers []*entity.User

	for _, us := range u.Users {
		if us.ID != user.ID {
			newUsers = append(newUsers, us)
		}
	}

	u.Users = newUsers

	return nil
}

func (u *UserRepository) FindAll() ([]*entity.User, error) {
	return u.Users, nil
}

func (u *UserRepository) FindById(user *entity.User) error {
	for _, us := range u.Users {
		if us.ID == user.ID {
			*user = *us
			return nil
		}
	}
	return util.ErrorNotFound
}
