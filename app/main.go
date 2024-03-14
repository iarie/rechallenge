package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/iarie/rechallenge/data"
	"github.com/iarie/rechallenge/internal"
)

func Run(cfg *Config) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/place-order/", postOrderHandler(cfg.Packer))

	log.Fatal(http.ListenAndServe(cfg.Addr(), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	tmpl.Execute(w, nil)
}

func postOrderHandler(p internal.Packer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("postOrderHandler")

		// Validate
		r.ParseForm()

		// string to int
		qty, err := strconv.Atoi(r.PostForm.Get("qty"))
		if err != nil {
			renderError(err, w)
			return
		}

		// Process
		pkg_50 := data.Package{Sku: "xxxx0200", Size: 50}
		pkg_100 := data.Package{Sku: "xxxx0200", Size: 100}
		inventory := []data.Package{
			pkg_50,
			pkg_100,
		}
		o := p.Pack(qty, inventory)

		tmpl := template.Must(template.ParseFiles("templates/order.html"))
		tmpl.Execute(w, o)
	}
}

func renderError(err error, w http.ResponseWriter) {
	log.Println("Error: ", err)
	tmpl := template.New("t")
	tmpl.Parse(fmt.Sprintf("Error: %v", err.Error()))
	tmpl.Execute(w, err)
}
