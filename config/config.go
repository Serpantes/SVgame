package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var LogLevel string
var timestampFormat = "2006-01-02 15:04:05"


func InitConfig(){
	log.Debug("Reading config...")
	// проверили и если передали перееменную, то скорее всего сборка будет в контейнер
	// и .env файл там не нужен, т.к. параметры динамически меняются
	if os.Getenv("env") != "PROD" {
		log.Info("not in PROD. Reading .env")

		var errEnv error

		//Check if .env.local exists
		if _, err := os.Stat(".env.local"); err == nil {
			log.Info(".env.local found")
			errEnv = godotenv.Load(".env.local")
		} else {
			log.Info(".env found")
			errEnv = godotenv.Load(".env")
		}

		if errEnv != nil {
			log.Fatal("Can't read .env file. ", errEnv)
		}
	}

	log.Trace("Reading env vars")

	LogLevel = os.Getenv("LOG_LEVEL")
}


func InitLogger() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}

	formatter.TimestampFormat = timestampFormat
	log.SetFormatter(formatter)

	switch strings.ToLower(LogLevel) {

	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	}

	log.Debug("Logger ready")
}
