package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/iarie/rechallenge/internal"
)

func Run(cfg *Config) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler(cfg.InventoryRepo))

	http.HandleFunc("/place-order/", postOrderHandler(cfg.Packer, cfg.InventoryRepo))

	log.Fatal(http.ListenAndServe(cfg.Addr(), nil))
}

func indexHandler(inventoryRepo internal.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate;")
		w.Header().Set("pragma", "no-cache")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, inventoryRepo.Get())
	}
}

func postOrderHandler(p internal.Packer, inventoryRepo internal.Repository) func(w http.ResponseWriter, r *http.Request) {
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

		o, err := p.Pack(qty, inventoryRepo.Get())
		if err != nil {
			renderError(err, w)
			return
		}

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
