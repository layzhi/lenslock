package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"lenslocked/controllers"
	"lenslocked/views"
)

type Router struct{}

func GalleriesCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		galleriesID := chi.URLParam(r, "galleriesID")
		fmt.Println(galleriesID)
	})
}

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.Route("/galleries", func(r chi.Router) {
		tpl = views.Must(views.Parse(filepath.Join("templates", "galleries.gohtml")))
		r.Get("/", controllers.StaticHandler(tpl))
		// Subrouters:
		r.Route("/{galleriesID}", func(r chi.Router) {
			r.Use(GalleriesCtx)
			r.Get("/", controllers.StaticHandler(tpl))
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
