package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var GlobalConfig Config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Name     string `mapstructure:"name" json:"name"`
}

type SMTPConfig struct {
	From     string `mapstructure:"from" json:"from"`
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

type Config struct {
	Mysql MysqlConfig `mapstructure:"database" json:"database"`
	SMTP  SMTPConfig  `mapstructure:"smtp" json:"smtp"`
}

// DSN returns the Data Source Name
func DSN(ci MysqlConfig) string {
	// Example: root:@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Host +
		":" +
		ci.Port +
		")/" +
		ci.Name +
		"?charset=utf8mb4&parseTime=True"
}
func loadconfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config")
	//viper.SetConfigFile("./config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("error while reading config file:", zap.Error(err))
		return nil, fmt.Errorf("error while reading config file: %s", err)
	}
	err = viper.UnmarshalKey("database", &GlobalConfig.Mysql)
	viper.UnmarshalKey("SMTP", &GlobalConfig.SMTP)
	fmt.Println(GlobalConfig.Mysql)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling config file: %s", err)
	}

	fmt.Printf("config:%+v", GlobalConfig)
	return &GlobalConfig, nil
}

func SetDbConfig() (string, error) {
	Config, err := loadconfig()
	if err != nil {
		return "", err
	}
	dsn := DSN(Config.Mysql)
	fmt.Printf("dsn:%s", dsn)
	return dsn, nil
}
