package interfaces

import "github.com/IlyaZayats/restus/internal/entity"

type FoodRepository interface {
	GetFoods() ([]entity.Food, error)
	UpdateFood(food entity.Food) error
	InsertFood(food entity.Food) error
	DeleteFood(id int) error
}
