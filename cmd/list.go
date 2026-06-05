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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all of your tasks",
	Long: `list all of your tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(TitleStyle.Render(" My Todo List "))
		tasks,err := myutils.LoadTask()
		if (err!=nil || len(tasks)==0){
			fmt.Println("no tasks here")
		}
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
	},
}
func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
