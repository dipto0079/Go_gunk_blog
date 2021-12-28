package user

import (
	"blog/category/storage"
	tpu "blog/gunk/v1/user"
	"context"
)

type UserCoreStore interface {
	Create(context.Context, storage.User) (int64, error)
}

type Svc struct {
	tpu.UnimplementedUserRegServiceServer
	core UserCoreStore
}

func NewUserServer(c UserCoreStore) *Svc {
	return &Svc{
		core: c,
	}
}

