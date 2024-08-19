package domain

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"realty/internal/common/errors"
	"regexp"
)

var (
	emailRegexp = regexp.MustCompile("[^@ \\t\\r\\n]+@[^@ \\t\\r\\n]+\\.[^@ \\t\\r\\n]+")
)

type User struct {
	ID                int64
	Email             string
	UserType          string
	PasswordEncrypted string
}

func (u *User) Validate() error {
	if !emailRegexp.MatchString(u.Email) {
		return errors.NewInvalidInputError("incorrect email format", "email")
	}

	return nil
}

func (u *User) EncryptPassword() error {
	data, err := bcrypt.GenerateFromPassword([]byte(u.PasswordEncrypted), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("encrypting password: %w", err)
	}

	u.PasswordEncrypted = string(data)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordEncrypted), []byte(password)) == nil
}
