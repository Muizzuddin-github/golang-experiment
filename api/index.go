// package main
package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var server http.Server

func init() {

	router := httprouter.New()

	router.GET("/",func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		token, _ := r.Cookie("token")

		fmt.Println(token)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("golang okeh " + token.Name + "="+ token.Value))
	})

	c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:5173"},
    AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		
})
	server = http.Server{
		Handler:  c.Handler(router),
		Addr: "localhost:" + "5000" ,
	}


	// fmt.Println("server is listening")
	// server.ListenAndServe()
	
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server.Handler.ServeHTTP(w, r)
}

