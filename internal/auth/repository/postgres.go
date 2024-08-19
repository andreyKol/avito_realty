package postgresql

import (
	"fmt"
	"realty/internal/domain"
	"realty/pkg/storage/postgres"
)

//go:generate ifacemaker -f postgres.go -o ../repository.go -i Repository -s PostgresRepository -p auth -y "Controller describes methods, implemented by the repository package."
type PostgresRepository struct {
	db postgres.Postgres
}

func NewPostgresRepository(db postgres.Postgres) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) CreateUser(user *domain.User) (*domain.User, error) {
	var userID int64
	err := p.db.QueryRow(`
		INSERT INTO users(email, password_enc, user_type)
		VALUES ($1, $2, $3) RETURNING id`,
		user.Email,
		user.PasswordEncrypted,
		user.UserType,
	).Scan(&userID)

	if err != nil {
		return nil, fmt.Errorf("user has already been created: %w", err)
	}

	createdUser := &domain.User{
		ID:                userID,
		Email:             user.Email,
		UserType:          user.UserType,
		PasswordEncrypted: user.PasswordEncrypted,
	}

	return createdUser, nil
}

func (p *PostgresRepository) GetUserByID(id int64) (*domain.User, error) {
	var user domain.User

	err := p.db.QueryRow(`
		SELECT email, password_enc, user_type
		FROM users
		WHERE id = $1`, id,
	).Scan(&user.Email, &user.PasswordEncrypted, &user.UserType)

	if err != nil {
		return nil, fmt.Errorf("querying user: %w", err)
	}

	return &user, nil
}

func (p *PostgresRepository) CheckEmailUnique(email string) error {
	var count int
	err := p.db.QueryRow(`
		SELECT COUNT(*)
		FROM users
		WHERE email = $1`, email,
	).Scan(&count)

	if err != nil {
		return fmt.Errorf("querying email: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("email already exists: %w", err)
	}

	return nil
}
