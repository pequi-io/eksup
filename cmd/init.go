package cmd

import (
	"eksup/pkg/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Generates a .eksup.yaml config file",
	Example: "eksup init",
	RunE:    runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, _ []string) error {
	home, _ := os.UserHomeDir()
	file := "/.eksup.yaml"
	path := home + file

	create, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		fmt.Println("Config file does not exist, creating it...")
	}
	defer create.Close()
	fmt.Printf("Generating eksup config file at %s\n", path)

	_, err = create.WriteString(config.SampleCfg)
	if err != nil {
		return err
	}

	fmt.Println("Config file created. Edit it in order to configure the AWS authentication method.")
	return nil
}
