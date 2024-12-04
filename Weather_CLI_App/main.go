package main

import (
	"fmt"
	"log"
	"os"

	"Weather_CLI_App/services"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error laoding .env flie")
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatal("API_KEY not set in environment variables")
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: weather-cli [CITY]")
		return
	}

	// city := os.Args[1]

	// response, err := services.FetchWeatherData(city, apiKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// displayWeather(response)

	var city string
	var rootCmd = &cobra.Command{
		Use:   "weather-cli",
		Short: "A CLI to get weather information",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if city argument is provided
			if city == "" {
				fmt.Println("Please provide a city.")
				return
			}

			// Fetch weather data
			response, err := services.FetchWeatherData(city, apiKey)
			if err != nil {
				log.Fatal(err)
			}

			// Display weather data
			displayWeather(response)
		},
	}

	// Add a flag for the city
	rootCmd.Flags().StringVarP(&city, "city", "c", "", "City name to get weather information for")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func displayWeather(weather *services.WeatherResponse) {
	fmt.Printf("Weather in %s:\n", weather.Name)
	fmt.Printf("Temperature: %.1f°C\n", weather.Main.Temp)
	fmt.Printf("Condition: %s\n", weather.Weather[0].Description)
}
