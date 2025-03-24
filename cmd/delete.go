/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <task_id>",
	Short: "Delete a live task",
	Long: `A longer description that spans multiple lines`,

	Run: func(cmd *cobra.Command, args []string) {
        if(len(args) == 0){
            fmt.Fprintln(os.Stderr, "Error: Task ID is required\nUsage: delete <task_id>")
            return
        }
        taskId := args[0]
        fmt.Printf("task with task id %s is deleted\n", taskId)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
