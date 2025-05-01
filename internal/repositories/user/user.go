package user

import "gocionics/internal/entities/user"

type IUserRepository interface {
	Create(user *user.User) (string, error)
}
