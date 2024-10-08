// Code generated by ifacemaker; DO NOT EDIT.

package auth

import (
	"realty/internal/domain"
)

// Controller describes methods, implemented by the repository package.
type Repository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserByID(id int64) (*domain.User, error)
	CheckEmailUnique(email string) error
}
