package services

import (
	"github.com/IlyaZayats/restus/internal/entity"
	"github.com/IlyaZayats/restus/internal/interfaces"
	"strconv"
)

type CourseService struct {
	repo interfaces.CourseRepository
}

func NewCourseService(repo interfaces.CourseRepository) (*CourseService, error) {
	return &CourseService{
		repo: repo,
	}, nil
}

func (s *CourseService) GetCourses() ([]map[string]string, error) {
	courses, err := s.repo.GetCourses()
	if err != nil {
		return nil, err
	}
	coursesSlice := []map[string]string{}
	for _, item := range courses {
		coursesMap := map[string]string{
			"id":            strconv.Itoa(item.Id),
			"restaurant_id": strconv.Itoa(item.RestaurantId),
			"name":          item.Name,
		}
		coursesSlice = append(coursesSlice, coursesMap)
	}
	return coursesSlice, nil
}

func (s *CourseService) InsertCourse(restaurantId int, name string) error {
	return s.repo.InsertCourse(entity.Course{Id: 0, RestaurantId: restaurantId, Name: name})
}

func (s *CourseService) UpdateCourse(id int, name string) error {
	return s.repo.UpdateCourse(entity.Course{Id: id, RestaurantId: 0, Name: name})
}

func (s *CourseService) DeleteCourse(id int) error {
	return s.repo.DeleteCourse(id)
}
