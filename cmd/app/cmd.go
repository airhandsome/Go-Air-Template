package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "My Go application",
	Long:  `My Go application with HTTP and gRPC API`,
	Run: func(cmd *cobra.Command, args []string) {
		// Run the application
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}

func initConfig() {
	// Use viper to read configuration
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
	psqlConfig := viper.Sub("psql")
	fmt.Printf("PostgreSQL Host: %s\n", psqlConfig.GetString("host"))

	redisConfig := viper.Sub("redis")
	fmt.Printf("Redis Address: %s\n", redisConfig.GetString("address"))

	mongodbConfig := viper.Sub("mongodb")
	fmt.Printf("MongoDB URI: %s\n", mongodbConfig.GetString("uri"))
}
