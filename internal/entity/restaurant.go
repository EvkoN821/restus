package entity

type Restaurant struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
