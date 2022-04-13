package main

import (
	"fmt"
	"os"
	"time"

	"github.com/amatute/go-playground/concurrency/database"
	"github.com/amatute/go-playground/concurrency/prices"
	"github.com/amatute/go-playground/concurrency/prices/repository"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	loadConfig()
	database.DBConn = initDatabase()
	database.DBConn.AutoMigrate(&repository.Price{})
	// seedTestDatadb() TODO: fill the database with prices

	repo := repository.NewRepository(database.DBConn)
	service := prices.NewService(repo)

	err := service.ReadData()
	if err != nil {
		logrus.Error("error reading data from price service 😑")
	}

	fmt.Println("we are ready to Go! ✅")
}

func initDatabase() *gorm.DB {
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")

	fmt.Printf("name %s user %s password %s host %s port %s", DBName, DBUser, DBPassword, DBHost, DBPort)

	fmt.Println("connecting to database... ⏲️")
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", DBUser, DBPassword, DBHost, DBPort, DBName)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to DB 💥")
	}
	fmt.Println("successfully connected to database 👌")
	return db
}

func loadConfig() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatalf("error loading .env file, can't config the app 💥")
	}
}

func seedTestDatadb() {
	for i := 0; i < 1000; i++ {
		p := &repository.Price{
			ID:            int(uuid.New().ID()),
			ProductNumber: fmt.Sprintf("PROD_%d", i),
			StoreNumber:   fmt.Sprintf("STORE_%d", i),
			Price:         "500",
			SalePrice:     "250",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		database.DBConn.Create(p)
	}
}
