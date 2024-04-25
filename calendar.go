package main

import (
	"fmt"
	"time"
)

func validateDay(day int, firstDay time.Time, lastDay time.Time) bool {
	if day < firstDay.Day() || day > lastDay.Day() {
		return false
	}
	return true
}
func handleSelectedDay(day int) {
	currentTime := time.Now()
	year, month, _ := currentTime.Date()
	selectedDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	formattedDate := selectedDate.Format("2 January 2006 года")
	fmt.Println()
	fmt.Printf("Выбранная дата: %s\n", formattedDate)
}

func calendar() {
	currentTime := time.Now()
	year, month, _ := currentTime.Date()

	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)

	fmt.Println("Calendar for", month, year)
	fmt.Println("Mon Tue Wed Thu Fri Sat Sun")

	startDay := int(firstDay.Weekday())
	if startDay == 0 {
		startDay = 7
	}

	for i := 1; i < startDay; i++ {
		fmt.Print("    ")
	}

	for d := firstDay; !d.After(lastDay); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Monday && d.Day() != 1 {
			fmt.Println()
		}

		day := d.Day()

		if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
			fmt.Printf("\x1b[1;33m%3d\x1b[0m ", day)
		} else {
			fmt.Printf("%3d ", day)
		}
	}
	fmt.Println()
}
func printCalendarWithCircle(selectedDay int) {
	currentTime := time.Now()
	year, month, _ := currentTime.Date()

	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)

	fmt.Println("Calendar for", month, year)
	fmt.Println("Mon Tue Wed Thu Fri Sat Sun")

	startDay := int(firstDay.Weekday())
	if startDay == 0 {
		startDay = 7
	}

	for i := 1; i < startDay; i++ {
		fmt.Print("    ")
	}

	for d := firstDay; !d.After(lastDay); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Monday && d.Day() != 1 {
			fmt.Println()
		}

		day := d.Day()
		if d.Day() == selectedDay {
			fmt.Printf("\033[1;31m%3d\033[0m ", d.Day())
		} else if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
			fmt.Printf("\x1b[1;33m%3d\x1b[0m ", day)
		} else {
			fmt.Printf("%3d ", d.Day())
		}

	}
}
