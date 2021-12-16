package category

import (
	"context"
	"blog/category/storage"
)

type categoryStore interface {
	Create_sto(context.Context, storage.Category) (int64, error)
}

type CoreSve struct{
	store categoryStore
}

func NewCoreSve(b categoryStore) *CoreSve {
	return &CoreSve{
		store: b,
	}
}

func (cs CoreSve) Create_ser(ctx context.Context, t storage.Category) (int64, error) {
	return cs.store.Create_sto(ctx, t)
	// return 0, nil
}