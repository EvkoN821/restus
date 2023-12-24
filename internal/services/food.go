package services

import (
	"github.com/IlyaZayats/restus/internal/entity"
	"github.com/IlyaZayats/restus/internal/interfaces"
	"strconv"
)

type FoodService struct {
	repo interfaces.FoodRepository
}

func NewFoodService(repo interfaces.FoodRepository) (*FoodService, error) {
	return &FoodService{
		repo: repo,
	}, nil
}

func (s *FoodService) GetFoods() ([]map[string]string, error) {
	foods, err := s.repo.GetFoods()
	if err != nil {
		return nil, err
	}
	foodsSlice := []map[string]string{}
	for _, item := range foods {
		foodsMap := map[string]string{
			"id":        strconv.Itoa(item.Id),
			"course_id": strconv.Itoa(item.CourseId),
			"name":      item.Name,
			"weight":    strconv.Itoa(item.Weight),
			"price":     strconv.Itoa(item.Price),
			"calories":  strconv.Itoa(item.Calories),
			"info":      item.Info,
		}
		foodsSlice = append(foodsSlice, foodsMap)
	}
	return foodsSlice, nil
}

func (s *FoodService) InsertFood(name, info string, courseId, weight, price, calories int) error {
	return s.repo.InsertFood(entity.Food{Id: 0, CourseId: courseId, Name: name, Weight: weight, Price: price, Calories: calories, Info: info})
}

func (s *FoodService) UpdateFood(name, info string, id, weight, price, calories int) error {
	return s.repo.UpdateFood(entity.Food{Id: id, CourseId: 0, Name: name, Weight: weight, Price: price, Calories: calories, Info: info})
}

func (s *FoodService) DeleteFood(id int) error {
	return s.repo.DeleteFood(id)
}
