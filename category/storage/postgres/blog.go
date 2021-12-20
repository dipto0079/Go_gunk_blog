package postgres

import (
	"blog/category/storage"
	"context"
)


const insertBlog = `
	INSERT INTO blogs(
		cat_id,
		title,
		description,
		image

	)VALUES(
		:cat_id,
		:title,
		:description,
		:image
	)RETURNING id;
`

func (s *Storage) Create(ctx context.Context, t storage.Blog) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertBlog)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) ListBlog(ctx context.Context) ([]storage.Blog, error) {

	var b []storage.Blog

	if err := s.db.Select(&b, "SELECT * FROM blogs order by id desc"); err != nil {
		return b, err
	}
	return b, nil
}


func (s *Storage) GetBlog(ctx context.Context, id int64) (storage.Blog, error) {
	var b storage.Blog

	if err := s.db.Get(&b, "SELECT * FROM blogs WHERE id=$1", id); err != nil {
		return b, err
	}
	return b, nil
}

