package mserver

import (
	"encoding/json"
	"net/http"

	"wepa.com/go/mserial"
)

func acmak(w http.ResponseWriter, r *http.Request) {
	mserial.SendChar("1")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ok")
}

func yapmak(w http.ResponseWriter, r *http.Request) {
	mserial.SendChar("0")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ok")
}


func CreateApi() {
	server := http.NewServeMux()
	
	
	server.HandleFunc("/api/ac", acmak)
	server.HandleFunc("/api/yap", yapmak)


	fs := http.FileServer(http.Dir("./public"))

	server.Handle("/", fs)

	http.ListenAndServe(":3333", server)
}
