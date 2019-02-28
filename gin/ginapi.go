package ginapi

import (

	//"go-online-util/reflectinvoke"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Mainflow() {
	//r := setupRouter()
	//r.Run(":8080")
	router := setupRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8000),
		Handler:        router,
		ReadTimeout:    10,
		WriteTimeout:   10,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	fmt.Println("111")
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
