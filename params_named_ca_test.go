package belajar_golang_httprouter

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterParamsNamed(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/item/:itemid", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemid")
		productName := "Product " + id + " Item " + itemId
		fmt.Fprint(writer, productName)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/products/1/item/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestRouterParamsCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		fmt.Fprint(writer, image)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/images/mountain/tree/grass.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "/mountain/tree/grass.png", string(body))
}
