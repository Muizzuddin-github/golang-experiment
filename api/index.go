// package main

package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var server http.Server

func init() {

	router := httprouter.New()

	router.GET("/api",func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		token, err := r.Cookie("token")
		if err == http.ErrNoCookie{
			w.WriteHeader(http.StatusOK)
		  w.Write([]byte("golang okeh"))
		  return
		}


		w.WriteHeader(http.StatusOK)
		w.Write([]byte("golang okeh " + token.Name + "=" + token.Value))
	})

	c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:5173","https://express-cookie-experiment.vercel.app"},
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

