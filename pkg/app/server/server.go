package server

import (
	"log"
	"xr-central/pkg/config"
	"xr-central/pkg/db"
)

type Server struct {
	Config *config.Config
}

func New(cfg *config.Config) *Server {

	return &Server{
		Config: cfg,
	}

}

func (t *Server) Serve() {

	dbConn, err := db.GormOpen(&t.Config.DB, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.MainDB = dbConn

	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	if t.Config.Hello.Show != "" {
		show = t.Config.Hello.Show
	}

	//InitLogger("", t.Config.GCP.ProjectID, t.Config.GCP.LogName)

	addr := ":" + t.Config.Server.Port
	log.Printf("======= Server start to listen (%s) and serve =======\n", addr)
	r := Router()
	r.Run(addr)

	log.Printf("======= Server Exit =======\n")
	//CloseLogger()
}
