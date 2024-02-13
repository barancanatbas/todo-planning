package main

import (
	"github.com/creasty/defaults"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"task-manager/internal/config"
	"task-manager/internal/router"
	"task-manager/pkg/mariadb"
	"task-manager/pkg/viper"
)

func main() {
	config := readConfig()
	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	db := mariadb.NewMariaDbConnection(config.MariaDB)

	router := router.NewRouter(app)
	router.RegisterRoutes(db)

	if err := app.Listen(":" + config.App.Port); err != nil {
		log.Fatal().Err(err).Msg("error starting application")
	}
}

func readConfig() config.Config {
	var cnf config.Config
	if err := defaults.Set(&cnf); err != nil {
		log.Fatal().Err(err).Msg("error setting config default values")
	}

	viper.SetupConfig(&cnf)

	if level, err := zerolog.ParseLevel(cnf.Log.Level); err != nil {
		log.Fatal().Err(err).Msg("error parsing log level")
	} else {
		log.Level(level)
	}

	return cnf
}
