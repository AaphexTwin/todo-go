package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var tasksFile = "test.json"

func readLine(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

//var reader = bufio.NewReader(os.Stdin)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var nextID int = 1
var tasks []Task

func printInfo() {
	fmt.Printf("\nВыберите нужный пункт:\n1. Добавить задачу\n2. Показать задачи\n3. Отметить выполненной\n4. Удалить задачу\n5. Выход\n")
}

func loadTasksFromJson() {
	data, err := os.ReadFile(tasksFile)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл не найден")
			return
		}
		fmt.Printf("Не удалось прочитать файл %s: %v\n", tasksFile, err)
		return
	}

	if len(data) == 0 {
		nextID = 1
		fmt.Println("Файл пуст")
		return
	}

	var loadedTasks []Task
	err = json.Unmarshal(data, &loadedTasks)
	if err != nil {
		fmt.Printf("Файл: %s содержит некорректный JSON формат %v\n", tasksFile, err)
		return
	}

	tasks = loadedTasks
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	nextID = maxID + 1

}

func saveTasksToJSON() {
	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		fmt.Printf("Ошибка формирования JSON: %v\n", err)
		return
	}
	err = os.WriteFile(tasksFile, data, 0644)

	if err != nil {
		fmt.Printf("Не удалось сохранить файл %s: %v\n", tasksFile, err)
		return
	}
	fmt.Println("Задачи сохранены")
}

func addTask() {

	task := readLine("Введите задачу: ")

	if task == "" {
		fmt.Println("❌ Название задачи не может быть пустым")
		return
	}
	newTask := Task{
		ID:    nextID,
		Title: task,
		Done:  false,
	}
	nextID++
	tasks = append(tasks, newTask)
	fmt.Println("Задача добавлена")
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Список пуст")
		return
	}
	for _, task := range tasks {
		if task.Done {
			fmt.Printf("[✓] %d: %s\n", task.ID, task.Title)
		} else {
			fmt.Printf("[ ] %d: %s\n", task.ID, task.Title)
		}
	}
}

// изделия из дерева
func isDone() {
	fmt.Print("Введите ID выполненной задачи: ")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	var newID int
	n, err := fmt.Sscanf(line, "%d", &newID)
	if err != nil || n != 1 {
		fmt.Println("Некорректный ID (нужно целое число)")
		return
	}

	for i := range tasks {
		if tasks[i].ID == newID {
			tasks[i].Done = true
			fmt.Printf("[✓] Задача %d: %s выполнена\n", tasks[i].ID, tasks[i].Title)
			return
		}
	}
	fmt.Printf("Задача с ID %d не найдена\n", newID)
}

func deleteTask() {
	fmt.Print("Введите ID задачи которую хотите удалить: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	var newID int
	n, err := fmt.Sscanf(line, "%d", &newID)
	if err != nil || n != 1 {
		fmt.Println("Некорректный ID (нужно целое число)")
		return
	}

	if newID <= 0 {
		fmt.Println("ID должен быть положительным")
		return
	}
	for i := range tasks {
		if tasks[i].ID == newID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Задача %d удалена\n", newID)
			return
		}
	}
	fmt.Printf("Задача с ID %d не найдена\n", newID)
}

func main() {

	loadTasksFromJson()
	var choice int

	for {
		printInfo()
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода")
			continue
		}

		_, err = fmt.Sscanf(strings.TrimSpace(line), "%d", &choice)
		if err != nil {
			fmt.Println("Введите число от 1 до 5")
			continue
		}
		switch choice {
		case 1:
			fmt.Println("Добавление")
			addTask()
			saveTasksToJSON()
		case 2:
			fmt.Println("Список")
			showTasks()
		case 3:
			fmt.Println("Выполнено")
			isDone()
			saveTasksToJSON()
		case 4:
			fmt.Println("Удаление")
			deleteTask()
			saveTasksToJSON()
		case 5:
			fmt.Println("Выход")
			saveTasksToJSON()
			return
		default:
			fmt.Println("Такого пункта не существует")
		}
		fmt.Println()
	}
}
