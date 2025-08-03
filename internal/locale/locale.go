package locale

var wordlist = map[string]string{
	"Patchy rain nearby":     "Поруч невеликий дощ",
	"Rain with thunderstorm": "Дощ з грозою",
	"Clear":                  "Ясно",
	"Partly cloudy":          "Мінлива хмарність",
	"Smoke":                  "Дим",
	"Sunny":                  "Сонячно",
	"Haze":                   "Туман, димка",
	"Light rain shower":      "Легкий дощ",
	"Mist":                   "Туман",
	"Light drizzle":          "Мигичка",
	"Cloudy":                 "Хмарно",
	"Overcast":               "Похмуро",
}

func Translate(word string) string {

	if wordlist[word] != "" {
		return wordlist[word]
	}

	return word
}
