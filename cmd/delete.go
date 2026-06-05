/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"todo/myutils"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Args: cobra.ExactArgs(1),
	Short: "Enter ID for argument to delete that task from tasks",
	Long: `Enter ID for argument to delete that task from tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		tasks,err := myutils.LoadTask()
		if err!=nil{
			fmt.Println("Error occured while loading tasks : ",err)
		}
		deletd := false
		for i:=range(tasks){
			if tasks[i].ID == args[0]{
				tasks = append(tasks[:i],tasks[i+1:]... )
				deletd = true
			}
		}
		if !deletd{
			out := ErrorBox.Render(
					lipgloss.JoinVertical(lipgloss.Center,
					fmt.Sprintf("No task to Delete with id : %v",args[0]),
				),
				)
			fmt.Println(out)
			os.Exit(0)
		}
		out:=SuccessBox.Render(
					lipgloss.JoinVertical(lipgloss.Center,
					fmt.Sprintf("Task Deleted with id : %v",args[0]),
				),
				)
		fmt.Println(out)
		if(len(tasks)>0){
			for i, t := range tasks {
            // 2. Render items with different styles
			var prefix string
			if !t.Completed {
            prefix = lipgloss.NewStyle().Foreground(secondaryColor).Render("○")
			}else{
				prefix = lipgloss.NewStyle().Foreground(secondaryColor).Render("✓")
			}
            content := TaskStyle.Render(fmt.Sprintf("%d. %s", i+1, t.Title))
            
            fmt.Println(prefix + content)
        }
		}
		myutils.SaveTask(tasks)
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
