package main

import (
	"go-backend/internal/db"
	"go-backend/internal/env"
	"go-backend/internal/mailer"
	"go-backend/internal/store"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const version = "0.0.2"

var Validate *validator.Validate

func init() {
	// This setting is pretty useful as it ensures
	// nil struct pointers will still undergo validation
	// Without this setting, nil struct pointers would
	// bypass data validation which can cause critical bugs
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

//	@title			GopherSocial API
//	@description	API for GopherSocial, a social network for gohpers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www. swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description

func main() {
	cfg := config{
		addr: env.MustGetString("ADDR"),
		db: dbConfig{
			addr:         env.MustGetString("DB_ADDR"),
			maxOpenConns: env.MustGetInt("DB_MAX_OPEN_CONNS"),
			maxIdleConns: env.MustGetInt("DB_MAX_IDLE_CONNS"),
			maxIdleTime:  env.MustGetString("DB_MAX_IDLE_TIME"),
		},
		env: env.MustGetString("ENV"),
		swagger: swaggerConfig{
			host:    env.MustGetString("SWAGGER_HOST"),
			jsonURL: env.MustGetString("SWAGGER_JSON_URL"),
		},
		mail: mailConfig{
			exp: time.Hour * 24 * 3, // 3 days
			fromEmail: env.MustGetString("FROM_EMAIL"),
			sengrid: sengridConfig{
				apiKey: env.MustGetString("SENDGRID_API_KEY"),
			},
		},
		frontendURL: "http://localhost:5000",
		auth: authConfig{
			basic: basicAuthConfig{
				username: env.MustGetString("BASIC_AUTH_USERNAME"),
				pass: env.MustGetString("BASIC_AUTH_PASS"),
			},
		},
	}
	// logger
	prodConfig := zap.NewProductionConfig()
	prodConfig.DisableStacktrace = true
	zapLogger, _ := prodConfig.Build()
	logger := zapLogger.Sugar()
	defer logger.Sync() // flushes buffer, if any

	// setup the database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxIdleTime,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	store := store.NewPostgreSQLStorage(db)
	
	// mailer
	mailer := mailer.NewSendgrid(cfg.mail.sengrid.apiKey, cfg.mail.fromEmail)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
		mailer: mailer,
	}

	logger.Fatal(app.run())
}
