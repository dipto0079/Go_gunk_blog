package category


// import (
// 	//"blog/category/storage"
// 	tpc "blog/gunk/v1/category"
// 	"context"
	
// 	"blog/category/storage"
// 	 "google.golang.org/grpc/codes"
// 	 "google.golang.org/grpc/status"
// )
// //Get(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error)
// func(s *Svc) Get(ctx context.Context,id int64) (*tpc.GetCategoryResponse, error){
	
// 	category :=storage.Category{
// 		Title: req.Category.Title,
// 	}
	
// 	id,err:=s.core.Create_ser(context.Background(),category)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal,"failed to create category: %s",err)
// 	}
// 	return &tpc.GetCategoryResponse{
// 		ID: id,
// 	},nil


// }