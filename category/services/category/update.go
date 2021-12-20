package category

import (
	"blog/category/storage"
	tpc "blog/gunk/v1/category"
	"context"
	

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
func (s *Svc) Update(ctx context.Context, req *tpc.UpdateCategoryRequest ) (*tpc.UpdateCategoryResponse, error) {
	category := storage.Category{
		ID: req.GetCategory().ID,
		Title: req.GetCategory().Title,
	}
	err := s.core.Update(context.Background(), category)
	if err != nil{
		return nil, status.Error(codes.Internal, "failed to Update category.")
	}
	return  &tpc.UpdateCategoryResponse{}, nil
}