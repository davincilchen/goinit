package server

import (
	"encoding/json"
	"fmt"
	"time"
	devUCase "xr-central/pkg/app/device/usecase"
	"xr-central/pkg/app/infopass"

	// gcpLogging "cloud.google.com/go/logging"
	"github.com/gin-gonic/gin"
)

type APILog struct {
	Version      string
	RequestURI   string
	Method       string
	Duration     time.Duration
	DurationText string
	InfoTxt      string
	DBErrorTxt   string
	ErrorTxt     string
	AdvErrorTxt  string

	RequestBody  *string `json:"requestBody"`
	SessionToken *string `json:"sessionToken"`
	DevData      interface{}
}

// type Gcp struct {
// 	client    *gcpLogging.Client
// 	logger    *gcpLogging.Logger
// 	stdLogger *log.Logger
// }

//gcloud auth application-default login
//gcloud auth application-default revoke

// var gcp *Gcp

func InitLogger(key string, project string, logName string) error {
	// //client, err := gcpLogging.NewClient(context.Background(), project, option.WithCredentialsFile(key))
	// client, err := gcpLogging.NewClient(context.Background(), project)
	// if err != nil {
	// 	return err
	// }
	// gcp = &Gcp{client: client,
	// 	logger:    client.Logger(logName),
	// 	stdLogger: client.Logger(logName).StandardLogger(gcpLogging.Info)}

	return nil
}

func CloseLogger() {
	// if gcp != nil {
	// 	gcp.client.Close()
	// }
}

func Logger(ctx *gin.Context) {
	now := time.Now()

	ctx.Next()

	log := APILog{}

	theError := infopass.GetError(ctx)
	theAdvError := infopass.GetAdvError(ctx)
	theDBError := infopass.GetDBError(ctx)
	log.RequestURI = ctx.Request.RequestURI
	log.Method = ctx.Request.Method
	log.Duration = time.Since(now)
	log.DurationText = fmt.Sprintf("%v", log.Duration)

	log.RequestBody, _ = infopass.GetRequestBodyInGin(ctx)
	log.SessionToken = infopass.GetSessionToken(ctx)
	log.DevData = devUCase.GetCacheDevice(ctx)
	// log.PlayerSession, _ = GetPlayerSessionInGin(ctx)

	if theError != nil {
		log.ErrorTxt = theError.Error()
	}
	if theAdvError != nil {
		log.AdvErrorTxt = theAdvError.Error()
	}
	if theDBError != nil {
		log.DBErrorTxt = theDBError.Error()
	}

	logger(log)

	// if theError != nil {
	// 	log.Error = theError
	// 	log.ErrorTxt = log.Error.Error()
	// 	Warning(log)
	// } else {
	// 	Info(log)
	// }

}

func logger(log APILog) {

	// if gcp == nil {
	// 	fmt.Printf("[gcp == nil] %#v", log)
	// 	fmt.Println(log)
	// 	return
	// }
	fmt.Println()
	s := fmt.Sprintf("%#v", log)
	fmt.Println(s)

	if log.SessionToken != nil {
		s := fmt.Sprintf("SessionToken %#v", *log.SessionToken)
		fmt.Println(s)
	}

	//if log.DevData != nil { 無法檢查interface, 這裡可以不檢查
	b, err := json.Marshal(log.DevData)
	if err == nil {
		fmt.Println("DevData:", string(b))
	}

	// gcp.stdLogger.Println(s)

	// //defer gcp.logger.Flush() // Ensure the entry is written.
	// gcp.logger.Log(gcpLogging.Entry{
	// 	// Log anything that can be marshaled to JSON.
	// 	Payload:  log,
	// 	Severity: gcpLogging.Debug,
	// })
}
