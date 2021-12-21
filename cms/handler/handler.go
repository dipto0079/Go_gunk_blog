package handler

import (
	"net/http"
	"text/template"

	tpb "blog/gunk/v1/blog"
	tpc "blog/gunk/v1/category"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

const sessionsName = "cms-session"

type Handler struct {
	templates *template.Template
	decoder   *schema.Decoder
	sess      *sessions.CookieStore
	tc        tpc.CategoryServiceClient
	tb        tpb.BlogServiceClient
}

func New(decoder *schema.Decoder, sess *sessions.CookieStore, tc tpc.CategoryServiceClient, tb tpb.BlogServiceClient) *mux.Router {
	h := &Handler{
		decoder: decoder,
		sess:    sess,
		tc:      tc,
		tb:      tb,
	}

	h.parseTemplates()
	r := mux.NewRouter()
	r.HandleFunc("/", h.Home)
	r.HandleFunc("/category/list", h.CategoryList)
	r.HandleFunc("/category/create", h.CategoryCreate)
	r.HandleFunc("/category/store", h.CategoryStore)
	r.HandleFunc("/category/{id:[0-9]+}/delete", h.Delete)
	r.HandleFunc("/category/{id:[0-9]+}/edit", h.Edit)
	r.HandleFunc("/category/{id:[0-9]+}/update", h.Update)


	r.HandleFunc("/blog/create", h.BlogCreate)
	r.HandleFunc("/blog/store", h.BlogStore)
	r.HandleFunc("/blog/{id:[0-9]+}/update", h.BlogUpdate)
	r.HandleFunc("/blog/list", h.BlogList)
	r.HandleFunc("/blog/{id:[0-9]+}/single", h.BlogSingle)
	r.HandleFunc("/blog/{id:[0-9]+}/delete", h.BlogDelete)
	r.HandleFunc("/blog/{id:[0-9]+}/edit", h.BlogEdit)

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./"))))

	r.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := h.templates.ExecuteTemplate(rw, "404.html", nil); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	return r
}

func (h *Handler) parseTemplates() {
	h.templates = template.Must(template.ParseFiles(
		"cms/assets/templates/404.html",
		"cms/assets/templates/index.html",
		"cms/assets/templates/create_Category.html",
		"cms/assets/templates/list_category.html",
		"cms/assets/templates/edit_Category.html",
		"cms/assets/templates/create_blog.html",
		"cms/assets/templates/list-blog.html",
		"cms/assets/templates/edit_Blog.html",
		"cms/assets/templates/single_blog.html",
	))
}
