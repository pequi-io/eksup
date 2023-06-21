package cmd

import (
	"github.com/spf13/cobra"
)

var clustersCmd = &cobra.Command{
	Use:     "clusters",
	Aliases: []string{"c"},
	Short:   "Describe EKS clusters",
}

func init() {
	rootCmd.AddCommand(clustersCmd)
}
