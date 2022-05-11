// исползование log.Fatal()
package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Fatal(sysLog)
	fmt.Println("Will you see this?")
}

// $ grep "Some program" /var/log/mail.log
