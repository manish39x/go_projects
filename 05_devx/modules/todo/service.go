package todo

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
	"time"

	"github.com/google/uuid"
)

var StorageFile = "todos.json"

// LoadTasks reads all tasks from JSON file
func LoadTasks() ([]Task, error) {
	var tasks []Task

	file, err := os.ReadFile(StorageFile)
	if errors.Is(err, os.ErrNotExist) {
		return tasks, nil
	} else if err != nil {
		return nil, err
	}

	if len(file) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func ListTasks() ([]Task, error) {
	return LoadTasks()
}

// SaveTasks writes all tasks to JSON file
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(StorageFile, data, 0644) //
}

// AddTask create and saves a new task
func AddTask(title string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	return SaveTasks(tasks)
}

func MarkDone(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if index < 1 || index > len(tasks) {
		return errors.New("Invalid task number")
	}

	tasks[index-1].Completed = true
	return SaveTasks(tasks)
}

func DeleteTask(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if index < 1 || index > len(tasks) {
		return errors.New("Invalid task number")
	}

	// Remove task at index-1
	tasks = slices.Delete(tasks, index-1, index)
	return SaveTasks(tasks)
}

func ClearTasks() error {
	var tasks []Task

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(StorageFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
