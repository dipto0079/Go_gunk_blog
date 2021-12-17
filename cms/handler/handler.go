package handler

import (
	
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	tpb "blog/gunk/v1/category"
	
)

const sessionsName = "cms-session"

type Handler struct {
	templates *template.Template
	decoder   *schema.Decoder
	sess      *sessions.CookieStore
	tc       tpb.CategoryServiceClient
}

func New(decoder *schema.Decoder, sess *sessions.CookieStore ,tc tpb.CategoryServiceClient) *mux.Router {
	h := &Handler{
		decoder: decoder,
		sess:    sess,
		tc:tc,
	}

	h.parseTemplates()
	r := mux.NewRouter()

	r.HandleFunc("/category/create", h.CategoryCreate)
	r.HandleFunc("/category/store", h.CategoryStore)

	

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

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
		"cms/assets/templates/create_Category.html",
	))
}

// func (h *Handler) authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		session, _ := h.sess.Get(r, "library")
// 		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
// 			http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
// 			return
// 		}
// 		next.ServeHTTP(rw, r)
// 	})
// // }
// func (h *Handler) loginMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		session, err := h.sess.Get(r, sessionsName)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		authuser := session.Values["authenticated"]
// 		if authuser != nil {
// 			http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
// 		} else {
// 			next.ServeHTTP(rw, r)
// 		}
// 	})
// }
