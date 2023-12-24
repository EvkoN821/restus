package services

import (
	"github.com/IlyaZayats/restus/internal/entity"
	"github.com/IlyaZayats/restus/internal/interfaces"
	"strconv"
)

type RestaurantService struct {
	repo interfaces.RestaurantRepository
}

func NewRestaurantService(repo interfaces.RestaurantRepository) (*RestaurantService, error) {
	return &RestaurantService{
		repo: repo,
	}, nil
}

func (s *RestaurantService) GetRestaurants() ([]map[string]string, error) {
	restaurants, err := s.repo.GetRestaurants()
	if err != nil {
		return nil, err
	}
	restaurantsSlice := []map[string]string{}
	for _, item := range restaurants {
		facultiesMap := map[string]string{
			"id":   strconv.Itoa(item.Id),
			"name": item.Name,
		}
		restaurantsSlice = append(restaurantsSlice, facultiesMap)
	}
	return restaurantsSlice, nil
}

func (s *RestaurantService) InsertRestaurant(name string) error {
	return s.repo.InsertRestaurant(entity.Restaurant{Id: 0, Name: name})
}

func (s *RestaurantService) UpdateRestaurant(id int, name string) error {
	return s.repo.UpdateRestaurant(entity.Restaurant{Id: id, Name: name})
}

func (s *RestaurantService) DeleteRestaurant(id int) error {
	return s.repo.DeleteRestaurant(id)
}
