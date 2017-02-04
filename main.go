package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func embedString(str string) string {
	return fmt.Sprintf("")
}

type example struct {
	Name    string
	Content string
}

var Examples []example

func loadExamples() {
	files, _ := ioutil.ReadDir("./examples")
	for _, f := range files {
		b, err := ioutil.ReadFile("./examples/" + f.Name()) // just pass the file name
		if err != nil {
			fmt.Print(err)
			continue
		}

		ex := example{
			Name:    f.Name(),
			Content: string(b),
		}

		Examples = append(Examples, ex)
	}
}

func main() {

	loadExamples()

	http.Handle("/", http.FileServer(http.Dir("./html")))

	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {

		t, err := template.ParseFiles("index.html")

		if err != nil {
			panic(err)
		}

		err = t.ExecuteTemplate(w, "index.html", Examples)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":3000", nil)
}
