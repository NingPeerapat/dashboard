package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURI           string
	DatabaseName       string
	CollectionName     string
	CollectionTempName string
}

func LoadConfig() (*Config, error) {
	// โหลดค่า config จากไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	return &Config{
		MongoURI:           os.Getenv("MONGO_URI"),
		DatabaseName:       os.Getenv("DATABASE_NAME"),
		CollectionName:     os.Getenv("COLLECTION_NAME"),
		CollectionTempName: os.Getenv("COLLECTION_TEMP_NAME"),
	}, nil
}

// เชื่อมต่อ MongoDB ด้วยการใช้ URI จาก config
func ConnectMongo(cfg *Config) (*mongo.Collection, *mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	// ตรวจสอบว่าเชื่อมต่อได้หรือไม่
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error pinging MongoDB: %v", err)
	}

	collection := client.Database(cfg.DatabaseName).Collection(cfg.CollectionName)
	collectionTemp := client.Database(cfg.DatabaseName).Collection(cfg.CollectionTempName)

	log.Println("Connected to MongoDB")
	return collection, collectionTemp, nil
}
