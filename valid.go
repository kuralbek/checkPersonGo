package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func validateIIN(iin string) (bool, string) {
	if len(iin) != 12 || !regexp.MustCompile(`^\d{12}$`).MatchString(iin) {
		return false, "IIN должен состоять из 12 цифр"
	}

	weights1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	weights2 := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}

	calculateChecksum := func(iin string, weights []int) int {
		sum := 0
		for i := 0; i < 11; i++ {
			digit, _ := strconv.Atoi(string(iin[i]))
			sum += digit * weights[i]
		}
		return sum % 11
	}

	checksum := calculateChecksum(iin, weights1)
	if checksum == 10 {
		checksum = calculateChecksum(iin, weights2)
	}

	if checksum == 10 || checksum != int(iin[11]-'0') {
		return false, "Неверная контрольная цифра"
	}

	return true, ""
}

func getBirthDateAndGender(iin string) (string, string, error) {
	year := 0
	month := 0
	day := 0
	code := int(iin[6] - '0')

	fmt.Sscanf(iin[0:6], "%02d%02d%02d", &year, &month, &day)

	var century int
	switch code {
	case 1, 2:
		century = 1800
	case 3, 4:
		century = 1900
	case 5, 6:
		century = 2000
	default:
		return "", "", fmt.Errorf("Неверный код")
	}

	year += century

	var gender string
	if code%2 != 0 {
		gender = "male"
	} else {
		gender = "female"
	}

	birthDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if birthDate.IsZero() {
		return "", "", fmt.Errorf("Неверная дата рождения")
	}

	birthDateString := birthDate.Format("02.01.2006")

	return birthDateString, gender, nil
}
