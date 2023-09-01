# golang-unit-tests

ejemplos de test unitarios en go usando bun router

### Requisitos

- Instalar go y configurar las rutas necesarias como variables de entorno
- Instalar vs studio code o editor de preferencia
- Instalar go tools y el plugin de go
- Iniciar repositorio

## Paso a paso

1.  Iniciar modulo

           go mod init "nombre del modulo"

2.  Para crear nuestro go.sum(archivo que registrará los requerimientos y dependencias) go mod tidy
3.  Crear archivo main.go

```go
   package main

   func main() {

   }

```

4.  Vamos a descargar el enrutador http que usaremos

           go get github.com/uptrace/bunrouter

5.  Agregamos el import al main

```go
import "github.com/uptrace/bunrouter"
```

6. Creamos la instancia al router que nos proporciona esta librería

```go
router := bunrouter.New()
```

7. Ver documentación oficial [Creando un go router ][doc-bun-router]

8. Según la documentación oficial nos recomiendan ordenar los endpoint en grupos

```go
router.WithGroup("/api/categories", func(group *bunrouter.Group) {
    // /api/categories/:category/items
    group.WithGroup("/:category/items", func(group *bunrouter.Group) {})

    // /api/categories/archive
    group.WithGroup("/archive", func(group *bunrouter.Group) {})
})
```

9. Agregaremos el archivo handler.go

```go
package main

import (
	"net/http"

	"github.com/uptrace/bunrouter"

)

type handler struct {
}

func (h handler)GetById(w http.ResponseWriter, req bunrouter.Request) error {

	return nil
}

func (h handler)Create(w http.ResponseWriter, req bunrouter.Request) error {

	return nil
}

func (h handler)Delete(w http.ResponseWriter, req bunrouter.Request) error {

	return nil
}

func (h handler)Update(w http.ResponseWriter, req bunrouter.Request) error {

	return nil
}
```

10. Actualizamos el main para incluir el import del handler y especificar los endpoint - El método, la ruta y los parametros

```go
package main

import (
	"go-tests/handler"

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
}
```

11. En el handler del metodo getbyid agregamos para imprimir en consola detalles de la request

```go

    fmt.Println(req.Method, req.Route(), req.Params().Map())

```

12. Podemos usar curl o armar un postman ((o cliente http a elección))

`curl --location 'http://localhost:9999/api/categories/category/items/1' `

13. Agregar la línea en main.go donde levantaremos el servidor

```go
log.Println(http.ListenAndServe(":9999", router))
```

14. Ejecutar con go run main.go y realizar la request

15. Debiesemos ver en consola

GET /api/categories/:category/items/:id map[category:category id:1]

16. Vamos a modificar el handler agregando un slice de categorias

```go
type Category struct{
	Id int
	Name string
}

var	categories = []Category{Category{
		Id: 1,
		Name: "News",
}}

func (h Controlador)GetById(w http.ResponseWriter, req bunrouter.Request) error {
 // req embeds *http.Request and has all the same fields and methods
    fmt.Println(req.Method, req.Route(), req.Params().Map())

		id := req.Params.ByName("id")
		cat :=""
		for i := range categories {
			if fmt.Sprint(categories[i].Id) ==id {
						// Found!
						cat = categories[i].Name
				}
		}
		if cat ==""{
			w.WriteHeader(404)
			return nil
		}
			//caso feliz
	w.WriteHeader(200)
	return nil
}
```

17. Agregamos el archivo handler_test.go

```go
package handler_test

import "testing"

func TestHandlerGetById(t *testing.T) {

}
func TestHandlerCreate(t *testing.T) {

}
func TestHandlerGetDelete(t *testing.T) {

}
func TestHandlerGetUpdate(t *testing.T) {

}
```

18. Agregar libreria para asserts

go get "github.com/stretchr/testify/assert"

19. Actualizamos con los test

```go
var router = bunrouter.New()

func TestHandlerGetById(t *testing.T) {
c :=&handler.Controlador{}

router.GET("/api/categories/category/items/:id", c.GetById)

	t.Run("Happy Path 200 status code", func(t *testing.T) {
		w := httptest.NewRecorder()
		baseURL, _ := url.Parse("/api/categories/category/items/1")
    params := url.Values{}
    params.Add("id", "1")
    baseURL.RawQuery = params.Encode()

		req, _ := http.NewRequest(http.MethodGet, baseURL.String(), nil)
		err := router.ServeHTTPError(w, req)
	  assert.Equal(t, 200, w.Code)
		require.NoError(t, err)

	})

	t.Run("UnHappy Path 404 status code", func(t *testing.T) {
		w := httptest.NewRecorder()
		baseURL, _ := url.Parse("/api/categories/category/items/2")
    params := url.Values{}
    params.Add("id", "2")
    baseURL.RawQuery = params.Encode()

		req, _ := http.NewRequest(http.MethodGet, baseURL.String(), nil)
		err := router.ServeHTTPError(w, req)
	  assert.Equal(t, 404, w.Code)
		require.NoError(t, err)

	})
}
```

20. Agregamos service.go con la estructura base

```go
package service

type CountriesCategory struct {
	Country  string
	Category int
	Enabled  bool
}

type Service interface {
	GetById(id int, country string) (string, error)
}

type svc struct {
}

func NewService() Service {

	return &svc{}
}

func (s svc) GetById(id int, country string) (string, error) {

	return "", nil
}
```

21. Instalar mockery go install github.com/vektra/mockery/v2@v2.33.0

22. Crear carpeta models y archivo category.go

23. Crear carpeta repository y archivo category.go
24. Vamos a crear una capa de repositorio donde tengamos una separacion de logica de operaciones sobre un repositorio de datos estándar sql a la implementación concreta del Sistema Gestor de Base de dato.

```go
package repository

import (
	"database/sql"
	"go-tests/models"
)

type CategoryRepo interface {
	GetById(id int, country string) (models.CountriesCategory, error)
}

type catRepo struct{
	db *sql.DB
}
func NewCategoryRepo(db *sql.DB) CategoryRepo{
	return &catRepo{}
}

func (c catRepo) 	GetById(id int, country string) (models.CountriesCategory, error){

	return models.CountriesCategory{}, nil

}
```

[doc-bun-router]: https://bunrouter.uptrace.dev/guide/golang-router.html#creating-go-router
