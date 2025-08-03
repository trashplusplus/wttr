package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wttr/internal/locale"
)

const GetWeatherURL = "https://wttr.in/%s?format=j2&lang=uk"
const GetWeatherIconURL = "https://wttr.in/%s?0&A&lang=uk&format=%%c"
const GetWeatherIconURLAuto = "https://wttr.in/?0&A&lang=uk&format=%c"

type WeatherInfo struct {
	City        string
	Temperature string
	Description string
	Icon        string
}

type apiResponse struct {
	CurrentCondition []struct {
		TempC       string `json:"temp_C"`
		WeatherDesc []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
	} `json:"current_condition"`
	NearestArea []struct {
		AreaName []struct {
			Value string `json:"value"`
		} `json:"areaName"`
	} `json:"nearest_area"`
}

func GetWeather(city string) (WeatherInfo, error) {
	url := fmt.Sprintf(GetWeatherURL, city)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherInfo{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherInfo{}, err
	}

	var data apiResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return WeatherInfo{}, err
	}

	if len(data.CurrentCondition) == 0 || len(data.NearestArea) == 0 {
		return WeatherInfo{}, fmt.Errorf("некоректні дані з wttr.in")
	}

	icon, err := GetWeatherIcon(city)
	if err != nil {
		fmt.Println("Помилка отримання іконки:", err)
		return WeatherInfo{}, err
	}

	info := WeatherInfo{
		City:        data.NearestArea[0].AreaName[0].Value,
		Temperature: data.CurrentCondition[0].TempC,
		Description: data.CurrentCondition[0].WeatherDesc[0].Value,
		Icon:        icon,
	}

	return info, nil
}

func GetWeatherIcon(city string) (string, error) {

	url := fmt.Sprintf(GetWeatherIconURL, city)

	if city == "" {
		url = GetWeatherIconURLAuto
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP статус: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func Show(weatherInfo WeatherInfo) {
	fmt.Println("Місто: ", weatherInfo.City)
	fmt.Println("Температура: ", weatherInfo.Temperature, "°C")
	fmt.Println("Опис: ", locale.Translate(weatherInfo.Description))
	fmt.Println("Іконка: ", weatherInfo.Icon)
}

func ShowLocaledDesc(weatherInfo WeatherInfo) {
	fmt.Println(weatherInfo.City, ": ", locale.Translate(weatherInfo.Description))
}
