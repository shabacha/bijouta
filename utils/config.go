package utils

import (
	"github.com/hashicorp/go-hclog"
	"github.com/lib/pq"
	"github.com/spf13/viper"
)

type Configurations struct {
	ServerAddress              string
	DBHost                     string
	DBName                     string
	DBUser                     string
	DBPass                     string
	DBPort                     string
	DBConn                     string
	AccessTokenPrivateKeyPath  string
	AccessTokenPublicKeyPath   string
	RefreshTokenPrivateKeyPath string
	RefreshTokenPublicKeyPath  string
	JwtExpiration              int // in minutes
}

func NewConfigurations(logger hclog.Logger) *Configurations {
	viper.AutomaticEnv()
	dbUrl := viper.GetString("DATABASE_URL")
	conn, _ := pq.ParseURL(dbUrl)
	logger.Debug("Found database url")
	logger.Debug("DB connection string --> ", conn)
	viper.SetDefault("SERVER_ADDRESS", "0.0.0.0:9090")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "bookite")
	viper.SetDefault("DB_USER", "vignesh")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("ACCESS_TOKEN_PRIVATE_KEY_PATH", "./access-private.pem")
	viper.SetDefault("ACCESS_TOKEN_PUBLIC_KEY_PATH", "./access-public.pem")
	viper.SetDefault("REFRESH_TOKEN_PRIVATE_KEY_PATH", "./refresh-private.pem")
	viper.SetDefault("REFRESH_TOKEN_PUBLIC_KEY_PATH", "./refresh-public.pem")
	viper.SetDefault("JWT_EXPIRATION", 30)
	configs := &Configurations{
		ServerAddress:              viper.GetString("SERVER_ADDRESS"),
		DBHost:                     viper.GetString("DB_HOST"),
		DBName:                     viper.GetString("DB_NAME"),
		DBUser:                     viper.GetString("DB_USER"),
		DBPass:                     viper.GetString("DB_PASSWORD"),
		DBPort:                     viper.GetString("DB_PORT"),
		DBConn:                     conn,
		JwtExpiration:              viper.GetInt("JWT_EXPIRATION"),
		AccessTokenPrivateKeyPath:  viper.GetString("ACCESS_TOKEN_PRIVATE_KEY_PATH"),
		AccessTokenPublicKeyPath:   viper.GetString("ACCESS_TOKEN_PUBLIC_KEY_PATH"),
		RefreshTokenPrivateKeyPath: viper.GetString("REFRESH_TOKEN_PRIVATE_KEY_PATH"),
		RefreshTokenPublicKeyPath:  viper.GetString("REFRESH_TOKEN_PUBLIC_KEY_PATH"),
	}
	port := viper.GetString("PORT")
	if port != "" {
		logger.Debug("Using port : ", port)
		configs.ServerAddress = "0.0.0.0:" + port
	}
	logger.Debug("Serve port --> ", configs.ServerAddress)
	logger.Debug("DB host --> ", configs.DBHost)
	logger.Debug("DB name --> ", configs.DBName)
	logger.Debug("DB port --> ", configs.DBPort)
	logger.Debug("JWT expiration --> ", configs.JwtExpiration)
	return configs
}
