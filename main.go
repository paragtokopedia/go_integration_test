package main


import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	w.Write([]byte(`{"alive": true,"version":1}`))
}

func main() {
	log.Println("Starting Server at 80 ")
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", r))
}