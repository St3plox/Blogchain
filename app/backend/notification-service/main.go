package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	gomail "gopkg.in/mail.v2"

	"github.com/St3plox/Blogchain/business/core/email"
	"github.com/St3plox/Blogchain/business/web/broker/consumer"
	"github.com/St3plox/Blogchain/foundation/logger"
	"github.com/ardanlabs/conf/v3"
	"github.com/rs/zerolog"
)

var build = "develop"

const cfgPath = "app/backend/notification-service/config.json"

func main() {
	log := logger.New("NOTIFICATION - SERVICE")

	if err := run(log); err != nil {
		log.Error().Err(err).Msg("startup")
		os.Exit(1)
	}
}

func run(log *zerolog.Logger) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS
	log.Info().Str("startup", "GOMAXPROCS").Int("GOMAXPROCS", runtime.
		GOMAXPROCS(0)).
		Str("BUILD", build).
		Msg("startup")

	// -------------------------------------------------------------------------
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
			ShutdownTimeout time.Duration `conf:"default:20s,mask"`
			APIHost         string        `conf:"default::3000"`
			DebugHost       string        `conf:"default::4000"`
		}
		Email struct {
			AdminKey   string `json:"admin_key"`
			AdminEmail string `json:"admin_email"`
		}
		LikeConsumer struct {
			RetryDelay time.Duration `conf:"default:5s"`
			Topic      string        `conf:"default:likes"`
			Address    string        `conf:"default:localhost:9092"`
			Group      string        `conf:"default:blogchain-group"`
			BufferSize int           `conf:"default:8"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	// Load the configuration from a file
	cfgFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}

	err = json.Unmarshal(cfgFile, &cfg)
	if err != nil {
		return fmt.Errorf("error unmarshal: %w", err)
	}

	prefix := "CFG"
	help, err := conf.Parse(prefix, &cfg)

	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// -------------------------------------------------------------------------
	// App Starting

	log.Info().Str("version", build).Msg("starting service")
	defer log.Info().Msg("shutdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info().Str("config", out).Msg("startup")

	// -------------------------------------------------------------------------
	// initializing notifications support

	d := gomail.NewDialer("smtp.gmail.com", 587, cfg.Email.AdminEmail, cfg.Email.AdminKey)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	emailCore := email.NewCore(cfg.Email.AdminEmail, d)

	// -------------------------------------------------------------------------
	// initializing consumer support


	kafkaAdress := os.Getenv("KAFKA_ADDRESS")
	if kafkaAdress == "" {
		kafkaAdress = cfg.LikeConsumer.Address
	}


	likeConsumer, err := consumer.NewLikeConsumer(
		kafkaAdress,
		cfg.LikeConsumer.Group,
		cfg.LikeConsumer.Topic,
		log,
		runtime.GOMAXPROCS(0),
		time.Microsecond*10,
	)

	likeController := consumer.New(likeConsumer, time.Second, log, emailCore)

	serverErrors := make(chan error, 1)
	go func() {
		log.Info().
			Str("status", "notification service started").
			Msg("startup")
		serverErrors <- likeController.ListenForEvents(context.Background())
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

		// case sig := <-shutdown:
		// 	log.Info().
		// 		Str("status", "shutdown started").
		// 		Str("signal", sig.String()).
		// 		Msg("shutdown")
		// 	defer log.Info().Str("status", "shutdown complete").
		// 		Str("signal", sig.String()).
		// 		Msg("shutdown")

		// 	ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		// 	defer cancel()


	}

	// return nil
}
