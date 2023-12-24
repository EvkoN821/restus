package requests

type InsertRestaurantRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateRestaurantRequest struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteRestaurantRequest struct {
	Id int `json:"id" binding:"required"`
}
