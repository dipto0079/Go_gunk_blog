package storage

type Category struct{
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	IsComplete  bool   `db:"is_completed"`
}