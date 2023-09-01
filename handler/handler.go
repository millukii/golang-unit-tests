package handler

import (
	"fmt"
	"go-tests/models"
	"net/http"
	"strings"

	"github.com/uptrace/bunrouter"

)

type Controlador struct {
}




func (h Controlador)GetById(w http.ResponseWriter, req bunrouter.Request) error {
 // req embeds *http.Request and has all the same fields and methods
    fmt.Println(req.Method, req.Route(), req.Params().Map())
var	categories = []models.Category{models.Category{
		Id: 1,
		Name: "News",
}}
		id := req.Param("id")
		fmt.Println("------------------------------->id ", id)
		for i := range categories {
			if fmt.Sprint(categories[i].Id) ==id {
						// Found!
						cat := categories[i].Name
							w.WriteHeader(200)
							raw := strings.NewReader(cat)
							raw.WriteTo(w)
				}
		}

			w.WriteHeader(404)
			return nil
}

func (h Controlador)Create(w http.ResponseWriter, req bunrouter.Request) error {

	return nil
}

func (h Controlador)Delete(w http.ResponseWriter, req bunrouter.Request) error {



	return nil
}

func (h Controlador)Update(w http.ResponseWriter, req bunrouter.Request) error {

	return nil
}
