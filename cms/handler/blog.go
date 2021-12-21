package handler

import (
	//"fmt"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	//"strconv"

	tpb "blog/gunk/v1/blog"
	tpc "blog/gunk/v1/category"

	validation "github.com/go-ozzo/ozzo-validation"
	//"github.com/gorilla/mux"
)

type Blog struct{
	ID          int64
	CatID       int64
	Title       string
	Description string
	Image       string
	CatName     string
	Category    []Category
	Errors      map[string]string
}



func (b *Blog) Validate() error {
	return validation.ValidateStruct(b,
		//validation.Field(&b.CatID, validation.Required.Error("This Filed cannot be blank")),
		validation.Field(&b.Title, validation.Required.Error("This Filed cannot be blank"), validation.Length(3, 0)),
		validation.Field(&b.Description, validation.Required.Error("This Filed cannot be blank"), validation.Length(3, 0)),
		//validation.Field(&b.Image, validation.Required.Error("This Filed cannot be blank")),
	)
}

const MAX_UPLOAD_SIZE = 1024 * 10024 // 1MB


func (h *Handler) BlogList(rw http.ResponseWriter, r *http.Request) {
	blogData, err := h.tb.ListBlog(r.Context(), &tpb.ListBlogRequest{})
	//fmt.Printf("=============data===================",blogData)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.templates.ExecuteTemplate(rw, "list-blog.html", blogData); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) BlogCreate(rw http.ResponseWriter, r *http.Request) {

	vErrs := map[string]string{}
	data, err := h.tc.GetAllData(r.Context(), &tpc.GetAllDataCategoryRequest{})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	cats := []Category{}

	for _, v := range data.Category {
		cats = append(cats, Category{
			ID:   v.ID,
			Title: v.Title,
		})
	}

	
	h.createPostData(rw, 0, "", "", cats, vErrs)
	return

}

func (h *Handler) createPostData(rw http.ResponseWriter, catId int64, title string, desc string, cats []Category, errs map[string]string) {

	form := Blog{
		CatID:       catId,
		Title:       title,
		Description: desc,
		Category:    cats,
		Errors:      errs,
	}

	fmt.Println(form)
	if err := h.templates.ExecuteTemplate(rw, "create_blog.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}



func (h *Handler) BlogStore(rw http.ResponseWriter, r *http.Request) {
	data, err := h.tc.GetAllData(r.Context(), &tpc.GetAllDataCategoryRequest{})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	cats := []Category{}
	for _, v := range data.Category {
		cats = append(cats, Category{
			ID:   v.ID,
			Title: v.Title,
		})
	}

	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	r.Body = http.MaxBytesReader(rw, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(rw, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("Image")

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	now := strconv.Itoa(int(time.Now().UnixNano()))
	img := "upload-*" + now + filepath.Ext(fileHeader.Filename)
	tempFile, err := ioutil.TempFile("cms/assets/uploads", img)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	tempFile.Write(fileBytes)
	imgName := tempFile.Name()

	var blog Blog
	if err := h.decoder.Decode(&blog, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := blog.Validate(); err != nil {
		valError, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range valError {
				vErrs[key] = value.Error()
			}
			h.createPostData(rw, blog.CatID, blog.Title, blog.Description, cats, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.tb.CreateBlog(r.Context(), &tpb.CreateBlogRequest{
		Blog: &tpb.Blog{
			CatID:       blog.CatID,
			Title:       blog.Title,
			Description: blog.Description,
			Image:       imgName,
		},
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}