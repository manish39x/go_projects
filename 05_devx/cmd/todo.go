package cmd

import (
	"fmt"
	"strconv"

	"github.com/manish39x/devx/modules/todo"
	"github.com/manish39x/devx/tui"
	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage your to-do list",
}

var addCmd = &cobra.Command{
	Use:   "add [task title]",
	Short: "Add a new to-do task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		if err := todo.AddTask(title); err != nil {
			fmt.Println("❌ Error: ", err)
			return
		}
		fmt.Println("✅ Task added: ", title)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all to-do tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := todo.LoadTasks()
		if err != nil {
			fmt.Println("❌ Error loading tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("🟢 Your task list is empty. You're all caught up!")
			return
		}

		fmt.Println("📋 Your Tasks:")
		for index, task := range tasks {
			status := "[ ]"
			if task.Completed {
				status = "[x]"
			}
			fmt.Printf("%d. %s %s \n", index+1, status, task.Title)
		}
	},
}

var doneCmd = &cobra.Command{
	Use:   "done [task number]",
	Short: "Mark a to-do as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Please enter a valid task number.")
			return
		}
		err = todo.MarkDone(index)
		if err != nil {
			fmt.Println("❌ Error: ", err)
			return
		}

		fmt.Printf("✅ Task %d marked as complete.\n", index)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a to-do from your list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Please enter a valid task number.")
			return
		}

		if err = todo.DeleteTask(index); err != nil {
			fmt.Println("❌ Error:", err)
			return
		}

		fmt.Printf("🗑️ Task %d deleted.\n", index)

	},
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all todos",
	Long:  `This command removes all tasks from the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.ClearTasks(); err != nil {
			fmt.Println("❌Error clearing tasks:", err)
			return
		}
		fmt.Println("🗑️ All tasks have been cleared")
	},
}

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch DevDash terminal dashboard UI",
	Run: func(cmd *cobra.Command, args []string) {
		tui.Start()
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
	rootCmd.AddCommand(tuiCmd)
	todoCmd.AddCommand(addCmd)
	todoCmd.AddCommand(listCmd)
	todoCmd.AddCommand(doneCmd)
	todoCmd.AddCommand(deleteCmd)
	todoCmd.AddCommand(clearCmd)
}
