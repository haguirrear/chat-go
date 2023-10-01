package repository

import (
	"chat/entities"
	"context"
	"database/sql"
	"log"
	"time"
)

type repository struct {
	db *sql.DB
}

func NewSqlUserRepository(db *sql.DB) UserRepository {
	return repository{db: db}
}

func (r repository) CreateUser(ctx context.Context, user *entities.User) error {
	var id int
	query := `
		--sql 
		INSERT INTO users(email, password, created_at) VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, user.Email, user.Password, time.Now()).Scan(&id)

	if err != nil {
		return err
	}

	user.Id = id

	return nil

}

func (r repository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `
		--sql
		SELECT
		    u.id,
		    u.email,
		    u.password,
		    u.is_active,
			u.created_at,
			u.updated_at
		FROM
		    users u
		WHERE
		    u.email = $1
			and u.is_active
	`

	u := entities.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&u.Id,
		&u.Email,
		&u.Password,
		&u.IsActive,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	switch err {
	case nil:
		return &u, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Panicf("error retriveing user: %v", err)
		return nil, nil

	}

}
