package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/iarie/rechallenge/data"
	"github.com/iarie/rechallenge/internal"
)

func Run(cfg *Config) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler(cfg.InventoryRepo))

	http.HandleFunc("/place-order/", postOrderHandler(cfg.Packer, cfg.InventoryRepo))

	http.HandleFunc("/add-package/", postPackageHandler(cfg.InventoryRepo))

	http.HandleFunc("/delete-package/", deletePackageHandler(cfg.InventoryRepo))

	log.Fatal(http.ListenAndServe(cfg.Addr(), nil))
}

func indexHandler(inventoryRepo internal.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate;")
		w.Header().Set("pragma", "no-cache")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		tmpl := template.Must(template.ParseFiles("index.html"))

		data := struct {
			Version  string
			Packages []data.Package
		}{
			Version:  os.Getenv("APP_VERSION"),
			Packages: inventoryRepo.Get(),
		}

		tmpl.Execute(w, data)

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

func postPackageHandler(inventoryRepo internal.Repository) func(w http.ResponseWriter, r *http.Request) {
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

		err = inventoryRepo.New(qty)

		if err != nil {
			renderError(err, w)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/packs.html"))
		tmpl.Execute(w, inventoryRepo.Get())
	}
}

func deletePackageHandler(inventoryRepo internal.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("deleteOrderHandler")

		// Validate
		parts := strings.Split(r.URL.Path, "/")

		// string to int
		size, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			renderError(err, w)
			return
		}

		err = inventoryRepo.Delete(size)

		if err != nil {
			renderError(err, w)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/packs.html"))
		tmpl.Execute(w, inventoryRepo.Get())
	}
}

func renderError(err error, w http.ResponseWriter) {
	log.Println("Error: ", err)
	tmpl := template.New("t")
	tmpl.Parse(fmt.Sprintf("Error: %v", err.Error()))
	tmpl.Execute(w, err)
}
