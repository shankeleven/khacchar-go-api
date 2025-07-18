package khachhar

import (
	"log/slog"
	"net/http"
)

// this file would ofcourse return handlerfunction

func New() http.HandlerFunc{

			slog.Info("Adding a khachhar") 
	  return func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Pranaam!")) }
}

   