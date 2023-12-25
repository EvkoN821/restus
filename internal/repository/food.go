package repository

import (
	"context"
	"github.com/IlyaZayats/restus/internal/entity"
	"github.com/IlyaZayats/restus/internal/interfaces"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresFoodRepository struct {
	db *pgxpool.Pool
}

func NewPostgresFoodRepository(db *pgxpool.Pool) (interfaces.FoodRepository, error) {
	return &PostgresFoodRepository{
		db: db,
	}, nil
}

func (r *PostgresFoodRepository) GetFoods() ([]entity.Food, error) {
	var foods []entity.Food
	q := "SELECT id, course_id, name, weight, price, calories, info, comp, prep FROM Foods"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return foods, err
	}
	//foods, err =
	return r.parseRowsToSlice(rows)

}

func (r *PostgresFoodRepository) InsertFood(food entity.Food) error {
	q := "INSERT INTO Foods (course_id, name, weight, price, calories, info, comp, prep) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	if _, err := r.db.Exec(context.Background(), q, food.CourseId, food.Name, food.Weight, food.Price, food.Calories, food.Info, food.Comp, food.Prep); err != nil {
		return err
	}
	return nil
}

func (r *PostgresFoodRepository) UpdateFood(food entity.Food) error {
	q := "UPDATE Foods SET (name, weight, price, calories, info, comp, prep) = ($1, $2, $3, $4, $5, $6, $7) WHERE id=$8"
	if _, err := r.db.Exec(context.Background(), q, food.Name, food.Weight, food.Price, food.Calories, food.Info, food.Comp, food.Prep, food.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresFoodRepository) DeleteFood(id int) error {
	q := "DELETE FROM Foods WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresFoodRepository) parseRowsToSlice(rows pgx.Rows) ([]entity.Food, error) {
	var slice []entity.Food
	defer rows.Close()
	for rows.Next() {
		var id, courseId, weight, price, calories, prep int
		var name, info, comp string
		if err := rows.Scan(&id, &courseId, &name, &weight, &price, &calories, &info, &comp, &prep); err != nil {
			return slice, err
		}
		slice = append(slice, entity.Food{
			Id: id, CourseId: courseId, Name: name, Weight: weight, Price: price, Calories: calories, Info: info, Comp: comp, Prep: prep})
	}
	return slice, nil
}
