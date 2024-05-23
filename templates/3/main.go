package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	temp := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := temp.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 60},
		{"Python", 45},
	})
	if err != nil {
		panic(err)
	}
	// fileServer := http.FileServer(http.Dir("./public"))
	// mux := http.NewServeMux()
	// mux.Handle("/", fileServer)
	// log.Fatal(http.ListenAndServe(":8080", mux))
}
