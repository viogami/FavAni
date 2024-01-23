package repos

import (
	"FavAni/database"
)

type Repository interface {
	User() UserRepository
}

type Migrant interface {
	Migrate() error
}

type UserRepository interface {
	Login(username, password string) (*database.User, error)
	Logout() error
	Register(newUser database.User) error
	GetUserByName(string) (*database.User, error)
	List() (database.Users, error)
	Create(*database.User) (*database.User, error)
	Update(*database.User) (*database.User, error)
	Delete(user database.User) error
}
