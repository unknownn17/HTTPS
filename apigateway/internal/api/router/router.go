package router

import (
	"api/internal/connections"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	r := http.NewServeMux()

	handler := connections.NewHandler()

	r.HandleFunc("POST /items/create", handler.Create)
	r.HandleFunc("GET /items/{id}", handler.Get)
	r.HandleFunc("GET /items", handler.Gets)
	r.HandleFunc("PUT /items/{id}", handler.Update)
	r.HandleFunc("DELETE /items/{id}", handler.Delete)

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	srv := &http.Server{
		Addr:      ":8081",
		Handler:   r,
		TLSConfig: tlsConfig,
	}
	fmt.Println("Server started on port 8081")
	err := srv.ListenAndServeTLS("./tls/localhost.pem", "./tls/localhost-key.pem")
	r.Handle("/swagger/", httpSwagger.WrapHandler)
	logger.Error(err.Error())
	os.Exit(1)
}
