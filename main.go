package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"lenslocked/controllers"
	"lenslocked/templates"
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

	var usersC controllers.Users

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.Route("/galleries", func(r chi.Router) {
		tpl := views.Must(views.ParseFS(templates.FS, "galleries.gohtml"))
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
