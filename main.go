package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"notificationSubscriber/logger"
	"notificationSubscriber/subscribers"
	"notificationSubscriber/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		zapLog := logger.GetForFile("startup-errors")
		zapLog.Error("No .env file found", zap.Error(err))
	}
}

func main() {
	// Define a timeout duration
	const timeout = 30 * time.Second
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)

	// Create instances of EmailSubscriber and SmsSubscriber

	client, err := pubsub.NewClient(ctx, os.Getenv("PROJECT_ID"), option.WithCredentialsFile(os.Getenv("KEY_FILE_PATH")))
	if err != nil {
		zapLog := logger.GetForFile("startup-errors")
		zapLog.Error("Error creating Pub/Sub client", zap.Error(err))
		return
	}

	defer func(client *pubsub.Client) {
		err := client.Close()
		if err != nil {
			zapLog := logger.GetForFile("startup-errors")
			zapLog.Error("Error closing client", zap.Error(err))
		}
	}(client)

	subscriber := utils.NewSubscriber(ctx)

	emailSub := subscriber.CreateSubscription(os.Getenv("SUBSCRIPTION_ID_EMAIL"), client)

	smsSub := subscriber.CreateSubscription(os.Getenv("SUBSCRIPTION_ID_SMS"), client)

	adminEmailSub := subscriber.CreateSubscription(os.Getenv("SUBSCRIPTION_ID_ADMIN_EMAIL"), client)

	go subscribers.ProcessEmail(emailSub, ctx)
	go subscribers.ProcessSMS(smsSub, ctx)
	go subscribers.AdminProcessEmail(adminEmailSub, ctx)

	// Set up HTTP server

	// srv := &http.Server{Addr: ":3000"}
	srv := &http.Server{Addr: ":" + os.Getenv("HTTP_PORT")}
	// Set up HTTP server
	http.HandleFunc("/v1/health", utils.HealthCheckHandler)
	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Wait for termination signal
	<-sig
	log.Println("Received termination signal. Shutting down...")

	// Graceful shutdown of HTTP server
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server Shutdown Failed:%+v", err)
	} else {
		log.Println("HTTP server gracefully stopped")
	}

	// The select {} at the end of the main function is a common idiom in Go for blocking indefinitely.
	// This is useful here to keep the main goroutine running while other goroutines listen for messages.
	//select {}

}
