 package category


import (
	//"blog/category/storage"
	tpc "blog/gunk/v1/category"
	"context"
	//"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Svc) GetAllData(ctx context.Context,) (*tpc.GetAllDataCategoryResponse, error){
	
	 //category := storage.Category{}
	
	id,err:=s.core.Get_AllData_ser(context.Background())
	if err != nil {
		return nil, status.Errorf(codes.Internal,"failed to create category: %s",err)
	}
	return &tpc.GetAllDataCategoryResponse{
		Category: &tpc.Category{
			ID: id.ID,
			Title: id.Title,
		},
	},nil

}