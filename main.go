package main

import (
	"context"
	"known-anchors/router"
	"known-anchors/service"
	"known-anchors/util/close"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	sc := service.NewServiceContext()
	cl := close.NewClose()
	cl.AddCloseable(sc)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	router := router.NewRouter(sc, cl)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		log.Println("Starting server...")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	cancel()
	log.Println("Closing service gracefully...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := cl.CloseAll(); err != nil {
		log.Fatal(err)
	}
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Service closed")
}
