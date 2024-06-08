package risingwave

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var instantiatedConnection *pgx.Conn
var once sync.Once

func GetConnection() *pgx.Conn {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	once.Do(func() {
		// Please replace the placeholders with the actual credentials.
		connStr := os.Getenv("RISINGWAVE_ENDPOINT")
		conn, err := pgx.Connect(context.Background(), connStr)
		if err != nil {
			logrus.Fatalf("Unable to connect to RisingWave: %v\n", err)
		}

		//defer conn.Close(context.Background())
		instantiatedConnection = conn
	})
	return instantiatedConnection
}

func RunQuery(sql string) (pgx.Rows, error) {
	return GetConnection().Query(context.Background(), sql)
}
