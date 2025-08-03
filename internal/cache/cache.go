package cache

import (
	"time"
	"wttr/internal/weather"
)

type CachedWeather struct {
	Data      weather.WeatherInfo
	Timestamp time.Time
}

var cache = make(map[string]CachedWeather)

const cacheRate = 15 * time.Second

func GetWeatherFromCache(city string) (weather.WeatherInfo, error) {
	entry, found := cache[city]
	if found && time.Since(entry.Timestamp) < cacheRate {
		return entry.Data, nil
	}

	data, err := weather.GetWeather(city)
	if err != nil {
		return weather.WeatherInfo{}, err
	}
	cache[city] = CachedWeather{
		Data:      data,
		Timestamp: time.Now(),
	}
	return data, nil
}
