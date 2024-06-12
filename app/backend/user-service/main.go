package main

import (
	"os"

	"github.com/St3plox/Blogchain/foundation/logger"
	"github.com/rs/zerolog"
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
	
	
	
	return nil
}
