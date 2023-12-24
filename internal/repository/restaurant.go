package repository

import (
	"context"
	"github.com/IlyaZayats/restus/internal/entity"
	"github.com/IlyaZayats/restus/internal/interfaces"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRestaurantRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRestaurantRepository(db *pgxpool.Pool) (interfaces.RestaurantRepository, error) {
	return &PostgresRestaurantRepository{
		db: db,
	}, nil
}

func (r *PostgresRestaurantRepository) GetRestaurants() ([]entity.Restaurant, error) {
	var faculties []entity.Restaurant
	q := "SELECT id, name FROM Restaurants"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return faculties, err
	}
	//faculties, err =
	return r.parseRowsToSlice(rows)

}

func (r *PostgresRestaurantRepository) InsertRestaurant(restaurant entity.Restaurant) error {
	q := "INSERT INTO Restaurants (name) VALUES ($1)"
	if _, err := r.db.Exec(context.Background(), q, restaurant.Name); err != nil {
		return err
	}
	return nil
}

func (r *PostgresRestaurantRepository) UpdateRestaurant(restaurant entity.Restaurant) error {
	q := "UPDATE Restaurants SET name=$1 WHERE id=$2"
	if _, err := r.db.Exec(context.Background(), q, restaurant.Name, restaurant.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresRestaurantRepository) DeleteRestaurant(id int) error {
	q := "DELETE FROM Restaurants WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresRestaurantRepository) parseRowsToSlice(rows pgx.Rows) ([]entity.Restaurant, error) {
	var slice []entity.Restaurant
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return slice, err
		}
		slice = append(slice, entity.Restaurant{Id: id, Name: name})
	}
	return slice, nil
}
