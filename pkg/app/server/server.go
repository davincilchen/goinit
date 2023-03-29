package server

import (
	"log"
	"os"

	//"syscall"
	"xr-central/pkg/config"
	"xr-central/pkg/db"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"
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

	consoleQuickMode(false)
	//InitLogger("", t.Config.GCP.ProjectID, t.Config.GCP.LogName)

	addr := ":" + t.Config.Server.Port
	log.Printf("======= Server start to listen (%s) and serve =======\n", addr)
	r := Router()
	r.Run(addr)

	log.Printf("======= Server Exit =======\n")
	//CloseLogger()
}

// # input flags
// ENABLE_PROCESSED_INPUT = 0x0001
// ENABLE_LINE_INPUT      = 0x0002
// ENABLE_ECHO_INPUT      = 0x0004
// ENABLE_WINDOW_INPUT    = 0x0008
// ENABLE_MOUSE_INPUT     = 0x0010
// ENABLE_INSERT_MODE     = 0x0020
// ENABLE_QUICK_EDIT_MODE = 0x0040
// ENABLE_EXTENDED_FLAGS  = 0x0080

// # output flags
// ENABLE_PROCESSED_OUTPUT   = 0x0001
// ENABLE_WRAP_AT_EOL_OUTPUT = 0x0002
// ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004 # VT100 (Win 10)

type master struct {
	in     windows.Handle
	inMode uint32

	out     windows.Handle
	outMode uint32

	err     windows.Handle
	errMode uint32
}

func consoleQuickMode(enable bool) { //for windows
	//hStdin := syscall.Handle(os.Stdin.Fd())
	hStdin := windows.Handle(os.Stdin.Fd())

	var originalMode uint32
	var newMode uint32
	err := windows.GetConsoleMode(hStdin, &originalMode)
	if err != nil {
		logrus.Errorf("windows.GetConsoleMode error: %s", err.Error())
		return
	}

	logrus.Infof("windows.GetConsoleMode mode: %x", originalMode)
	newMode = originalMode
	if enable {
		//(0x100|0x80|0x00|0x00|0x10|0x4|0x2|0x1)
		//_enable := 0x0060 //|0x20|0x40
		// newMode &^= windows.ENABLE_INSERT_MODE
		// newMode &^= windows.ENABLE_QUICK_EDIT_MODE
		newMode |= windows.ENABLE_INSERT_MODE
		newMode |= windows.ENABLE_QUICK_EDIT_MODE

		//newMode = originalMode | uint32(_enable)
		logrus.Infof("[_enable] NewMode : %x ", newMode)
	} else { //disable
		//(0x100|0x80|0x00|0x00|0x10|0x4|0x2|0x1)
		//_disable := 0x019F //|0x20|0x40
		newMode &^= windows.ENABLE_INSERT_MODE
		newMode &^= windows.ENABLE_QUICK_EDIT_MODE

		//newMode = originalMode & uint32(_disable)
		logrus.Infof("[_disable] NewMode : %x ", newMode)
	}
	err = windows.SetConsoleMode(hStdin, newMode)
	if err != nil {
		logrus.Errorf("windows.SetConsoleMode Error %s", err)
		return
	}
	logrus.Infof("windows.SetConsoleMode Success")

	err = windows.GetConsoleMode(hStdin, &originalMode)
	if err != nil {
		logrus.Errorf("windows.GetConsoleMode error: %s", err.Error())
		return
	}

	logrus.Infof("windows.GetConsoleMode new mode: %x", originalMode)
}

// func consoleQuickMode(enable bool) { //for windows
// 	//hStdin := syscall.Handle(os.Stdin.Fd())
// 	hStdin := os.Stdin.Fd()

// 	var originalMode uint32
// 	var newMode uint32
// 	err := windows.GetConsoleMode(hStdin, &originalMode)
// 	if err != nil {
// 		logrus.Errorf("syscall.GetConsoleMode error: %s", err.Error())
// 		return
// 	}

// 	logrus.Errorf("syscall.GetConsoleMode mode: %x", originalMode)
// 	newMode = originalMode
// 	if enable {
// 		//(0x100|0x80|0x00|0x00|0x10|0x4|0x2|0x1)
// 		_enable := 0x0060 //|0x20|0x40
// 		newMode = originalMode | uint32(_enable)
// 		logrus.Errorf("[_enable] NewMode : %x , flag %x", newMode, uint32(_enable))
// 	} else {
// 		//(0x100|0x80|0x00|0x00|0x10|0x4|0x2|0x1)
// 		_disable := 0x019F //|0x20|0x40
// 		newMode = originalMode & uint32(_disable)
// 		logrus.Errorf("[_disable] NewMode : %x , flag %x", newMode, uint32(_disable))
// 	}
// 	err = syscall.SetConsoleMode(syscall.Stdin, newMode)
// }

//example of SetConsoleMode
//fd windows.Handle: trans -> os.Stdin.Fd() or os.Stdout.Fd()
//windows.Handle(os.Stdin.Fd())
// func makeInputRaw(fd windows.Handle, mode uint32) error {
// 	// See
// 	// -- https://msdn.microsoft.com/en-us/library/windows/desktop/ms686033(v=vs.85).aspx
// 	// -- https://msdn.microsoft.com/en-us/library/windows/desktop/ms683462(v=vs.85).aspx

// 	// Disable these modes
// 	mode &^= windows.ENABLE_ECHO_INPUT
// 	mode &^= windows.ENABLE_LINE_INPUT
// 	mode &^= windows.ENABLE_MOUSE_INPUT
// 	mode &^= windows.ENABLE_WINDOW_INPUT
// 	mode &^= windows.ENABLE_PROCESSED_INPUT

// 	// Enable these modes
// 	mode |= windows.ENABLE_EXTENDED_FLAGS
// 	mode |= windows.ENABLE_INSERT_MODE
// 	mode |= windows.ENABLE_QUICK_EDIT_MODE

// 	if vtInputSupported {
// 		mode |= windows.ENABLE_VIRTUAL_TERMINAL_INPUT
// 	}

// 	if err := windows.SetConsoleMode(fd, mode); err != nil {
// 		return fmt.Errorf("unable to set console to raw mode: %w", err)
// 	}

// 	return nil
// }
