package intermediate

// We use go get to get zap logger

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error in initializing zap logger")
	}
	defer logger.Sync()
	logger.Info("This is an info message")
	logger.Info("User logged in", zap.String("username", "John Smith"), zap.String("method", "GET"))
}
