package postgres

import (
	"github.com/SaidovZohid/jwt_token/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(u *repo.User) (*repo.User, error) {
	query := `
		INSERT INTO users(
			name,
			username,
			email,
			password
		) VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	var user repo.User

	err := ur.db.QueryRow(
		query,
		u.Name,
		u.Username,
		u.Email,
		u.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) Get(user_id int64) (*repo.User, error) {
	query := `
	SELECT
		id,
		name,
		username,
		email,
		password,
		created_at
	FROM users WHERE id = $1
`
	var user repo.User

	err := ur.db.QueryRow(
		query,
		user_id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
