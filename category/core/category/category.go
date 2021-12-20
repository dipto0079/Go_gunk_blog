package category

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

func (cs CoreSve) Create_ser(ctx context.Context, t storage.Category) (int64, error) {
	return cs.store.Create_sto(ctx, t)
}

func (cs CoreSve) Get_AllData_ser(ctx context.Context) ([]storage.Category, error) {
	return cs.store.Get_all_Data(ctx)
}
