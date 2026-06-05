/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"
	"todo/myutils"
	"github.com/spf13/cobra"
	"github.com/charmbracelet/lipgloss"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [ID] [Title]",
	Args: cobra.ExactArgs(2),
	Short: "Add task by giving id and Title in argument",
	Long: `Add task by giving id and Title in argument`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		id := args[0]
		title := args[1]
		completed := false
		createdAt := time.Now()
		task := myutils.Task{ID: id,Title: title,Completed: completed,CreatedAt: createdAt}

		data,err := myutils.LoadTask()
		
		if err!=nil {
			myutils.SaveTask([]myutils.Task{task})
			out:=SuccessBox.Render(
					lipgloss.JoinVertical(lipgloss.Center,
					fmt.Sprintf("Task Added : %v",title),
				),
				)
			fmt.Println(out)
		}
		data = append(data, task)
		out:=SuccessBox.Render(
					lipgloss.JoinVertical(lipgloss.Center,
					fmt.Sprintf("Task Added : %v",title),
				),
				)
			fmt.Println(out)
		if(len(data)>0){
			for i, t := range data {
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
		myutils.SaveTask(data)

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
