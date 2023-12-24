package requests

//type GetCoursesRequest struct {
//	Id int `json:"id" binding:"required"`
//}

type InsertCourseRequest struct {
	RestaurantId int    `json:"restaurant_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
}

type UpdateCourseRequest struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteCourseRequest struct {
	Id int `json:"id" binding:"required"`
}
