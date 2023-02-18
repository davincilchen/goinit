package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"xr-central/pkg/config"
	"xr-central/pkg/db"

	"gorm.io/gorm/logger"
)

const confPath = "./config.json"

func main() {

	cfg, err := config.New(confPath)
	if err != nil {
		log.Fatal(err)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	l := db.Logger{
		Logger: newLogger,
	}
	fmt.Println("GormCreateDB --> Start")

	err = db.GormCreateDB(&cfg.DB, &l)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GormCreateDB --> Done")
}
