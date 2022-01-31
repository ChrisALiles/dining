package dining

import (
	"fmt"
	"os"
	"time"
)

var logchan chan string
var nl bool

// Lgger creates the log file and writes the entries passed to
// via the channel.
func Logger(log chan string, nolog bool) {
	logchan = log
	nl = nolog
	if nolog {
		return
	}
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

// Pass log entries to the Logger via the channel.
func Log(text string) {
	if !nl {
		logchan <- text
	}
}
