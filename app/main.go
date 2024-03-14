package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Config struct {
	port int
}

func NewConfig(port int) *Config {
	return &Config{port: port}
}

func (ac *Config) Addr() string {
	return fmt.Sprintf(":%v", ac.port)
}

func Run(cfg *Config) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/solve/", solveHandler)

	log.Fatal(http.ListenAndServe(cfg.Addr(), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	tmpl.Execute(w, nil)
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	qty := r.PostForm.Get("qty")

	tmpl := template.New("t")
	tmpl.Parse("<li>1 - 250</li><li>2 - 250</li><li>3 - 250</li>")
	tmpl.Execute(w, nil)

	log.Printf("SOLVE: %v", qty)
}
