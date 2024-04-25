package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Category struct {
	name  string
	tasks []Task
}

func showCategories(categories map[int]Category) {
	for key, value := range categories {
		fmt.Printf("%d: %s\n", key, value.name)
	}
}

func category() {
	var selectedCategory Category
	categories := make(map[int]Category)
	categories[1] = Category{name: "Важные", tasks: []Task{
		{description: "Позвонить маме", completed: false},
	}}
	categories[2] = Category{name: "Работа", tasks: []Task{
		{description: "Подготовить отчет", completed: false},
		{description: "Подготовиться к совещанию", completed: false}}}
	categories[3] = Category{name: "Покупки", tasks: []Task{
		{description: "Купить продукты", completed: false},
	}}
	categories[4] = Category{name: "Дом", tasks: []Task{}}

	showCategories(categories)
	fmt.Println("Введите номер категории или введите '+' для создания новой:\nДля возврата в главное меню введите start")
	fmt.Println()
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	choiceStr := input.Text()

	if choiceStr == "+" {
		fmt.Println("Введите название новой категории:")
		input.Scan()
		newCategoryName := input.Text()
		newCategory := Category{name: newCategoryName, tasks: []Task{}}
		maxKey := 0
		for key := range categories {
			if key > maxKey {
				maxKey = key
			}
		}
		categories[maxKey+1] = newCategory
		fmt.Println("Новая категория добавлена.")
		for key, value := range categories {
			fmt.Printf("%d: %s\n", key, value.name)
		}

		for {
			addTask(&selectedCategory, category)
		}
	}
	if choiceStr == "start" {
		start(nil)
	}

	choice, err := strconv.Atoi(choiceStr)
	if err != nil {
		fmt.Println("Ошибка при чтении ввода.")
		return
	}

	if choice < 1 || choice > len(categories) {
		fmt.Println("Неверный выбор категории.")
		return
	} else {
		selectedCategory = categories[choice]
		fmt.Printf("Выбранная категория: %s\n", selectedCategory.name)
	}
	addTask(&selectedCategory, category)
}
