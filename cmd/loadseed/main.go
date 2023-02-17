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
	_, err = db.GormOpen(&cfg.DB, &l)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("LoadSeed --> Start")

	err = seed()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("LoadSeed --> Done")
}

func seed() error {

	fmt.Println("Seed: Create User --> Start")

	users, err := loadSeed(seedPath)
	if err != nil {
		return err
	}

	IO := usecase.User{}
	for _, user := range users {
		u, err := IO.Register(&user) //char(32) will be error
		if err != nil {
			fmt.Printf("err =%s, for Register User: %+v", err.Error(), u)
		} else {
			fmt.Printf("Register User: %+v\n", u)
		}

	}

	fmt.Println("Seed: Create User  --> Done")
	return nil
}

func loadSeed(path string) ([]models.User, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	type LoopUsers struct {
		models.User
		Count int `json:"count"`
	}
	type Seed struct {
		Users     []models.User `json:"Users"`
		LoopUsers []LoopUsers   `json:"LoopUsers"`
	}

	var seed Seed

	err = json.Unmarshal(buf, &seed)
	if err != nil {
		return nil, err
	}

	for _, v := range seed.LoopUsers {
		fmt.Println(v)
		for i := 0; i < v.Count; i++ {
			u := v.User
			no := fmt.Sprintf("%d", i+1)
			u.Name = u.Name + no
			u.Account = u.Account + no
			u.Password = u.Password + no
			seed.Users = append(seed.Users, u)
		}
	}
	return seed.Users, nil
}
