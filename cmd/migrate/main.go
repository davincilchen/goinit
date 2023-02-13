package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
	"xr-central/pkg/app/users/usecase"
	"xr-central/pkg/config"
	"xr-central/pkg/db"
	"xr-central/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const confPath = "./config.json"
const seedPath = "./seed.json"

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
	d, err := db.GormOpen(&cfg.DB, &l)
	if err != nil {
		log.Fatal(err)
	}

	migration(d)

	loadSeed := false
	for _, v := range os.Args {
		if v == "loadseed" {
			loadSeed = true
		}
	}
	fmt.Println("loadSeed ", loadSeed)
	if loadSeed {
		err = seed()
		if err != nil {
			log.Fatal(err)
		}
	}

}

func migration(db *gorm.DB) {

	log.Println("Run DB Migration --> Start")

	db.AutoMigrate(&models.Edge{})
	db.AutoMigrate(&models.Streaming{})
	db.AutoMigrate(&models.EdgeStreaming{})

	db.AutoMigrate(&models.Device{})
	db.AutoMigrate(&models.EdgeOrder{})

	db.AutoMigrate(&models.Platform{})
	db.AutoMigrate(&models.App{})
	//db.AutoMigrate(&models.User{})

	log.Println("Run DB Migration --> Done")

}

func seed() error {

	log.Println("Create Seed --> Start")

	user, err := loadSeed(seedPath)
	if err != nil {
		return err
	}

	IO := usecase.User{}
	u, _ := IO.Register(user) //char(32) will be error

	fmt.Printf("%+v", u)

	log.Println("Create Seed --> Done")
	return nil
}

func loadSeed(path string) (*models.User, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	type Seed struct {
		User models.User `json:"User"`
	}

	var seed Seed

	err = json.Unmarshal(buf, &seed)
	if err != nil {
		return nil, err
	}

	return &seed.User, nil
}
