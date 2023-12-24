package interfaces

import "github.com/IlyaZayats/restus/internal/entity"

type RestaurantRepository interface {
	GetRestaurants() ([]entity.Restaurant, error)
	UpdateRestaurant(restaurant entity.Restaurant) error
	InsertRestaurant(restaurant entity.Restaurant) error
	DeleteRestaurant(id int) error
}
