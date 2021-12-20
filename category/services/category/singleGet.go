package category


import (
	"blog/category/storage"
	tpc "blog/gunk/v1/category"
	"context"
	
	  "google.golang.org/grpc/codes"
	  "google.golang.org/grpc/status"
)

func(s *Svc) Get(ctx context.Context,req *tpc.GetCategoryRequest) (*tpc.GetCategoryResponse, error){
	
	var cat storage.Category

	cat, err := s.core.Get_single_ser(context.Background(), req.GetID())
	if err != nil{
		return nil, status.Error(codes.Internal, "failed to get category.")
	}

	return  &tpc.GetCategoryResponse{
		Category : &tpc.Category{
			ID: cat.ID,
			Title: cat.Title,
			IsComplete: cat.IsComplete,
		},
	}, nil

}