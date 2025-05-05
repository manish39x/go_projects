package tests

import (
	"path/filepath"
	"testing"

	"github.com/manish39x/devx/modules/todo"
)

func setupTestFile(t *testing.T) string {
	dir := t.TempDir()
	testFile := filepath.Join(dir, "test_todos.json")
	todo.StorageFile = testFile
	return testFile
}

func TestAddAndListTodos(t *testing.T) {
	setupTestFile(t)

	if err := todo.AddTask("Test task 1"); err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	tasks, err := todo.ListTasks()
	if err != nil {
		t.Fatalf("LoadTask failed: %v ", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Title != "Test task 1" {
		t.Errorf("Expected title 'Test task 1', got '%s'", tasks[0].Title)
	}
}

func TestMarkDone(t *testing.T) {
	setupTestFile(t)

	_ = todo.AddTask("Test task")
	_ = todo.MarkDone(1)

	tasks, _ := todo.ListTasks()
	if !tasks[0].Completed {
		t.Errorf("Task was not marked as completed")
	}
}

func TestDeleteTask(t *testing.T) {
	setupTestFile(t)

	_ = todo.AddTask("Task to delete")
	_ = todo.DeleteTask(1)

	tasks, _ := todo.ListTasks()
	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks after deletion, got %d", len(tasks))
	}
}
