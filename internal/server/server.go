package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-keycloak-example1/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(r *mux.Router) {
	// create a server object
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler:      r,
		WriteTimeout: time.Hour,
		ReadTimeout:  time.Hour,
	}

	fmt.Println(fmt.Sprintf("Server is started and listening on port : %s", config.Port))

	server.ListenAndServe()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
