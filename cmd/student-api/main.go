package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/shankeleven/student-api/internal/config"
)

// some structure basics:
/*
config usually has two config files dev and prod for obvious usecases
we also would have to serialise config as some golang struct for it to be used inside the project by go program  s
internals or pkg in projects is for the packages that are internally used i.e, inside the project and not exposed as such
we had virt-launcher for example

*/


func main(){
	 // load config
		fmt.Println("Server started")


		cfg:= config.MustLoad()
		// database setup
		// server routing 

	router:=	http.NewServeMux()
	router.HandleFunc("GET /",func(w http.ResponseWriter , r *http.Request){
		w.Write([]byte("Pranaam!"))
	})

		// setup the server
		server:= http.Server{
			Addr: cfg.Addr,
			Handler: router,
		}


		err:= server.ListenAndServe()
		if(err!=nil){
			log.Fatal("server not started because :", err.Error())
		}

	fmt.Println("working fine") // not visible because server.ListenAndServe is blocking
}
