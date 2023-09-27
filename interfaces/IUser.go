package interfaces

import "tecnotree-go-project/entities"

type IUser interface {
	Register(user *entities.Register) (string, error)
	Login(user *entities.Login) (string, error)
	Logout() string
}
