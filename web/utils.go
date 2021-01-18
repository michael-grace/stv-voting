package web

import (
	"fmt"
	"html/template"
	"net/http"
)

func generateVoteList(n int) (toRtn []string) {
	for i := 0; i < n; i++ {
		toRtn = append(toRtn, fmt.Sprintf("%v", i+1))
	}
	return
}

func serveTemplate(w http.ResponseWriter, path string, data interface{}, code int) error {
	w.WriteHeader(code)

	tmpl, err := template.New("base.html").ParseFiles("templates/base.html", fmt.Sprintf("templates/%s", path))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error Creating Templates")
	}

	return tmpl.Execute(w, data)

}

func serve400(w http.ResponseWriter) {
	if err := serveTemplate(w, "error.html",
		struct {
			Code int
			Text string
		}{
			Code: 400,
			Text: "Bad Request",
		}, 400); err != nil {
		fmt.Println(err)
	}
}

// Serve404 gives a 404 Not Found page
func Serve404(w http.ResponseWriter) {
	if err := serveTemplate(w, "error.html",
		struct {
			Code int
			Text string
		}{
			Code: 404,
			Text: "Not Found",
		}, 404); err != nil {
		fmt.Println(err)
	}
}

func serve403(w http.ResponseWriter) {
	if err := serveTemplate(w, "error.html",
		struct {
			Code int
			Text string
		}{
			Code: 403,
			Text: "Forbidden",
		}, 403); err != nil {
		fmt.Println(err)
	}
}
