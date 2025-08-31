/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/dayemsiddiqui/box-cli/internal/tasks"
	"github.com/spf13/cobra"
)

// Add task subcommand
var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long: `Add a task`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
	
		tasks.CreateTaskAction(name, description)
		fmt.Println("task created")
	},
}

// List task subcommand
var listTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `List all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		taskList,err := tasks.GetAllTasksAction()
		if err != nil {
			fmt.Println("Error listing tasks")
			return
		}
		for _, task := range taskList {
			if task.Description == "" {
				fmt.Printf("ID: %s, Name: %s \n", task.ID, task.Title)
			} else {
			fmt.Printf("ID: %s, Name: %s, Description: %s\n", task.ID, task.Title, task.Description)
			}
		}
		

	},
}

var clearTasksCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all tasks",
	Long: `Clear all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks.DeleteAllTasksAction()
	},
}

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
	Long: `Manage tasks`,
}

func init() {
	rootCmd.AddCommand(taskCmd)

	addTaskCmd.Flags().String("name", "", "Name of the task")
	addTaskCmd.Flags().String("description", "", "Description of the task")

	taskCmd.AddCommand(addTaskCmd)
	taskCmd.AddCommand(listTaskCmd)
	taskCmd.AddCommand(clearTasksCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
