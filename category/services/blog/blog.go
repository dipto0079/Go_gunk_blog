package blog

import (
	"blog/category/storage"
	tpc "blog/gunk/v1/blog"
	"context"
)

type blogCoreStore interface {
	Create(context.Context, storage.Blog) (int64, error)
	ListBlog(context.Context) ([]storage.Blog, error)
	GetBlog(context.Context,int64) (storage.Blog, error)
	// Delete(context.Context, int64)  error
	// Update(context.Context, storage.Category) error
}

type Svc struct {
	tpc.UnimplementedBlogServiceServer
	core blogCoreStore
}

func NewCategoryServer(b blogCoreStore) *Svc {
	return &Svc{
		core: b,
	}
}
