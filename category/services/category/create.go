package category

import (
	"blog/category/storage"
	tpc "blog/gunk/v1/category"
	"context"
	

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) Create(ctx context.Context,req *tpc.CreateCategoryRequest) (*tpc.CreateCategoryResponse, error){
	//log.Printf("Request Category : %#v\n",req.GetCategory())

	category :=storage.Category{
		Title: req.Category.Title,
	}
	
	id,err:=s.core.Create_ser(context.Background(),category)
	if err != nil {
		return nil, status.Errorf(codes.Internal,"failed to create category: %s",err)
	}
	return &tpc.CreateCategoryResponse{
		ID: id,
	},nil

}