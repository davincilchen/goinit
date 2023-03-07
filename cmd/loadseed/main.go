package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
	"xr-central/pkg/app/user/usecase"
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
	ddb, err := db.GormOpen(&cfg.DB, &l)
	if err != nil {
		log.Fatal(err)
	}
	db.MainDB = ddb

	fmt.Println("LoadSeed --> Start")

	err = saveSeed()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("LoadSeed --> Done")
}

type Seed struct {
	Users     []models.User
	Platforms []models.Platform
}

func saveSeed() error {

	fmt.Println("Seed: Create User --> Start")

	s, err := loadSeed(seedPath)
	if err != nil {
		return err
	}

	IO := usecase.User{}
	for _, user := range s.Users {
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

func loadSeed(path string) (*Seed, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	type LoopUsers struct {
		models.User
		Count int `json:"count"`
	}
	type TmpSeed struct {
		Users     []models.User `json:"Users"`
		LoopUsers []LoopUsers   `json:"LoopUsers"`
	}

	var tmpSeed TmpSeed

	err = json.Unmarshal(buf, &tmpSeed)
	if err != nil {
		return nil, err
	}

	seed := &Seed{}

	for _, v := range tmpSeed.LoopUsers {
		fmt.Println(v)
		for i := 0; i < v.Count; i++ {
			u := v.User
			no := fmt.Sprintf("%d", i+1)
			u.Name = u.Name + no
			u.Account = u.Account + no
			u.Password = u.Password + no
			tmpSeed.Users = append(seed.Users, u)
		}
	}

	seed.Users = tmpSeed.Users
	return seed, nil
}
