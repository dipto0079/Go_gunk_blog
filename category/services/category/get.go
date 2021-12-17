 package category


// import (
// 	"blog/category/storage"
// 	tpc "blog/gunk/v1/category"
// 	"context"
// 	//"log"

// //	"google.golang.org/grpc/codes"
// //	"google.golang.org/grpc/status"
// )

// func Get(ctx context.Context,req *tpc.GetCategoryRequest) (*tpc.GetCategoryResponse, error){
// 	category :=storage.Category{}


	
// 	id,err:=s.core.Get_ser(context.Background(),category)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal,"failed to create category: %s",err)
// 	}
// 	return &tpc.CreateCategoryResponse{
// 		ID: id,
// 	},nil

// }