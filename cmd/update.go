/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo/myutils"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update task to true of false with id",
	Long: `Update task to true of false with id`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		tasks,err := myutils.LoadTask()
		if len(tasks)==0 {
			out := SuccessBox.Render(
					lipgloss.JoinVertical(lipgloss.Center,
					"No tasks to update",
				),
				)
			fmt.Println(out)
		}
		var task myutils.Task
		if err!=nil{
			fmt.Println("Error occured while getting tasks from json file : ",err)
		}
		updated := false
		for t := range(len(tasks)){
			if tasks[t].ID==args[0] {
				task = tasks[t]
				tasks[t].Completed = true
				message := fmt.Sprintf("Finished task : %v",task.Title)
				out:=SuccessBox.Render(
					lipgloss.JoinVertical(lipgloss.Center,
					"🎉 Nice Job! ",message,
				),
				)
				fmt.Println(out)
				updated=true
			}
		}
		if !updated{
			out := ErrorBox.Render(
					lipgloss.JoinVertical(lipgloss.Center,
					fmt.Sprintf("No task to update with id : %v",args[0]),
				),
				)
			fmt.Println(out)
		}else{
			myutils.SaveTask(tasks)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
