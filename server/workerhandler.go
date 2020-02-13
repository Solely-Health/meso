package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/meso-org/meso/workers"
)

type workerHandler struct {
	s workers.Service
}

func (h *workerHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/worker", func(chi.Router) {
		r.Post("/", h.registerWorker)
		r.Get("/", h.listWorkers)
		/*
			if we were to add more sub routing:
			r.Route("/pattern", func(chi.Router) {
				r.Verb("/pattern", handlerFunc)
			})
		*/
	})
	r.Get("/ping", h.testPing)
	return r
}

func (h *workerHandler) testPing(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var response = struct {
		Domain string `json:"domain"`
		Ping   string `json:"ping"`
	}{
		Domain: "worker",
		Ping:   "pong",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		encodeError(ctx, err, w)
		return
	}
}

func (h *workerHandler) registerWorker(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()
	// r.ParseForm()
	// fmt.Println("checking the form from r.PrintForm() ", r.FormValue("test"))

	var request struct {
		Test string `json:"test"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("unable to decode json: %v", err)
	}

	fmt.Println("heres the request string ", request.Test)

}

func (h *workerHandler) listWorkers(w http.ResponseWriter, r *http.Request) {

}
