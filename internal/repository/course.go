package repository

import (
	"context"
	"github.com/IlyaZayats/restus/internal/entity"
	"github.com/IlyaZayats/restus/internal/interfaces"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresCourseRepository struct {
	db *pgxpool.Pool
}

func NewPostgresCourseRepository(db *pgxpool.Pool) (interfaces.CourseRepository, error) {
	return &PostgresCourseRepository{
		db: db,
	}, nil
}

func (r *PostgresCourseRepository) GetCourses() ([]entity.Course, error) {
	var courses []entity.Course
	q := "SELECT id, restaurant_id, name FROM Courses"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return courses, err
	}
	return r.parseRowsToSlice(rows)

}

func (r *PostgresCourseRepository) InsertCourse(course entity.Course) error {
	q := "INSERT INTO Courses (restaurant_id, name) VALUES ($1, $2)"
	if _, err := r.db.Exec(context.Background(), q, course.RestaurantId, course.Name); err != nil {
		return err
	}
	return nil
}

func (r *PostgresCourseRepository) UpdateCourse(course entity.Course) error {
	q := "UPDATE Courses SET name=$1 WHERE id=$2"
	if _, err := r.db.Exec(context.Background(), q, course.Name, course.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresCourseRepository) DeleteCourse(id int) error {
	q := "DELETE FROM Courses WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresCourseRepository) parseRowsToSlice(rows pgx.Rows) ([]entity.Course, error) {
	var slice []entity.Course
	defer rows.Close()
	for rows.Next() {
		var id, restaurantId int
		var name string
		if err := rows.Scan(&id, &restaurantId, &name); err != nil {
			return slice, err
		}
		slice = append(slice, entity.Course{Id: id, RestaurantId: restaurantId, Name: name})
	}
	return slice, nil
}
