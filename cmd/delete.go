package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task]",
	Short: "Delete a todo task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Convert the first argument to an integer
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error parsing ID: %v\n", err)
			return
		}
		// Call the DeleteTodo function with the ID and file path
		todo.DeleteTodo(id, "todo.json")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
