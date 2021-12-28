package storage

type Category struct {
	ID         int64  `db:"id"`
	Title      string `db:"title"`
	IsComplete bool   `db:"is_completed"`
}

type Blog struct {
	ID          int64  `db:"id"`
	CatID       int64  `db:"cat_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Image       string `db:"image"`
	CatName     string `db:"catname"`
}

type User struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	EmailVerified string `db:"email_verified"`
	IsVerify bool   `db:"is_verify"`
}
