package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DenisHoliahaR/go-beautyhub/internal/domain"
	"github.com/DenisHoliahaR/go-beautyhub/internal/repository"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
	INSERT INTO users (first_name, second_name, email, phone, password)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at
`

	err := r.db.QueryRowContext(ctx, query,
		user.FirstName,
		user.SecondName,
		user.Email,
		user.Phone,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("user repository create: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `
    SELECT id, first_name, second_name, email, phone, created_at
    FROM users
    WHERE id = $1
`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.SecondName,
		&user.Email,
		&user.Phone,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetList(ctx context.Context) ([]*domain.User, error) {
	query := `
    SELECT id, first_name, second_name, email, phone, created_at
    FROM users
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("user repository get list: %w", err)
	}
	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		u := &domain.User{}
		if err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.SecondName,
			&u.Email,
			&u.Phone,
			&u.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
		UPDATE users
		SET first_name = $1, second_name = $2, email = $3, phone = $4, password = $5
		WHERE id = $5
	`

	if _, err := r.db.ExecContext(ctx, query,
		user.FirstName,
		user.SecondName,
		user.Email,
		user.Phone,
		user.ID,
	); err != nil {
		return nil, fmt.Errorf("user repository update: %w", err)
	}

	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
