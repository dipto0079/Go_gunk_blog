package blog

import (
	"blog/category/storage"
	tpc "blog/gunk/v1/blog"
	"context"
)

type blogCoreStore interface {
	Create(context.Context, storage.Blog) (int64, error)
	// Get_AllData_ser(context.Context) ([]storage.Category, error)
	// Get_single_ser(context.Context,int64) (storage.Category, error)
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
