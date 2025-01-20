package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server Port
	Port          string
	SessionSecret string

	// Database settings
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBURL      string

	// Spotify Settings
	SpotifyClientId     string
	SpotifyClientSecret string
	SpotifyRedirectUrl  string

	/// Update interval in seconds
	UpdateInterval int
}

var AppConfig Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env, using enviromnent variables")
	}

	AppConfig = Config{
		Port:                getEnv("PORT", "8080"),
		SessionSecret:       getEnv("SESSION_SECRET", ""),
		DBHost:              getEnv("DATABASE_HOST", ""),
		DBPort:              getEnv("DATABASE_PORT", "5432"),
		DBUser:              getEnv("DATABASE_USER", ""),
		DBPassword:          getEnv("DATABASE_PASS", ""),
		DBName:              getEnv("DATABASE_NAME", "satpl"),
		SpotifyClientId:     getEnv("SPOTIFY_CLIENT_ID", ""),
		SpotifyClientSecret: getEnv("SPOTIFY_CLIENT_SECRET", ""),
		SpotifyRedirectUrl:  getEnv("SPOTIFY_REDIRECT_URL", "http://localhost:8080/callback"),
		UpdateInterval:      24 * 60 * 60,
	}
	AppConfig.DBURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBName)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("%s not found in env using defaultValue: %s\n", key, defaultValue)
	return defaultValue
}
