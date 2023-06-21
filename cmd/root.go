package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "eksup",
	Short: `
	███████╗██╗  ██╗███████╗██╗   ██╗██████╗ 
	██╔════╝██║ ██╔╝██╔════╝██║   ██║██╔══██╗
	█████╗  █████╔╝ ███████╗██║   ██║██████╔╝
	██╔══╝  ██╔═██╗ ╚════██║██║   ██║██╔═══╝ 
	███████╗██║  ██╗███████║╚██████╔╝██║     
	╚══════╝╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═╝     
	                                        
         EKS upgrading made easy
`,
}

var (
	version  bool
	styleRed = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff4765"))
	styleGreen = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#66ff33"))
	styleBlue = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00b4ff"))
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eksup.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".eksup")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}
