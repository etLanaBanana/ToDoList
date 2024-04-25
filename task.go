package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	description string
	completed   bool
	category    *Category
}

func addTask(category *Category, categoryFunc func()) {
	fmt.Println()
	fmt.Println("Для добавления задачи нажмите +, для выбора категорий введите -\nДля возврата в главное меню введите start")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	choiceStr := input.Text()

	if choiceStr == "+" {
		fmt.Println("Введите описание задачи:")
		input.Scan()
		newTask := Task{description: input.Text(), completed: false, category: category}
		category.tasks = append(category.tasks, newTask)
		fmt.Println("Задача добавлена в категорию.")
		showTasks(*category)
	}
	if choiceStr == "start" {
		start(nil)
	}
	if choiceStr == "-" {
		fmt.Println("Выберите категорию:")
		categoryFunc()
		return
	}
	return
}

func showTasks(category Category) {
	fmt.Printf("Список задач в категории %s:\n", category.name)
	for idx, task := range category.tasks {
		status := "✘"
		if task.completed {
			status = "✔"
		}
		fmt.Printf("%d. Описание: %s, Выполнено: %s\n", idx+1, task.description, status)
	}
}
