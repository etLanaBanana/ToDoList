package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var taskDisplayed bool
var categoryDisplayed bool

func start(categories map[int]Category) {

	for {
		fmt.Println("Введите 1, если хотите добавить напоминание для конкретного дня,2 если хотите создать задачу, или 3 для просмотра задач")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == "1" {
			fmt.Println()
			calendar()
			fmt.Println()
			fmt.Println("Выберите день")
			for {
				input = bufio.NewScanner(os.Stdin)
				input.Scan()
				dayStr := input.Text()
				day, err := strconv.Atoi(dayStr)
				if err != nil {
					fmt.Println("Ошибка при чтении дня.")
					continue
				}
				currentTime := time.Now()
				year, month, _ := currentTime.Date()
				firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
				lastDay := firstDay.AddDate(0, 1, -1)

				if !validateDay(day, firstDay, lastDay) {
					fmt.Println("Неверный день. Введите число от", firstDay.Day(), "до", lastDay.Day())
					continue
				}
				handleSelectedDay(day)
				fmt.Println("Введите напоминание:")
				input.Scan()
				reminder := input.Text()
				fmt.Println("Напоминание", reminder, "добавлено.")
				fmt.Println()
				break
			}

		} else if input.Text() == "2" {
			for {
				fmt.Println()
				fmt.Println("Выберите категорию:")
				category()
			}
		} else if input.Text() == "3" {
			fmt.Println("1. Просмотр всех задач\n2. Выбрать категорию")
			input.Scan()
			for {
				if input.Text() == "1" {
					if !taskDisplayed {

						taskDisplayed = true
					}
				} else if input.Text() == "2" {
					if !categoryDisplayed {
						showCategories(categories)
						categoryDisplayed = true
					}
				} else {
					fmt.Println("Ошибка при чтении ввода.")
					return
				}
			}
		} else {
			fmt.Println("Введите 1, 2 или 3!")
		}
	}

}
