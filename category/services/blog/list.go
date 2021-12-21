package blog

import (
	tpb "blog/gunk/v1/blog"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) ListBlog(ctx context.Context, req *tpb.ListBlogRequest) (*tpb.ListBlogResponse, error) {
	ids, err := s.core.ListBlog(context.Background())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create category: %s", err)
	}
	var bl []*tpb.Blog
	for _, v := range ids {
		bl = append(bl, &tpb.Blog{
			ID:         v.ID,
			CatID: v.CatID,
			Title:      v.Title,
			Description: v.Description,
			CatName: v.CatName,
		})
	}
	return &tpb.ListBlogResponse{
		Blog: bl,
	}, nil
}
