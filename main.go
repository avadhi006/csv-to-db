package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/avadhi006/csv-to-db/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Set up connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migrate schema
	db.AutoMigrate(&models.Record{})

	return db, nil
}

func parseCSV(filePath string) ([]models.Record, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var records []models.Record

	// Skip header row
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Read CSV rows
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Convert CSV row to Record struct
		records = append(records, models.Record{
			SiteID:                record[0],
			FixletID:              record[1],
			Name:                  record[2],
			Crtiticality:          record[3],
			RelevantComputerCount: record[4],
		})
	}
	return records, nil
}

func insertRecords(db *gorm.DB, records []models.Record) error {
	// Batch insert records into DB
	return db.Create(&records).Error
}

func main() {
	// Load environment variables
	loadEnv()

	// Initialize database connection
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	// Parse the CSV file
	records, err := parseCSV("fixlets.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Insert records into the database
	err = insertRecords(db, records)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data successfully imported to database!")
}
