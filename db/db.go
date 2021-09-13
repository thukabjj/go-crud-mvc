package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectComBanco() *sql.DB {
	conexao := "user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " host=" + os.Getenv("POSTGRES_HOST") + " dbname=" + os.Getenv("POSTGRES_DB") + " sslmode=" + os.Getenv("POSTGRES_SSL_MODE")
	fmt.Println("Connect: ", conexao)
	db, err := sql.Open(os.Getenv("POSTGRES_DRIVER_NAME"), conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
