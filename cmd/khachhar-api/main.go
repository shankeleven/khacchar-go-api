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
	"github.com/shankeleven/student-api/internal/http/handlers/khachhar"
)

// some structure basics:
/*
config usually has two config files dev and prod for obvious usecases
we also would have to serialise config as some golang struct for it to be used inside the project by go program  s
internals or pkg in projects is for the packages that are internally used i.e, inside the project and not exposed as such
we had virt-launcher for example

we created different folder for http (/internals/http) because it is very much possible that out app supports more interfaces
other than http so we leave some space for extensibility

*/


func main(){
	 // load config
		fmt.Println("Server started")// better to use logs instead of this to keep record of timestamps and some other features
		slog.Info("Server Started")


		cfg:= config.MustLoad()
		// database setup
		// server routing 

	router:=	http.NewServeMux()
	router.HandleFunc("GET /",khachhar.New()) // this is how we route ,
	// or else we could have written an anonymous function that would handle the request
	// and here in the New() we would pass the dependencies , this would be called dependency injection

	router.HandleFunc("POST /api/khachhars",khachhar.New())

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
		as stated in go documentation
		If no channel is provided, the signals will be discarded. If the channel is full, the signals will be discarded.
  So buffering helps reduce the chance of missing a signal.
		"won't we always block in the unbuffered channel till the channel is listened to?"
		In regular Go channels: yes — sending to an unbuffered channel blocks until a receiver is ready.
		But with os/signal, Go internally receives OS signals asynchronously, and then tries to forward them into your channel.
 This means the Go runtime must deliver the signal to your channel immediately, or else it might get dropped
		Signals are sent by the operating system, and the Go runtime buffers only a limited number of them internally. 
		If your program isn't ready to receive them, they can be:
		blocked(program doesn't shutdown)
		ignored or lost

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
			defer cancel() // read more about this , gd , this is actually pretty interesting


			err:= server.Shutdown(ctx) // but this could go on infinitely keeping the port acquired and wasting resources 

			if(err!=nil){
				slog.Error("failer to shutdown gracefilly : ", slog.String("error", err.Error())) // error is thrown if not completes in 5 seconds
			}
			slog.Info("Shut down successfully")


	fmt.Println("working fine") // not visible because server.ListenAndServe is blocking
}
