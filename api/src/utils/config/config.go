package config

import "os"

var (
	Host       = os.Getenv("POSTGRES_HOST")
	DBUser     = os.Getenv("POSTGRES_USER")
	DBPassword = os.Getenv("POSTGRES_PASSWORD")
	DBName     = os.Getenv("POSTGRES_DB")
	Port       = os.Getenv("POSTGRES_PORT")
)

var (
	AccessTokenKey  = os.Getenv("ACCESS_TOKEN_SECRETE")
	RefreshTokenKey = os.Getenv("REFRESH_TOKEN_SECRETE")
	Issuer          = os.Getenv("TOKEN_ISSUER")
)

func init() {
	// dbDriver := "pgx"
	host := "127.0.0.1"
	dbUser := "project"
	dbPassword := "password"
	dbName := "projectdb"
	port := "54320"

	if Host == "" {
		Host = host
	}
	if DBUser == "" {
		DBUser = dbUser
	}
	if DBPassword == "" {
		DBPassword = dbPassword
	}
	if DBName == "" {
		DBName = dbName
	}
	if Port == "" {
		Port = port
	}
}
