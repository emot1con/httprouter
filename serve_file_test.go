package belajar_golang_httprouter

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestRouterServeFile(t *testing.T) {
	router := httprouter.New()

	directory, _ := fs.Sub(resources, "resources/")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Hello HttpRouter", string(body))
}
