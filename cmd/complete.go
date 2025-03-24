/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "os"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <task_id>",
	Short: "Mark a task as complete",
	Long: `A longer description that spans multiple lines`,
    
	Run: func(cmd *cobra.Command, args []string) {
        if(len(args) == 0) {
            fmt.Fprintln(os.Stderr, "Error: Task Id is required\nUsage: complete <task_id>")
            return
        }
        taskId := args[0]
		fmt.Println("process with process id", taskId, "is completed")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
