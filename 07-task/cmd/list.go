package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"../db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil{  
			fmt.Println("Something Went Worng:",err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
			return
		}
		fmt.Println("You have following tasks:")
		for i,task := range tasks {
			fmt.Printf("%d. %s,Key=%d\n",i+1,task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}