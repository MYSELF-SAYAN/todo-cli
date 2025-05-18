package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI Todo application",
	Long:  `Manage your daily tasks right from the command line.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Use `todo --help` to see available commands")
	// },
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
