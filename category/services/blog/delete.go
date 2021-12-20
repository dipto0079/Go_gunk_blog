package blog
import (
//	"blog/category/storage"
	tpb "blog/gunk/v1/blog"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc)Delete(ctx context.Context, req *tpb.DeleteBlogRequest) (*tpb.DeleteBlogResponse, error){
	
	err:=s.core.BlogDelete(context.Background(),req.GetID())
	if err != nil {
		return nil, status.Errorf(codes.Internal,"failed to create category: %s",err)
	}
	return &tpb.DeleteBlogResponse{},nil
	
}