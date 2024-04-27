package handlers

import (
	helper "github/metasoftonic/coffee-api/internals/helpers"
	"github/metasoftonic/coffee-api/internals/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CoffeList() http.HandlerFunc {
	type d struct {
		Coffees []models.Coffee
	}
	coffeeListHtml, err := helper.ReadHTMLFile("static/coffee.html")
	if err != nil {
		log.Fatal(err)
	}
	templ := template.Must(template.New("").Parse(coffeeListHtml))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.repo.GetCoffees()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templ.Execute(w, d{Coffees: tt})
	}
}

func (h *Handler) CreateCoffee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		size, _ := strconv.ParseFloat(r.FormValue("size"), 64)
		available_units, _ := strconv.ParseInt(r.FormValue("available_units"), 10, 64)

		if err := h.repo.CreateCoffee(&models.Coffee{
			Name:           name,
			Description:    description,
			Size:           size,
			AvailableUnits: available_units,
			PricePerUnit:   price,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/coffees", http.StatusFound)
	}
}
