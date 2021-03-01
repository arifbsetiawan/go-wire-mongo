package main

import (
	"fmt"
	"go-wire-mongo/config"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	conf := config.InitConfig()
	database := config.InitDatabase(conf.MongoURI, conf.MongoDB)

	appModule := AppModule()
	authorModule := AuthorModule(database.GetDB())
	bookModule := BookModule(database.GetDB())

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", appModule.Index)
	r.Route("/author", func(r chi.Router) {
		r.Get("/", authorModule.GetIndex)
		r.Post("/", authorModule.PostStore)
		r.Get("/detail", authorModule.GetShow)
		r.Put("/update", authorModule.PutUpdate)
		r.Delete("/delete", authorModule.DeleteDestroy)
	})
	r.Route("/book", func(r chi.Router) {
		r.Get("/", bookModule.GetIndex)
		r.Post("/", bookModule.PostStore)
		r.Get("/detail", bookModule.GetShow)
		r.Put("/update", bookModule.PutUpdate)
		r.Delete("/delete", bookModule.DeleteDestroy)
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf(" %s\t\t|  %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	fmt.Println("\n\nService running at 0.0.0.0:" + conf.PORT)
	http.ListenAndServe(":"+conf.PORT, r)
}
