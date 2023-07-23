package dto

type UserData struct {
	ID       int64  `db:"id"`
	FullName string `db:"full_name"`
	Phone    string `db:"phone"`
	UserName string `db:"user_name"`
}
