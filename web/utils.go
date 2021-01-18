package web

import (
	"fmt"
	"net/http"
)

func serve400(w http.ResponseWriter) {
	fmt.Fprint(w, "Bad Request")
}

func serve404(w http.ResponseWriter) {
	fmt.Fprint(w, "Not Found")
}
