package server

import (
	"log"
	"os"
	"xr-central/pkg/config"
	"xr-central/pkg/db"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Config *config.Config
}

func initLogger() {
	//log輸出為json格式
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//輸出設定為標準輸出(預設為stderr)
	logrus.SetOutput(os.Stdout)
	//設定要輸出的log等級
	logrus.SetLevel(logrus.DebugLevel)
}

func New(cfg *config.Config) *Server {

	return &Server{
		Config: cfg,
	}

}

func (t *Server) Serve() {

	initLogger()

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

	//InitLogger("", t.Config.GCP.ProjectID, t.Config.GCP.LogName)

	addr := ":" + t.Config.Server.Port
	log.Printf("======= Server start to listen (%s) and serve =======\n", addr)
	r := Router()
	r.Run(addr)

	log.Printf("======= Server Exit =======\n")
	//CloseLogger()
}
