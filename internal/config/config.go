package config

import (
	"github.com/joho/godotenv"
	"os"
)

// const MongoUrl = "mongodb://mongo:27017/"
const Dbname = "strux_api"
const UsersCN = "users"
const Host = "0.0.0.0"
const Port = "3000"

func GetMongoAddress() string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	MongoAddress := os.Getenv("MONGO_ADDRESS")
	return MongoAddress
}

func GetUserServiceAddress() string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	UserServiceAddress := os.Getenv("USER_SERVICE_ADDRESS")
	return UserServiceAddress
}

func GetPkgServiceAddress() string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	PkgServiceAddress := os.Getenv("PACKAGE_SERVICE_ADDRESS")
	return PkgServiceAddress
}

const APILogFileName = "api_logs.log"
const UserServiceLogFileName = "user_service_logs.log"
const PackageServiceLogFileName = "package_service_logs.log"

const PackagesDirPath = "../UserPackages"
