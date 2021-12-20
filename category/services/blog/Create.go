package blog

import (
	"blog/category/storage"
	tpb "blog/gunk/v1/blog"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) Create(ctx context.Context,req *tpb.CreateBlogRequest) (*tpb.CreateBlogResponse, error){
	//log.Printf("Request Category : %#v\n",req.GetCategory())

	category :=storage.Blog{
		Cat_ID: req.Blog.Cat_ID,
		Title: req.Blog.Title,
		Description: req.Blog.Description,
		Image: req.Blog.Image,
	}
	
	id,err:=s.core.Create(context.Background(),category)
	if err != nil {
		return nil, status.Errorf(codes.Internal,"failed to create category: %s",err)
	}
	return &tpb.CreateBlogResponse{
		ID: id,
	},nil

}