package server

import (
	"central/pkg/config"
	"central/pkg/db"
	"log"
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

	_, err := db.GormOpen(&t.Config.DB)
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

	CloseLogger()
}
