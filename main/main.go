package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sykros.store-file-service.net/src/api"
	"sykros.store-file-service.net/src/api/files"
	db "sykros.store-file-service.net/src/pg-database"
	files2 "sykros.store-file-service.net/src/service/files"
	"time"
)

func main() {
	port := int(8000)
	if portStr, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		port = portStr
	}
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	pg, err := db.NewPG(os.Getenv("PG_URL"))
	if err != nil {
		panic(err)
	}
	if err := db.PGMigrate("./migrations", os.Getenv("PG_URL")); err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	fileService := files2.NewFileService().SetupLocalStorage(os.Getenv("FILES_DIR")).SetupDatabase(pg)

	api.NewHandler().SetFileAPI(files.InitFileAPI(fileService)).BaseRouter(router)

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}
	fmt.Println(fmt.Sprintf("Listening on PORT %s", addr))
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

	// shutdown server on os signal
	interrupt := make(chan os.Signal, 1)
	// relay os Interrupt signal to channel
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		return
	}
}
