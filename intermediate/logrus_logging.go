package intermediate

import "github.com/sirupsen/logrus"

func main() {
	// go mod init name will initialize mod file
	// go get github.com/sirupsen/logrus will import logrus package
	log := logrus.New()

	// Set log level
	log.SetLevel(logrus.InfoLevel)

	// Set log format
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")

	log.WithFields(logrus.Fields{
		"userName": "John Smith",
		"method":   "GET",
	}).Info("User logged in.")
}
