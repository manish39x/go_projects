package cmd

import (
	"fmt"
	"os"

	"github.com/manish39x/go-todo/storage"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := parseIndex(args[0])
		if err != nil {
			fmt.Println("❌ Invalid task number.")
			return
		}

		tasks, err := storage.LoadTask()
		if err != nil {
			fmt.Println("❌ Failed to load tasks:", err)
			return
		}

		if index < 1 || index > len(tasks) {
			fmt.Println("❌ Task number out of range.")
			return
		}

		// Delete the task by removing the item from the slice
		tasks = append(tasks[:index-1], tasks[index:]...)

		if err := storage.SaveTasks(tasks); err != nil {
			fmt.Println("❌ Failed to save tasks:", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Task deleted: %s\n", tasks[index-1].Title)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
