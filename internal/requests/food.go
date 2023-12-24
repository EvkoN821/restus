package requests

//type GetFoodsRequest struct {
//	Id int `json:"id" binding:"required"`
//}

type InsertFoodRequest struct {
	CourseId int    `json:"course_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Weight   int    `json:"weight" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Calories int    `json:"calories" binding:"required"`
	Info     string `json:"info" binding:"required"`
}

type UpdateFoodRequest struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Weight   int    `json:"weight" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Calories int    `json:"calories" binding:"required"`
	Info     string `json:"info" binding:"required"`
}

type DeleteFoodRequest struct {
	Id int `json:"id" binding:"required"`
}
