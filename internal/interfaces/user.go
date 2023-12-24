package interfaces

import (
	"github.com/IlyaZayats/restus/internal/entity"
)

type UserRepository interface {
	Login(user entity.User) (int, error)
}
