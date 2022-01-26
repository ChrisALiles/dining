package dining

import (
	"fmt"
	"os"
	"time"
)

var logchan chan string

func Logger(log chan string) {
	logchan = log
	f, err := os.Create(logfile)
	if err != nil {
		fmt.Println("Unable to create log file ", err)
		os.Exit(1)
	}
	defer f.Close()
	for text := range log {
		f.WriteString(time.Now().Format(time.StampMilli) + " " + text)
	}
}

func Log(text string) {
	logchan <- text
}
