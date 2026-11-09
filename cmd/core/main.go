package main

import (
	"log"
	"net/http"
	"time"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz" , func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
		"message" : "OK"
		}`))
	})
	server := http.Server{
		Addr: ":8090",
		Handler : mux,
		ReadTimeout : time.Second * 10,
		WriteTimeout : time.Second * 30,
		IdleTimeout : time.Second * 60,
	}
	err := server.ListenAndServe()
	if err != nil{
		log.Fatalf("Server failed to start : %v",err)
	}
	log.Println("Server started on port 8090")
}
