/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task-manager/db"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("failed to parse the argument")
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.ReadTasks()
		if err != nil {
			fmt.Println("something went wrong", err)
			return
		}
		_ = tasks
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
