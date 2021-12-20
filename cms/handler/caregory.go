package handler

import (
	"net/http"
	"strconv"
	//	"strconv"

	tpb "blog/gunk/v1/category"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"
)

type Category struct {
	ID          int64
	Title       string
	Errors      map[string]string
}

func (c *Category) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Title, validation.Required.Error("This Filed cannot be blank"), validation.Length(3, 0)),
	)
}

// Add
func (h *Handler) CategoryCreate(rw http.ResponseWriter, r *http.Request) {

	vErrs := map[string]string{"title": ""}
	title := ""
	h.createCategoryFormData(rw, title, vErrs)
	return

}

func (h *Handler) createCategoryFormData(rw http.ResponseWriter, title string, errs map[string]string) {

	form := Category{
		Title:       title,
		Errors:      errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "create_Category.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

//Store
func (h *Handler) CategoryStore(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	var category Category
	if err := h.decoder.Decode(&category, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if aErr := category.Validate(); aErr != nil {
		//fmt.Printf("%T", aErr)
		vErrors, ok := aErr.(validation.Errors)
		if ok {
			vErr := make(map[string]string)
			for key, value := range vErrors {
				vErr[key] = value.Error()
			}
			h.createCategoryFormData(rw, category.Title, vErr)
			return
		}

		http.Error(rw, aErr.Error(), http.StatusInternalServerError)
		return
	}

	

	_, err := h.tc.Create(r.Context(), &tpb.CreateCategoryRequest{
		Category: &tpb.Category{
			Title: category.Title,
		},
	})
	if err != nil {
		http.Error(rw, "Invalid URL", http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}


//Delete
func (h *Handler) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]

	id, erre := strconv.ParseInt(Id, 10, 64)

	if erre != nil {
		http.Error(rw, erre.Error(), http.StatusInternalServerError)
		return
	}

	_, err := h.tc.Delete(r.Context(), &tpb.DeleteCategoryRequest{
		ID: id,
	})
	if err != nil {
		http.Error(rw, "Invalid URL", http.StatusInternalServerError)
		return
	}
	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}


//Edit
func (h *Handler) Edit(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]

	id, erre := strconv.ParseInt(Id, 10, 64)

	if erre != nil {
		http.Error(rw, erre.Error(), http.StatusInternalServerError)
		return
	}

	res, err := h.tc.Get(r.Context(), &tpb.GetCategoryRequest{
		ID: id,
	})
	if err != nil {
		http.Error(rw, "Invalid URL", http.StatusInternalServerError)
		return
	}

	errs := map[string]string{}
	h.editBookData(rw, res.Category.ID,res.Category.Title, errs)
	return
}

func (h *Handler) editBookData(rw http.ResponseWriter, id int64, title string, errs map[string]string) {

	
	form := Category{
		ID:     id,
		Title:  title,
		Errors: errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "edit_Category.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

//Update
func (h *Handler) Update(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	Id := vars["id"]

	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	id, erre := strconv.ParseInt(Id, 10, 64)

	if erre != nil {
		http.Error(rw, erre.Error(), http.StatusInternalServerError)
		return
	}

	var category Category
	if err := h.decoder.Decode(&category, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	

	if err := category.Validate(); err != nil {
		valError, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range valError {
				vErrs[key] = value.Error()
			}
			h.editBookData(rw, id, category.Title, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err := h.tc.Update(r.Context(), &tpb.UpdateCategoryRequest{
		Category: &tpb.Category{
			ID: id,
			Title: category.Title,
		},
	})
	if err != nil {
		http.Error(rw, "Invalid URL", http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}


// func (h *Handler) bookActive(rw http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	Id := vars["id"]

// 	const updateStatusTodo = `UPDATE books SET status = true WHERE id=$1`
// 	res := h.db.MustExec(updateStatusTodo, Id)

// 	if ok, err := res.RowsAffected(); err != nil || ok == 0 {
// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(rw, r, "/Book/List", http.StatusTemporaryRedirect)
// }

// func (h *Handler) bookDeactivate(rw http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	Id := vars["id"]

// 	const updateStatusTodo = `UPDATE books SET status = false WHERE id=$1`
// 	res := h.db.MustExec(updateStatusTodo, Id)

// 	if ok, err := res.RowsAffected(); err != nil || ok == 0 {
// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(rw, r, "/Book/List", http.StatusTemporaryRedirect)
// }

