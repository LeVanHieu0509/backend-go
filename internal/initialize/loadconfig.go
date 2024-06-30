package initialize

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load từ file yaml ra giống như yaml là file env
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	//read configuration
	err := viper.ReadInConfig()

	if err != nil {
		//%w giữ nguyên chuỗi lỗi để gỡ lỗi
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}

	// Read server configuration
	fmt.Println("Server Port:", viper.GetInt("server.port"))
	fmt.Println("security Port:", viper.GetString("security.jwt.key"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

}
