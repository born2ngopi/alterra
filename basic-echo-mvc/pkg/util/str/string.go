package str

import (
	"math/rand"
)

func GenerateRandString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func ConvertDayFromEnToIdn(day string) string {
	switch day {
	case "monday":
		return "senin"
	case "tuesday":
		return "selasa"
	case "wednesday":
		return "rabu"
	case "thursday":
		return "kamis"
	case "friday":
		return "jum'at"
	case "saturday":
		return "sabtu"
	case "sunday":
		return "minggu"
	}
	return ""
}
