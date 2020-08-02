package handlers

import (
	"log"
	"net/http"
)

//Products is a http.Handler

type Products struct {
	l *log.Logger
}

//NewProducts creates a products handler with the given logger

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServerHttp is the main entry point for the handler and satisfies the http.Handler
// interface

func (p *Products) ServerHttp(rw http.ResponseWriter, r *http.Request) {
	//handle the request for a list of products

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts return the products from the data store

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Products")

	//fetch the products from the datastore
	lp := data.GetProducts()

	//serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
