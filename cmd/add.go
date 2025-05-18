package cmd

import (
	"github.com/spf13/cobra"
	"todo-cli/todo"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new todo task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todo.StoreTodo(args[0], "todo.json")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
