package entity

type User struct {
	Id       int    `db:"id"`
	Login    string `db:"login"`
	Pwd      string `db:"pwd"`
	UserType int    `db:"user_type"`
}
