package interfaces

import "github.com/IlyaZayats/restus/internal/entity"

type CourseRepository interface {
	GetCourses() ([]entity.Course, error)
	UpdateCourse(course entity.Course) error
	InsertCourse(course entity.Course) error
	DeleteCourse(id int) error
}
