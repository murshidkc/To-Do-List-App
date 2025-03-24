/*
Copyright © 2025 Muhammed Murshid KC <mrmurshidkc@gmail.com>

*/
package cmd

import (
	"fmt"
    "os"
    "github.com/murshidkc/To-Do-List-App/internals"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <task description>",
	Short: "Add a new task",
	Long: `A longer description that spans multiple lines`,

	Run: func(cmd *cobra.Command, args []string) {
        if(len(args) == 0){
            fmt.Fprintln(os.Stderr, "Error: Task description is required\nUsage: Add <task description>")
            return
        }
        description := args[0]
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
