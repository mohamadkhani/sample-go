package messageBroker

import (
	"os"

	"github.com/joho/godotenv"
)

var err = godotenv.Load()
var KafkaEndpoint = os.Getenv("KAFKA_ENDPOINT")
