package main

import (
	"context"
	"fmt"
	"log"
	"log/slog" // for structured logs
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		go func(){
			for i:=0;i>-1;i++{
				fmt.Println(i);
			}
		}()
	})

		// setup the server
		server:= http.Server{
			Addr: cfg.Addr,
			Handler: router,
		}

		done := make(chan os.Signal,1) //  buffer size one for some reason?

		/*
		this done channel is preferred to be buffered because:
		If unbuffered and the signal arrives before you're listening, the signal will be lost or your program might block waiting for the receiver.
		If buffered, the signal can be safely queued even if your goroutine isn't ready to receive it immediately.
		*/


		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // to notify about the signals in this channel

		go func(){
			err:= server.ListenAndServe() // this is blocking ofcourse
			if(err!=nil){
				log.Fatal(err.Error()) // this shuts down the service immediately which is anything but graceful 
				// as there could be tasks that are ongoing so we shall not force shut them so we listenandserve in a different go routine
			}

		}()

		<- done; // to wait for the server to run and make sure main thread is not finished and stopped
			slog.Info("shutting down the server")


			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // to give 5 seconds to gracefully shutdown
			defer cancel()


			err:= server.Shutdown(ctx) // but this could go on infinitely keeping the port acquired and wasting resources 

			if(err!=nil){
				slog.Error("failer to shutdown gracefilly : ", slog.String("error", err.Error())) // error is thrown if not completes in 5 seconds
			}
			slog.Info("Shut down successfully")


	fmt.Println("working fine") // not visible because server.ListenAndServe is blocking
}
