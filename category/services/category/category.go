package category

import (
	"blog/category/storage"
	tpc "blog/gunk/v1/category"
	"context"

)

type categoryCoreStore interface{
	Create_ser(context.Context,storage.Category) (int64, error)
	Get_ser(context.Context,storage.Category) (storage.Category, error)
}

type Svc struct{
	tpc.UnimplementedCategoryServiceServer
	core categoryCoreStore
}

func NewCategoryServer(c categoryCoreStore) *Svc {
	return &Svc{
		core: c,
	}
}