package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"task-manager/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("someting went wrong adding the task", err.Error())
			os.Exit(1)
		}
		if len(args) == 0 {
			fmt.Println("you need to specify the task to add")

		}
		fmt.Printf("Added task \"%s\" to the task list.", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
