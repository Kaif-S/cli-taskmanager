package cmd

import "github.com/charmbracelet/lipgloss"

var (
    // Brand colors
    primaryColor = lipgloss.Color("#7D56F4")
    secondaryColor = lipgloss.Color("#04B575")
    ErrorColor = lipgloss.Color("#b51f04")

    // Styles
    TitleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(primaryColor).
        MarginBottom(1)

    TaskStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("252")).
        PaddingLeft(2)

    DoneStyle = lipgloss.NewStyle().
        Strikethrough(true).
        Foreground(lipgloss.Color("240"))

    SuccessBox = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(secondaryColor).
        Padding(0, 1).
        MarginTop(1)

	ErrorBox = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(ErrorColor).
        Padding(0, 1).
        MarginTop(1)
)