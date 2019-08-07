package server

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"

	"go.uber.org/zap"

	"github.com/marcofranssen/gothermostat/storage"
)

var store *storage.Store

// Server provides an http.Server
type Server struct {
	logger *zap.Logger
	*http.Server
}

// NewServer creates a new instance of a server and configures the routes
func NewServer(cfg *viper.Viper, storage *storage.Store, logger *zap.Logger) (*Server, error) {
	store = storage
	r := &http.ServeMux{}
	r.Handle("/", http.FileServer(http.Dir("./web/build")))
	r.HandleFunc("/ping", ping)
	r.HandleFunc("/api/thermostat-data", api)

	errorLog, _ := zap.NewStdLogAt(logger, zap.ErrorLevel)
	srv := http.Server{
		Addr:         cfg.GetString("listenAddr"),
		Handler:      r,
		ErrorLog:     errorLog,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{logger, &srv}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *Server) Start() {
	srv.logger.Info("Starting server...")
	defer srv.logger.Sync()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.logger.Fatal("Could not listen on", zap.String("addr", srv.Addr), zap.Error(err))
		}
	}()
	srv.logger.Info("Server is ready to handle requests", zap.String("addr", srv.Addr))
	srv.gracefullShutdown()
}

func (srv *Server) gracefullShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	srv.logger.Info("Server is shutting down", zap.String("reason", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		srv.logger.Fatal("Could not gracefully shutdown the server", zap.Error(err))
	}
	srv.logger.Info("Server stopped")
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func returnServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func jsonWrite(w http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		returnServerError(w, err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func api(w http.ResponseWriter, r *http.Request) {
	data, err := store.GetTemperatureData()
	if err != nil {
		returnServerError(w, err)
	} else {
		jsonWrite(w, data)
	}
}
