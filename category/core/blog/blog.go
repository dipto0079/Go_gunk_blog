package blog

import (
	"blog/category/storage"
	"blog/category/storage/postgres"
	"context"
)

type CoreSve struct {
	store *postgres.Storage
}

func NewCoreSve(b *postgres.Storage) *CoreSve {
	return &CoreSve{
		store: b,
	}
}


func (cs CoreSve) Create(ctx context.Context, re storage.Blog) (int64, error) {
	return cs.store.Create(ctx, re)
}

func (cs CoreSve) ListBlog(ctx context.Context) ([]storage.Blog, error) {
	return cs.store.ListBlog(ctx)
}

func (cs CoreSve) GetBlog(ctx context.Context,id int64)  (storage.Blog, error) {
	return cs.store.GetBlog(ctx, id)
}
