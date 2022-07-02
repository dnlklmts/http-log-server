package app

import (
	"context"
	"fmt"
	"golang-log-server/api"
	"golang-log-server/internals/app/handlers"
	"golang-log-server/internals/app/processors"
	"golang-log-server/internals/app/storages"
	"golang-log-server/internals/cfg"
	"log"
	"net/http"
	"time"
)

type AppServer struct {
	config  *cfg.Config
	storage *storages.LogStorage
	srv     *http.Server
}

func NewAppServer(cfg *cfg.Config) *AppServer {
	return &AppServer{config: cfg}
}

func (server *AppServer) Serve() {
	log.Println("starting app server")

	logStorage := storages.NewLogStorage(server.config.LogFilePath)
	logProcessor := processors.NewLogProcessor(logStorage)
	logHandler := handlers.NewLogHandler(logProcessor)
	routes := api.CreateRoutes(logHandler)

	server.storage = logStorage
	server.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", server.config.Port),
		Handler: routes,
	}
	log.Println("server started")

	err := server.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func (server *AppServer) Shutdown() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("server stopped")
	if err = server.storage.Storage.Close(); err != nil {
		log.Fatalln(err)
	}

	if err = server.srv.Shutdown(ctx); err != nil {
		log.Fatalln("failed to shutdown gracefully:", err)
	}
	log.Println("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
