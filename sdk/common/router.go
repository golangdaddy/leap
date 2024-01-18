package common

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc `json:"-"`
}

func (app *App) ShowRoutes() interface{} {
	return app.routes
}

func (app *App) AddRoute(method, path string, handler gin.HandlerFunc) {
	route := Route{strings.ToUpper(method), path, handler}
	app.routes = append(app.routes, route)
	log.Println("adding route:", route.Method, route.Path)
	app.Gin().Handle(route.Method, route.Path, app.GinHandleCORS, route.Handler)
}

func (app *App) HandleCORS(w http.ResponseWriter, r *http.Request) bool {
	return app.doOptionsHandler("*", w, r)
}

func (app *App) GinHandleCORS(c *gin.Context) {
	app.doOptionsHandler("*", c)
}

func (app *App) doOptionsHandler(pattern string, args ...interface{}) bool {

	if len(args) == 0 {
		panic("OptionsHandler needs request context params")
	}

	methods := "HEAD, POST, GET, OPTIONS, PUT, PATCH, DELETE"
	headers := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

	switch c := args[0].(type) {
	case *gin.Context:

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", methods)
		c.Header("Access-Control-Allow-Headers", headers)

		if c.Request.Method == "OPTIONS" {
			c.Status(204)
			return true
		}

	case http.ResponseWriter:

		w := c
		r := args[1].(*http.Request)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", methods)
		w.Header().Set("Access-Control-Allow-Headers", headers)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return true
		}

	default:

		panic("invalid OptionsHander params")

	}

	return false
}

func (app *App) Serve() error {

	filter := map[string]bool{}
	for _, route := range app.routes {
		if !filter[route.Path] {
			app.Gin().OPTIONS(route.Path, app.GinHandleCORS)
			filter[route.Path] = true
		}
	}
	port := os.Getenv("PORT")
	s := &http.Server{
		Addr:           "0.0.0.0:" + port,
		Handler:        app.Gin(),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("starting server on port:", port)
	err := s.ListenAndServe()
	log.Println("SERVER SHUTTING DOWN...")
	return err
}
