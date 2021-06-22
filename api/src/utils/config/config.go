package config

import "os"

//postgres env
var (
	Host       = os.Getenv("POSTGRES_HOST")
	DBUser     = os.Getenv("POSTGRES_USER")
	DBPassword = os.Getenv("POSTGRES_PASSWORD")
	DBName     = os.Getenv("POSTGRES_DB")
	Port       = os.Getenv("POSTGRES_PORT")
)

//redis env
var (
	RedisHost     = os.Getenv("REDIS_HOST")
	RedisPort     = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
)

//token secret env
var (
	AccessTokenKey  = os.Getenv("ACCESS_TOKEN_SECRETE")
	RefreshTokenKey = os.Getenv("REFRESH_TOKEN_SECRETE")
	Issuer          = os.Getenv("TOKEN_ISSUER")
)

//postgres config
func postgresInit() {
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

//redis config
func redisInit() {
	host := "127.0.0.1"
	password := "redis_password"
	port := "56379"

	if RedisHost == "" {
		RedisHost = host
	}
	if RedisPort == "" {
		RedisPort = port
	}
	if RedisPassword == "" {
		RedisPassword = password
	}
}

func init() {
	postgresInit()
	redisInit()
}
