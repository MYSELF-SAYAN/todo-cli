package cmd

import (
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all todo tasks",
	Long:  `Display all your todo tasks in a formatted table.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Call the GetTodos function with the file path
		todo.GetTodos("todo.json")
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
