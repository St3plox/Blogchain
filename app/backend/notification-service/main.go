package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

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

	fmt.Println(cfg)

	return nil

}
