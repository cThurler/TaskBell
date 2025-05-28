package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const taskFile = "tasks.json"

func main() {
	tasks := loadTasks()
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("\nTaskBell")
		fmt.Println("1 - Listar tarefas")
		fmt.Println("2 - Adicionar tarefa")
		fmt.Println("3 - Marcar tarefa como concluída")
		fmt.Println("0 - Sair")

		fmt.Print("Escolha uma opção: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			clearConsole()
			listTasks(tasks)
		case "2":
			clearConsole()
			tasks = addTask(tasks, reader)
			saveTasks(tasks)
		case "3":
			clearConsole()
			tasks = markTaskDone(tasks, reader)
			saveTasks(tasks)
		case "0":
			clearConsole()
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func loadTasks() []Task {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(taskFile, data, 0644)
}

func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa cadastrada.")
		return
	}
	for _, t := range tasks {
		status := "Pendente"
		if t.Done {
			status = "Concluída"
		}
		fmt.Printf("[%d] %s - %s (%s)\n", t.ID, t.Title, t.DueDate.Format("02/01 15:04"), status)
	}
}

func addTask(tasks []Task, reader *bufio.Reader) []Task {
	fmt.Print("Título: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Descrição: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Data (formato: 2006-01-02 15:04): ")
	dateStr, _ := reader.ReadString('\n')
	dateStr = strings.TrimSpace(dateStr)

	due, err := time.Parse("2006-01-02 15:04", dateStr)
	if err != nil {
		fmt.Println("Data inválida. Usando data zero.")
		due = time.Time{}
	}

	newTask := Task{
		ID:          len(tasks) + 1,
		Title:       title,
		Description: description,
		DueDate:     due,
		Done:        false,
	}

	return append(tasks, newTask)
}

func markTaskDone(tasks []Task, reader *bufio.Reader) []Task {
	fmt.Print("ID da tarefa a marcar como concluída: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	var id int
	fmt.Sscanf(idStr, "%d", &id)

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			fmt.Println("Tarefa marcada como concluída.")
			break
		}
	}
	return tasks
}
