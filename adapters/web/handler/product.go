package handler

import (
	"encoding/json"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/adapters/web/dto"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(negroni.Wrap(getProduct(service)))).Methods("GET", "OPTIONS")
	r.Handle("/product", n.With(negroni.Wrap(createProduct(service)))).Methods("POST", "OPTIONS")
	r.Handle("/product/{id}/enable", n.With(negroni.Wrap(enableProduct(service)))).Methods("POST", "OPTIONS")
	r.Handle("/product/{id}/disable", n.With(negroni.Wrap(disableProduct(service)))).Methods("POST", "OPTIONS")
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestVars := mux.Vars(r)

		productId := requestVars["id"]

		product, err := service.Get(productId)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = service.Disable(product)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(ResponseJson(err.Error()))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestVars := mux.Vars(r)

		productId := requestVars["id"]

		product, err := service.Get(productId)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = service.Enable(product)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(ResponseJson(err.Error()))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var productDTO dto.ProductDTO

		err := json.NewDecoder(r.Body).Decode(&productDTO)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(ResponseJson(err.Error()))
			return
		}

		product, err := service.Create(productDTO.Name, productDTO.Price)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(ResponseJson(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestVars := mux.Vars(r)

		productId := requestVars["id"]

		product, err := service.Get(productId)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
