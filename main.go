package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/labstack/echo/v4"

	mockapiRepo "moneytransfer/internal/repository/mockapi"
	postgresqlRepo "moneytransfer/internal/repository/postgresql"

	"moneytransfer/service"
	"moneytransfer/internal/handler"
)

const (
	defaultAddress = ":9090"
)

func init() {

}

func main() {
	// prepare database
	// dbUser := os.Getenv("DATABASE_USER")
	// dbName := os.Getenv("DATABASE_NAME")
	dbUser := "postgres"
	dbName := "moneytransfer"
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbUser, dbName)	// should enable ssl later
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("failed to close connection from database", err)
		}
	}()

	// prepare echo
	e := echo.New()

	// prepare repository layer
	postgresqlTransactionRepo := postgresqlRepo.NewTransactionRepository(dbConn)
	mockapiAccountRepo := mockapiRepo.NewAccountRepository()
	mockapiTransactionRepo := mockapiRepo.NewTransactionRepository(postgresqlTransactionRepo)

	// prepare service & handler layer
	bankService := service.NewBankService(mockapiAccountRepo, mockapiTransactionRepo)
	handler.NewBankHandler(e, bankService)

	// start server
	log.Fatal(e.Start(defaultAddress))
}