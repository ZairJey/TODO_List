package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
)

type App struct {
	ctx  context.Context
	task []Task
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {}

func (a *App) LoadTasks() []Task {

	var tasks []Task
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("tasks.json not found, initializing with an empty task list.")
			return []Task{} // Возвращаем пустой список, если файла нет
		}
		log.Printf("Error reading tasks file: %v", err)
		return []Task{}
	}

	if len(file) == 0 {
		log.Println("Файл tasks.json пуст, возвращаем пустой список задач.")
		return []Task{}
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		log.Printf("Error unmarshaling tasks: %v", err)
		return []Task{}
	}

	return tasks
}

func (a *App) AddTask(taskName string) {
	tasks := a.LoadTasks()
	tasks = append(tasks, Task{
		Name:   taskName,
		Status: "new",
	})
	a.saveTasks(tasks)
}

func (a *App) DeleteTask(taskName string) {
	tasks := a.LoadTasks()

	// Фильтрация задач, чтобы исключить удаляемую
	var updatedTasks []Task
	for _, task := range tasks {
		if task.Name != taskName {
			updatedTasks = append(updatedTasks, task)
		}
	}
	// Сохранение обновленного списка задач
	a.saveTasks(updatedTasks)
}

func (a *App) saveTasks(tasks []Task) {

	if len(tasks) == 0 {
		if err := os.Remove("tasks.json"); err != nil {
			log.Printf("Error removing tasks.json: %v", err)
		}
		return
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("Error marshaling tasks: %v", err)
		return
	}

	if len(data) == 0 {
		log.Printf("No tasks to save.")
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		log.Printf("Error writing tasks to file: %v", err)
	}
}

func (a *App) CompleteTask(taskName string) {
	tasks := a.LoadTasks()
	for i, task := range tasks {
		if task.Name == taskName {
			tasks[i].Status = "completed"
		}
	}
	a.saveTasks(tasks)
}
