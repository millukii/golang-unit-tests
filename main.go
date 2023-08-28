package main

import (
	"go-tests/handler"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"

)
func main() {

	h:= &handler.Controlador{}

	router := bunrouter.New()
	router.WithGroup("/api/categories", func(group *bunrouter.Group) {
		// /api/categories/:category/items
		group.WithGroup("/:category/items", func(group *bunrouter.Group) {

			    group.GET("/:id", h.GetById)
					group.POST("", h.Create)
					group.PATCH("/:id", h.Update)
					group.DELETE("/:id", h.Delete)
		})

		// /api/categories/archive
		group.WithGroup("/archive", func(group *bunrouter.Group) {})
	})
	log.Println(http.ListenAndServe(":9999", router))
}
