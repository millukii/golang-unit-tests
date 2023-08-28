package handler_test

import (
	"go-tests/handler"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bunrouter"

)

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
func TestHandlerCreate(t *testing.T) {

}
func TestHandlerGetDelete(t *testing.T) {

}
func TestHandlerGetUpdate(t *testing.T) {

}