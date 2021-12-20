package handler

import (
	tpc "blog/gunk/v1/category"
	"fmt"
	"net/http"
)


func (h *Handler) Home(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.tc.GetAllData(ctx, &tpc.GetAllDataCategoryRequest{})
	fmt.Println("###############error#############")
	fmt.Printf("%+v", err)
	// fmt.Println("###############error#############")
	// fmt.Println("###############res#############")
	// fmt.Printf("%+v", res)
	// fmt.Println("###############res#############")

	if err := h.templates.ExecuteTemplate(rw, "index.html", res); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
