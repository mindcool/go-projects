package intermediate

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a log message")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message")

	// Log Flags
	// Instead of getting date and time we only gate date
	log.SetFlags(log.Ldate)
	log.Println("This is a log message with only date")

	// We can chain flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time and filename")
	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is a error message")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// This one going to print the error to console
		// It will also exit the program at the same time
		log.Fatalf("Failed to open the log file: ", err)
	}
	defer file.Close()

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message")
}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
