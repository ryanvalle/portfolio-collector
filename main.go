package router

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", controllers.base);
}
