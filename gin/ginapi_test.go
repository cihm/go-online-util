package ginapi

import (
	//"reflect"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	//"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMainflow(t *testing.T) {
	fmt.Println("begin")
	//Mainflow()
	fmt.Println("enf")
}
func TestSetupRouter(t *testing.T) {
	router := setupRouter()
	fmt.Println("222")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
