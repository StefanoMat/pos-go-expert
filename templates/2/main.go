package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	temp := template.Must(template.New("Curso template").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))

	err := temp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
	// fileServer := http.FileServer(http.Dir("./public"))
	// mux := http.NewServeMux()
	// mux.Handle("/", fileServer)
	// log.Fatal(http.ListenAndServe(":8080", mux))
}
