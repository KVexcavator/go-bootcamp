// использование пакетов log и log/syslog для записи в файлы системного журнала
package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

// $ go run log_files.go
// Will you see this?
// выходные данные попадают в несколко файлов журнала
// $ grep LOG_MAIL /var/log/mail.log
// $ grep LOG_LOCAL7 /var/log/cisco.log
// $ grep LOG_ /var/log/syslog
func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
}
