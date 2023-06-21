package cmd

import (
	"github.com/spf13/cobra"
)

var addonsCmd = &cobra.Command{
	Use:     "addons",
	Aliases: []string{"a"},
	Short:   "Describe add-ons installed on EKS clusters",
	RunE:    listAddons,
}

func init() {
	rootCmd.AddCommand(addonsCmd)
}
