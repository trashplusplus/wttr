package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"wttr/internal/cache"
	"wttr/internal/weather"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Місто (або Enter для авто, q! - вихід):")

	for {
		fmt.Print("> ")
		city, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Помилка вводу:", err)
			continue
		}
		city = strings.TrimSpace(city)

		if city == "q!" {
			fmt.Println("Пака :(")
			return
		}

		weatherInfo, err := cache.GetWeatherFromCache(city)
		if err != nil {
			fmt.Println("Помилка отримання погоди:", err)
			continue
		}

		weather.Show(weatherInfo)
	}
}
