package user

import "gocionics/internal/entities/user"

type IUserRepository interface {
	Create(user *user.User) (string, error)
	GetById(id string) (user.User, error)
	AssignCharacter(userID, characterID int) (string, error)
}
