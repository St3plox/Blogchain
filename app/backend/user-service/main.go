package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	defaultLog "log"

	"github.com/St3plox/Blogchain/app/backend/user-service/handlers"
	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/core/user/userdb"
	"github.com/St3plox/Blogchain/business/web/auth"
	"github.com/St3plox/Blogchain/business/web/v1/debug"
	"github.com/St3plox/Blogchain/foundation/keystore"
	"github.com/St3plox/Blogchain/foundation/logger"
	"github.com/ardanlabs/conf/v3"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var build = "develop"

func main() {
	log := logger.New("GATEWAY - SERVICE")

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
		Auth struct {
			KeysFolder string `conf:"default:zarf/keys/"`
			ActiveKID  string `conf:"default:private_key"`
			Issuer     string `conf:"default:service project"`
		}
		DB struct {
			Uri string `conf:"default:mongodb://localhost:27017"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	const prefix = "AUTH"
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
	// Database Support

	log.Info().Str("status", "startup").Str("uri", cfg.DB.Uri).Msg("initializing database support")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DB.Uri))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return err
	}

	userStore := userdb.NewStore(log, client)
	userCore := user.NewCore(userStore, "http://127.0.0.1:8545/")

	// -------------------------------------------------------------------------
	// Initialize authentication support

	log.Info().Str("status", "startup").Msg("initializing V1 API AUTH support")

	// Simple keystore versus using Vault.
	ks, err := keystore.NewFS(os.DirFS(cfg.Auth.KeysFolder))
	if err != nil {
		return fmt.Errorf("reading keys: %w", err)
	}

	authCfg := auth.Config{
		Log:       log,
		KeyLookup: ks,
	}

	auth, err := auth.New(authCfg)
	if err != nil {
		return fmt.Errorf("constructing auth: %w", err)
	}

	// -------------------------------------------------------------------------
	// Start Debug Service

	log.Info().
		Str("status", "debug v1 router started").
		Str("host", cfg.Web.DebugHost).
		Msg("startup")

	go func() {
		if err := http.ListenAndServe(cfg.Web.DebugHost, debug.StandardLibraryMux()); err != nil {
			log.Error().
				Str("status", "debug v1 router closed").
				Str("host", cfg.Web.DebugHost).
				Err(err).
				Msg("shutdown")

		}
	}()

	// -------------------------------------------------------------------------
	// Start API Service

	log.Info().Msg("initializing V1 API support")
	shutdown := make(chan os.Signal, 1)
	apiMux := handlers.APIMux(handlers.APIMuxConfig{
		Shutdown: shutdown,
		Log:      log,
		Auth:     auth,
		UserCore: userCore,
	})

	errorLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      apiMux,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     defaultLog.New(&errorLogger, "", 0),
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Info().
			Str("status", "api router started").
			Str("host", api.Addr).
			Msg("startup")
		serverErrors <- api.ListenAndServe()
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info().
			Str("status", "shutdown started").
			Str("signal", sig.String()).
			Msg("shutdown")
		defer log.Info().Str("status", "shutdown complete").
			Str("signal", sig.String()).
			Msg("shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			_ = api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
