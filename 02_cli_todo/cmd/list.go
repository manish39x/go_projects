package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/manish39x/go-todo/storage"
	"github.com/spf13/cobra"
)

var filterCompleted bool
var filterPriority string
var filterTag string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.LoadTask()
		if err != nil {
			fmt.Println("❌ Failed to load tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("📭 No tasks found.")
			return
		}

		for i, task := range tasks {
			// Apply filters
			if filterCompleted && !task.Completed {
				continue
			}
			if filterPriority != "" && strings.ToLower(task.Priority) != strings.ToLower(filterPriority) {
				continue
			}
			if filterTag != "" && !contains(task.Tags, filterTag) {
				continue
			}

			// Display
			status := "❌"
			if task.Completed {
				status = "✅"
			}
			fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
			if !task.DueDate.IsZero() {
				fmt.Printf("   📅 Due: %s\n", task.DueDate.Format("2006-01-02"))
			}
			if task.Priority != "" {
				fmt.Printf("   🔥 Priority: %s\n", task.Priority)
			}
			if len(task.Tags) > 0 {
				fmt.Printf("   🏷 Tags: %s\n", strings.Join(task.Tags, ", "))
			}
			fmt.Printf("   🕒 Created: %s\n", task.CreatedAt.Format(time.RFC822))
		}
	},
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.TrimSpace(strings.ToLower(item)) == strings.ToLower(val) {
			return true
		}
	}
	return false
}

func init() {
	listCmd.Flags().BoolVar(&filterCompleted, "completed", false, "Only show completed tasks")
	listCmd.Flags().StringVar(&filterPriority, "priority", "", "Filter by priority (low, medium, high)")
	listCmd.Flags().StringVar(&filterTag, "tag", "", "Filter by tag")
	rootCmd.AddCommand(listCmd)
}
