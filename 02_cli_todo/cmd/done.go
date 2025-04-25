package cmd

import (
	"fmt"
	"os"

	"github.com/manish39x/go-todo/storage"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task number]",
	Short: "Mark a task as completed",
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

		tasks[index-1].Completed = true

		if err := storage.SaveTasks(tasks); err != nil {
			fmt.Println("❌ Failed to save tasks:", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Task marked as done: %s\n", tasks[index-1].Title)
	},
}

func parseIndex(input string) (int, error) {
	var index int
	_, err := fmt.Sscanf(input, "%d", &index)
	return index, err
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
