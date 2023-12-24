package entity

type Course struct {
	Id           int    `db:"id"`
	RestaurantId int    `db:"restaurant_id"`
	Name         string `db:"name"`
}
