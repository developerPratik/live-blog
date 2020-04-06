package main

import (
	"log"
	"net/http"
	"os"
	"real_time_go/config"
	"real_time_go/database"
	"real_time_go/route"
)

func init() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	db := database.Get()
	err := db.CreateConnection()

	if err != nil {
		log.Fatal(err)
	}

	hub := route.NewHub(db.Session)

	go hub.Start()

	log.Println("connected to database")

	errorChannel := make(chan error)
	appRouter := route.GetRouter()

	appRouter.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		route.ServeWebSocket(hub, writer, request)
	})

	srv := http.Server{
		Addr:              os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"),
		Handler:           appRouter,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errorChannel <- err
		}
	}()

	serverErr := <-errorChannel

	log.Fatal(serverErr)

}
