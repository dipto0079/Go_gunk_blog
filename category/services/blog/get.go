package blog

import (
	"blog/category/storage"
	tpb "blog/gunk/v1/blog"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Svc) GetBlog(ctx context.Context,req *tpb.GetBlogRequest) (*tpb.GetBlogResponse, error){
	
	var blo storage.Blog

	blo, err := s.core.GetBlog(context.Background(), req.GetID())
	if err != nil{
		return nil, status.Error(codes.Internal, "failed to get Blog.")
	}

	return  &tpb.GetBlogResponse{
		Blog : &tpb.Blog{
			ID: blo.ID,
			Cat_ID: blo.Cat_ID,
			Title: blo.Title,
			Description: blo.Description,
			Image: blo.Image,
		},
	}, nil

}