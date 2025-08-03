package weather

import (
	"testing"
)

func TestDesc(t *testing.T) {

	cities := []string{
		"Неаполь",
		"Полтава",
		"Пномпен",
		"Пекин",
		"Орявчик",
		"Львів",
		"Мілан",
		"Багдад",
		"Тегеран",
		"Львів",
		"Берлін",
		"Краків",
		"Ріга",
		"Мехіко",
		"Сінгапур",
		"Гельсінки",
		"Магдебург",
		"Берн",
		"Базель",
		"Харків",
		"Женева",
		"Афіни",
		"Делі",
		"Сідней",
		"Мельбурн",
		"Тайбей",
	}
	for _, city := range cities {
		weatherInfo, err := GetWeather(city)
		if err != nil {
			t.Error(err)
		}
		ShowLocaledDesc(weatherInfo)
	}

}
