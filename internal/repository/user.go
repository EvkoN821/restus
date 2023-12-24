package repository

import (
	"context"
	"github.com/IlyaZayats/restus/internal/entity"
	"github.com/IlyaZayats/restus/internal/interfaces"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) (interfaces.UserRepository, error) {
	return &PostgresUserRepository{
		db: db,
	}, nil
}

func (r *PostgresUserRepository) Login(user entity.User) (int, error) {
	q := "SELECT user_type FROM Users WHERE login=$1 AND pwd=$2"
	rows, err := r.db.Query(context.Background(), q, user.Login, user.Pwd)
	if err != nil && err.Error() != "no rows in result set" {
		return -1, err
	}
	return r.parseRow(rows)
}

func (r *PostgresUserRepository) parseRow(rows pgx.Rows) (int, error) {
	var userType int
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&userType); err != nil {
			return -1, err
		}
	}
	return userType, nil
}
