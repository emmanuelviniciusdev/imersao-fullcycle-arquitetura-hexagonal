package server

import (
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/adapters/web/handler"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	ProductService application.ProductServiceInterface
}

func NewWebServer(productService application.ProductServiceInterface) *WebServer {
	return &WebServer{ProductService: productService}
}

func (w *WebServer) Serve() {
	r := mux.NewRouter()

	n := negroni.New(negroni.NewLogger())

	http.Handle("/", r)

	handler.MakeProductHandlers(r, n, w.ProductService)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
