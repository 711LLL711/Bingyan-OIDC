package config

import (
	"OIDC/utils"
	"fmt"

	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// DSN returns the Data Source Name
func DSN(ci MysqlConfig) string {
	// Example: root:@tcp(localhost:3306)/test
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Host +
		":" +
		ci.Port +
		")/" +
		ci.Name
}
func loadconfig() (*MysqlConfig, error) {
	// viper.SetConfigName("config")
	// viper.AddConfigPath(".")
	viper.SetConfigFile("./config.json")
	err := viper.ReadInConfig()
	if err != nil {
		utils.Logger.Error("error while reading config file:")
		return nil, fmt.Errorf("error while reading config file: %s", err)
	}
	var mysqlConfig MysqlConfig
	err = viper.UnmarshalKey("database", &mysqlConfig)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling config file: %s", err)
	}

	fmt.Printf("config:%+v", mysqlConfig)
	return &mysqlConfig, nil
}

func SetDbConfig() (string, error) {
	mysqlConfig, err := loadconfig()
	if err != nil {
		return "", err
	}
	dsn := DSN(*mysqlConfig)
	fmt.Printf("dsn:%s", dsn)
	return dsn, nil
}
