package handlers

import (
	"html/template"
	"net/http"
)

type Item struct {
	Title string
	Body  string
}

func HomeHandler(w http.ResponseWriter, req *http.Request, params map[string]string) {
	body := "Root Page Sucka! :boom: goes the dynamite"
	// for pretty url matching
	// values := fmt.Sprintf("<p> %v </p>", params)
	// for key, value := range params {
	// 	values += fmt.Sprintf("<p> %s := %s </p>", key, value)
	// }

	// query := req.URL.RawQuery
	// query := req.URL.Query  // type Values which is = map[string][]string

	// qOne := req.URL.Query().Get("one")

	i := &Item{Title: "Start Here", Body: body}
	// fmt.Fprintf(w, start+qOne)

	t, _ := template.ParseFiles("templates/root.html")
	t.Execute(w, i)
}
