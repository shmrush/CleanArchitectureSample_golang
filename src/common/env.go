package common

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type databaseEnv struct {
	Kind     string `envconfig:"DATABASE_KIND"`
	Port     string `envconfig:"DATABASE_PORT"`
	Host     string `envconfig:"DATABASE_HOST"`
	DbName   string `envconfig:"DATABASE_DBNAME"`
	User     string `envconfig:"DATABASE_USER"`
	Password string `envconfig:"DATABASE_PASSWORD"`
}

type kvsEnv struct {
	Port     string `envconfig:"KVS_PORT"`
	Host     string `envconfig:"KVS_HOST"`
	Protocol string `envconfig:"KVS_PROTOCOL"`
}

// DatabaseEnv provides database secrets.
var DatabaseEnv databaseEnv

// KvsEnv provides kvs secrets.
var KvsEnv kvsEnv

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	if err := envconfig.Process("", &DatabaseEnv); err != nil {
		log.Fatal("Error loading DatabaseEnv")
	}
	if err := envconfig.Process("", &KvsEnv); err != nil {
		log.Fatal("Error loading KvsEnv")
	}
}
