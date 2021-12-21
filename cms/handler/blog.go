package handler

import (
	//"fmt"
	"net/http"
	//"strconv"

	//	"strconv"

	//tpb "blog/gunk/v1/blog"
	//tpc "blog/gunk/v1/category"

	validation "github.com/go-ozzo/ozzo-validation"
	//"github.com/gorilla/mux"
)

type Blog struct{
	ID          int64  
	Cat_ID      int64  
	Title       string 
	Description string 
	Image       string 
}

type BlogData struct{
	Blog []Blog
	Errors      map[string]string
	Category Category
}

func (b *Blog) Validate() error {
	return validation.ValidateStruct(b,
		validation.Field(&b.Cat_ID, validation.Required.Error("This Filed cannot be blank")),
		validation.Field(&b.Title, validation.Required.Error("This Filed cannot be blank"), validation.Length(3, 0)),
		validation.Field(&b.Description, validation.Required.Error("This Filed cannot be blank"), validation.Length(3, 0)),
		validation.Field(&b.Image, validation.Required.Error("This Filed cannot be blank"), validation.Length(3, 0)),
	)
}

func (h *Handler) BlogCreate(rw http.ResponseWriter, r *http.Request) {

	// vErrs :=map[string]string{"title": "","description":"","image":""}
	// cat_id := 0
	// title := ""
	// description :=""
	// image :=""
	
	// ctx := r.Context()
	// catAll, err := h.tc.GetAllData(ctx, &tpc.GetAllDataCategoryRequest{})
	// fmt.Printf("%+v", err)
	

	// h.createBlogFormData(rw,cat_id, title,description,image,vErrs)
	return

}

func (h *Handler) createBlogFormData(rw http.ResponseWriter,catId int64, title string,description string,image string, errs map[string]string) {
	
	
	

	form := Category{
		Title:       title,
		Errors:      errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "create_Category.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}