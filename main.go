package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/ionnotion/golang-fcm-integration/api/v1"
	"github.com/ionnotion/golang-fcm-integration/config"
	"github.com/ionnotion/golang-fcm-integration/third-party/firebase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	//! LOAD ENV
	if err := godotenv.Load(".env"); err != nil {
		log.Panicf("Failed to load env : %s", err.Error())
	}

	//! DATABASE CONNECTION
	pgDb, err := config.DbPostgresConn()
	if err != nil {
		log.Panicf("Failed to connect to postgres : %s", err.Error())
	}

	pkey := os.Getenv("FCM_NOTIFICATION_PKEY")
	clientCount, _ := strconv.Atoi(os.Getenv("FCM_NOTIFICATION_COUNT"))
	firebase.CloudMessaging = firebase.Initialize(clientCount, pkey, pgDb)

	r := api.Routing{
		//! YOUR ROUTING STRUCT
	}
	e := echo.New()
	api.RegisterRouters(e, &r)

	go func() {
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("RUN_PORT"))); err != nil {
			log.Println("shutting down the server : " + err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
