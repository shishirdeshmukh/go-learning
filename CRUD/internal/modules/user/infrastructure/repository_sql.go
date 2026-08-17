package infrastructure

import (
	"context"
	"database/sql"

	"GO-Crud/internal/modules/user/domain"
	"GO-Crud/internal/modules/user/ports"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {

	query := `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
}

func (r *userRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.User, error) {

	query := `
		SELECT id, name, email, created_at
		FROM users
		WHERE id = $1
	`

	var user domain.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetAll(
	ctx context.Context,
) ([]domain.User, error) {

	query := `
		SELECT id, name, email, created_at
		FROM users
		ORDER BY id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Update(
	ctx context.Context,
	user *domain.User,
) error {

	query := `
		UPDATE users
		SET name = $1,
		    email = $2
		WHERE id = $3
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.ID,
	)

	return err
}

func (r *userRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		id,
	)

	return err
}
