package main

import (
	"net/http"
	"github.com/caiwp/impala-web/router"
	"github.com/caiwp/impala-web/router/middleware"
)

func main() {
	handler := router.Load(
		middleware.Version,
	)
	http.ListenAndServe(":3000", handler)
}
