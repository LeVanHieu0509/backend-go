package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerStruct struct {
	Port int `mapstructure:"port"`
}

type DatabaseStruct struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}

type Config struct {
	Server    ServerStruct     `mapstructure:"server"`
	Databases []DatabaseStruct `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	//read configuration
	err := viper.ReadInConfig()

	if err != nil {
		//%w giữ nguyên chuỗi lỗi để gỡ lỗis
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}

	// Read server configuration
	fmt.Println("Server Port:", viper.GetInt("server.port"))
	fmt.Println("security Port:", viper.GetString("security.jwt.key"))

	// configuration structure
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Println("Config Port::", config.Server.Port)

	for index, db := range config.Databases {
		fmt.Printf("database User-%v: %s, password: %s, host: %s \n", index, db.User, db.Password, db.Host)
	}

}
