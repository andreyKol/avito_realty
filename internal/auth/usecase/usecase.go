package usecase

import (
	"fmt"
	"realty/internal/auth"
	"realty/internal/common/jwt"
	"realty/internal/domain"
)

//go:generate ifacemaker -f *.go -o ../usecase.go -i UseCase -s UseCase -p auth -y "Controller describes methods, implemented by the usecase package."
type UseCase struct {
	authRepo auth.Repository
}

func NewAuthUseCase(authRepo auth.Repository) *UseCase {
	return &UseCase{authRepo: authRepo}
}

func (u *UseCase) Register(req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if req.Email != "" {
		err := u.authRepo.CheckEmailUnique(req.Email)
		if err != nil {
			return nil, fmt.Errorf("checking email uniqueness: %w", err)
		}
	}

	userInfo := domain.User{
		Email:             req.Email,
		PasswordEncrypted: req.Password,
		UserType:          req.UserType,
	}

	if err := userInfo.EncryptPassword(); err != nil {
		return nil, fmt.Errorf("encrypting password: %w", err)
	}

	res, err := u.authRepo.CreateUser(&userInfo)
	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}

	return &auth.RegisterResponse{
		UserID: res.ID,
	}, nil
}

func (u *UseCase) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := u.authRepo.GetUserByID(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by email: %w", err)
	}

	token, err := jwt.GenerateJWT(jwt.Claims{
		UserType: user.UserType,
	})
	if err != nil {
		return nil, fmt.Errorf("generating token: %w", err)
	}
	return &auth.LoginResponse{
		Token: token,
	}, nil
}
