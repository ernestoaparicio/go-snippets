package interfaces

import (
	"fmt"
	"log"
	"net/http"
)

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

func RunServeMux1() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list))
	//mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/list", db.list)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func RunServeDefaultServeMux() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
