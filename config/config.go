package config

import (
	"os"
)

//DatabaseConfig struct for database
type DatabaseConfig struct {
  Username string
  Password string
  Host string
  Schema string 
}

// Config project config
type Config struct {
  Datbase DatabaseConfig
}

//New creates instance of config
func New() *Config {
  return &Config {  
    Datbase: DatabaseConfig {
      Username: getEnv("MYSQL_USERS_USERNAME",""),
      Password: getEnv("MYSQL_USERS_PASSWORD",""),
      Host: getEnv("MYSQL_USERS_HOST",""),
      Schema: getEnv("MYSQL_USERS_SCHEMA",""),
    },
  }
}

func getEnv(key string, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
      return value
    }
    return defaultValue
}

