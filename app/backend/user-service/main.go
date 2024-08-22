package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	defaultLog "log"

	"github.com/St3plox/Blogchain/app/backend/user-service/handlers"
	"github.com/St3plox/Blogchain/business/core/media"
	"github.com/St3plox/Blogchain/business/core/media/mediadb"
	"github.com/St3plox/Blogchain/business/core/post"
	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/core/user/userdb"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger"

	contractAuth "github.com/St3plox/Blogchain/foundation/blockchain/auth"
	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/St3plox/Blogchain/foundation/web"

	"github.com/St3plox/Blogchain/business/web/auth"
	"github.com/St3plox/Blogchain/business/web/v1/debug"
	"github.com/St3plox/Blogchain/foundation/blockchain"
	"github.com/St3plox/Blogchain/foundation/blockchain/contract"
	"github.com/St3plox/Blogchain/foundation/keystore"
	"github.com/St3plox/Blogchain/foundation/logger"
	"github.com/ardanlabs/conf/v3"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	gorillaHandlers "github.com/gorilla/handlers"

	_ "github.com/St3plox/Blogchain/docs"
)

var build = "develop"

// @title Blogchain API
// @version 1.0
// @description This is a sample server Blogchain server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT License
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey jwtToken
// @in header
// @name Authorization
func main() {
	log := logger.New("USER - SERVICE")

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
		Redis struct {
			Url string `conf:"default:redis://@localhost:6379"`
		}
		Media struct {
			MaxFileSizeMb string `conf:"default:10"`
		}
		ETH struct {
			Rawurl   string `conf:"default:http://127.0.0.1:8545"`
			GasLimit uint64 `conf:"default:6000000"`
			AdminKey string `conf:"default:0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"`
		} // NOTE: AdminKey was taken from hardhat network. These accounts, and their private keys, are publicly known.
		//Any funds sent to them on Mainnet or any other live network WILL BE LOST.
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

	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		dbURI = cfg.DB.Uri
	}

	log.Info().Str("status", "startup").Str("uri", dbURI).Msg("initializing database support")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return err
	}

	// -------------------------------------------------------------------------
	// Redis Support
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		redisUrl = cfg.Redis.Url
	}

	log.Info().Str("status", "startup").Str("url", redisUrl).Msg("initializing redis support")

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return fmt.Errorf("error parsing redis url %w", err)
	}

	redisClient := cachestore.NewRedisClient(redis.NewClient(opt))

	// -------------------------------------------------------------------------
	//ETH client support

	ethRawUrl := os.Getenv("HARDHAT_NODE_URL")
	if ethRawUrl == "" {
		ethRawUrl = cfg.ETH.Rawurl
	}

	log.Info().Str("status", "startup").Str("url", ethRawUrl).Msg("initializing eth client support")

	ethclient, err := blockchain.NewClient(ethRawUrl)
	if err != nil {
		return fmt.Errorf("error creating eth client %w", err)
	}

	userStore := userdb.NewStore(log, client)
	userCore, err := user.NewCore(userStore, ethclient, redisClient)
	if err != nil {
		return err
	}

	admin, err := contractAuth.NewAdmin(cfg.ETH.AdminKey, ethclient.Client)
	if err != nil {
		return fmt.Errorf("error creating admin %w", err)
	}
	cAuth, err := admin.GenerateAuth(ctx)
	if err != nil {
		return fmt.Errorf("error creating auth %w", err)
	}

	cAuth.GasLimit = cfg.ETH.GasLimit

	postContractAddress, _, instance, err := contract.DeployContract(cAuth, ethclient.Client)
	if err != nil {
		return fmt.Errorf("error creating post contract %w", err)
	}

	log.Info().Str("status", "startup").Msg("deployed contract with address" + postContractAddress.Hex())

	postContract, err := contract.NewPostContract(ethclient.Client, instance)
	if err != nil {
		return err
	}

	postCore := post.NewCore(postContract, admin, redisClient)

	// -------------------------------------------------------------------------
	//media support

	maxSize, err := strconv.ParseInt(cfg.Media.MaxFileSizeMb, 10, 64)
	if err != nil {
		return err
	}

	mediaDb := mediadb.NewStore(log, client)
	mediaCore := media.NewCore(mediaDb, redisClient, userCore)
	mediaCore.MaxFileSizeMb = maxSize

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

	debugMux := debug.StandardLibraryMux()
	debugMux.Handle("/swagger/", httpSwagger.WrapHandler)
	go func() {
		if err := http.ListenAndServe(cfg.Web.DebugHost, debugMux); err != nil {
			log.Error().
				Str("status", "debug v1 router closed").
				Str("host", cfg.Web.DebugHost).
				Err(err).
				Msg("shutdown")

		}
	}()

	//initial

	// -------------------------------------------------------------------------
	// Start API Service

	log.Info().Msg("initializing V1 API support")
	shutdown := make(chan os.Signal, 1)
	apiMux := handlers.APIMux(handlers.APIMuxConfig{
		Shutdown:  shutdown,
		Log:       log,
		Auth:      auth,
		UserCore:  userCore,
		PostCore:  postCore,
		MediaCore: mediaCore,
	})

	apiMux.Handle("/swagger", "GET", swaggerHandler())

	corsMiddleware := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"*"}), // Allow all origins, customize as needed
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		gorillaHandlers.ExposedHeaders([]string{"Authorization"}),
	)

	corsWrappedMux := corsMiddleware(apiMux)

	errorLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      corsWrappedMux,
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
func swaggerHandler() web.Handler {
	swaggerHandler := httpSwagger.WrapHandler
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		swaggerHandler(w, r)
		return nil
	}
}
