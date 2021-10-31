package main

import (
	"fmt"
	"go-linebot-crypto/handler"
	"go-linebot-crypto/repository"
	"go-linebot-crypto/service"
	"time"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	initTimeZone()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := connectDatabase()
	if err != nil {
		panic(err)
	}

	// Connect to Line Bot
	bot := connectLineBot()

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)
	linBotHandler := handler.NewLineBotHandler(bot)

	router := mux.NewRouter()
	// Setup HTTP Server for receiving requests from LINE platform
	router.HandleFunc("/webhook", linBotHandler.HandleWebHook)
	router.HandleFunc("/customers", customerHandler.GetCustomers)
	router.HandleFunc("/customer/{customerID:[0-9]+}", customerHandler.GetCustomer).GetMethods()

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "4325"
		log.Printf("defaulting to port %s", port)
	}

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

func connectDatabase() (*sqlx.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSL"),
	)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

// Connect to Line Bot
func connectLineBot() *linebot.Client {
	log.Print("🚀 starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}
