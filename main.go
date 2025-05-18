package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Введите первую дату (в формате ДД.ММ.ГГ)")
	date1 := readDate()

	fmt.Println("Введите вторую дату (в формате ДД.ММ.ГГ)")
	date2 := readDate()

	days := calculateDaysBetween(date1, date2)
	years := float64(days) / 365
	month := int(years * 12)
	weeks := float64(days / 7)
	fmt.Println("Между датами", days, "дней")
	fmt.Printf("Разница в неделях: %.1f\n", weeks)
	fmt.Println("Разница в месяцах:", month)
	fmt.Printf("Разница в годах: %.1f\n", years)
}

func readDate() time.Time {
	var input string
	fmt.Scan(&input)
	layout := "02.01.2006"
	parsedDate, err := time.Parse(layout, input)

	if err != nil {
		fmt.Println("Неверный формат даты. Попробуйте снова")
		return readDate()
	}
	return parsedDate
}

func calculateDaysBetween(d1, d2 time.Time) int {
	if d1.After(d2) {
		d1, d2 = d2, d1
	}

	diff := d2.Sub(d1)
	return int(diff.Hours() / 24)
}
