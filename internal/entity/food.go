package entity

type Food struct {
	Id       int    `db:"id"`
	CourseId int    `db:"course_id"`
	Name     string `db:"name"`
	Weight   int    `db:"weight"`
	Price    int    `db:"price"`
	Calories int    `db:"calories"`
	Info     string `db:"info"`
	Comp     string `db:"comp"`
	Prep     int    `db:"prep"`
}
