package server

import (
	"fmt"
	"time"

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
	ErrorTxt     string
	Error        error
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

	log.RequestURI = ctx.Request.RequestURI
	log.Method = ctx.Request.Method
	log.Duration = time.Since(now)
	log.DurationText = fmt.Sprintf("%v", log.Duration)

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
	s := fmt.Sprintf("%#v", log)
	fmt.Println(s)

	// gcp.stdLogger.Println(s)

	// //defer gcp.logger.Flush() // Ensure the entry is written.
	// gcp.logger.Log(gcpLogging.Entry{
	// 	// Log anything that can be marshaled to JSON.
	// 	Payload:  log,
	// 	Severity: gcpLogging.Debug,
	// })
}
