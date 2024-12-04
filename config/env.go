package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tuananh31j/library-management-system/utils"
)

var (
	AppHost    string
	AppPort    int
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int
)

func init() {
	loadConfig()
	AppHost = viper.GetString("APP_HOST")
	AppPort = viper.GetInt("APP_PORT")
	DBHost = viper.GetString("DB_HOST")
	DBUser = viper.GetString("DB_USER")
	DBPassword = viper.GetString("DB_PASSWORD")
	DBName = viper.GetString("DB_NAME")
	DBPort = viper.GetInt("DB_PORT")

}

func loadConfig() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err == nil {
		utils.Log.Info("Config loaded successfully")
		return
	}

	utils.Log.Error(fmt.Sprintf("Error loading config file: %s", err))
}
