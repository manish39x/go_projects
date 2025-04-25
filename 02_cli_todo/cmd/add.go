package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/manish39x/go-todo/models"
	"github.com/manish39x/go-todo/storage"
	"github.com/spf13/cobra"
)

var due string
var priority string
var tags string

var addCmd = &cobra.Command{
	Use:   "add [task title]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")

		dueDate := time.Time{}
		if due != "" {
			parsed, err := time.Parse("2006-01-02", due)
			if err != nil {
				fmt.Println("❌ Invalid due date format. Use YYYY-MM-DD.")
				return
			}
			dueDate = parsed
		}

		task := models.Task{
			ID:        uuid.New().String(),
			Title:     title,
			Completed: false,
			CreatedAt: time.Now(),
			DueDate:   dueDate,
			Priority:  priority,
			Tags:      strings.Split(tags, ","),
		}

		tasks, _ := storage.LoadTask()
		tasks = append(tasks, task)
		storage.SaveTasks(tasks)

		fmt.Println("✅ Task added:", task.Title)
	},
}

func init() {
	addCmd.Flags().StringVar(&due, "due", "", "Due date in YYYY-MM-DD")
	addCmd.Flags().StringVar(&priority, "priority", "", "Priority level (low, medium, high)")
	addCmd.Flags().StringVar(&tags, "tags", "", "Comma-separated tags")
	rootCmd.AddCommand(addCmd)
}
