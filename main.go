package main

import (
	"azuki-server/api"
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	config := viper.New()
	config.AutomaticEnv()

	var envPath string
	flag.StringVar(&envPath, "env", "", "path to .env file to use")
	flag.Parse()

	var err error
	if envPath != "" {
		err = godotenv.Load(envPath)
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		log.Printf("failed to load .env file %s\n", err)
	}

	logger := log.Default()
	if config.GetString("ENVIRONMENT") == "dev" {
		logger.SetLevel(log.DebugLevel)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := run(ctx, os.Stdout, logger, config, os.Args[1:]); err != nil {
		log.Fatal("failed to run", "err", err)
	}
}

func run(ctx context.Context, _ io.Writer, logger *log.Logger, config *viper.Viper, _ []string) error {
	server := api.NewServer(config, logger)

	config.SetDefault("PORT", 8484)
	port := config.GetInt("PORT")

	go func() {
		log.Info("Listening", "port", port)
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			logger.Fatal("Failed to listen", "err", err)
		}

		log.Info("Serving")
		err = server.Serve(ln)
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("Received error from http server", "err", err)
		}
	}()

	<-ctx.Done()

	timeout := time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return server.Shutdown(ctx)
}
