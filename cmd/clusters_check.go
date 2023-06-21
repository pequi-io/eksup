package cmd

import (
	"context"
	"eksup/pkg/getver"
	awsClient "eksup/pkg/provider/aws"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/cobra"
)

var clustersCheckCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"c"},
	Short:   "Check for cluster version upgrades",
	Example: "eksup clusters check",
	RunE:    checklistClusters,
}

func init() {
	clustersCmd.AddCommand(clustersCheckCmd)
}

func checklistClusters(cmd *cobra.Command, args []string) error {

	e, c, err := awsClient.Initialize()
	if err != nil {
		fmt.Printf("Error while trying to initialize AWS client. Error: %v\n", err)
	}

	fmt.Println("List of clusters:")

	versions, err := getver.GetVersion()
	if err != nil {
		fmt.Println("Couldn't get versions. Error:", err)
	}

	for _, cluster := range c {
		i := eks.DescribeClusterInput{
			Name: &cluster,
		}

		a, err := e.DescribeCluster(context.TODO(), &i)
		if err != nil {
			fmt.Println("error")
		}

		if aws.ToString(a.Cluster.Version) != versions[0] {
			fmt.Println(aws.ToString(a.Cluster.Name), "is running version:", styleRed.Render(aws.ToString(a.Cluster.Version)), "and can be upgraded to version:", styleBlue.Render(versions[0]))
		} else {
			fmt.Println(aws.ToString(a.Cluster.Name), "is running version:", styleGreen.Render(aws.ToString(a.Cluster.Version)), "and is up to date")
		}
	}

	return err
}
