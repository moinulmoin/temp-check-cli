package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  float64 `json:"humidity"`
	} `json:"main"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
		return
	}

	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		fmt.Printf("API key is not set")
		return
	}

	city := "Dhaka"

	if len(os.Args) > 1 {
		city = os.Args[1]
	}

	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v&units=metric", city, apiKey)
	res, err := http.Get(apiUrl)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("Status code: %v", res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("%v, current temperature is %v° but feels like %v° and humidity is %v%%\n", city, weather.Main.Temp, weather.Main.FeelsLike, weather.Main.Humidity)
}
