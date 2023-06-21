package cmd

import (
	"context"
	awsClient "eksup/pkg/provider/aws"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/cobra"
	"os"
)

var addonsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List add-ons running on EKS clusters",
	Example: "eksup addons list",
	RunE:    listAddons,
}

func init() {
	addonsCmd.AddCommand(addonsListCmd)
}

func listAddons(cmd *cobra.Command, args []string) error {
	e, c, err := awsClient.Initialize()
	if err != nil {
		fmt.Printf("Error while trying to initialize AWS client. Error: %v\n", err)
	}

	if version != true {
		for _, n := range c {
			fmt.Println("Listing installed add-ons for cluster:", n)

			i := eks.ListAddonsInput{
				ClusterName: &n,
			}

			a, err := e.ListAddons(context.TODO(), &i)
			if err != nil {
				fmt.Println("Couldn't list add-ons for clusters. Error:", err)
			}

			for _, v := range a.Addons {
				fmt.Println(v)
			}

		}
	} else {
		for _, n := range c {
			clusterInput := eks.DescribeClusterInput{
				Name: &n,
			}

			clusterGet, err := e.DescribeCluster(context.TODO(), &clusterInput)
			if err != nil {
				fmt.Println("Couldn't describe current cluster")
			}

			fmt.Println("Listing add-ons for cluster:", n)

			addonsInput := eks.ListAddonsInput{
				ClusterName: &n,
			}

			addonsList, err := e.ListAddons(context.TODO(), &addonsInput)
			if err != nil {
				fmt.Println("Couldn't list add-ons for clusters. Error:", err)
			}

			for _, v := range addonsList.Addons {
				i := eks.DescribeAddonInput{
					ClusterName: &n,
					AddonName:   &v,
				}

				d, err := e.DescribeAddon(context.TODO(), &i)
				if err != nil {
					fmt.Println("Couldn't list add-ons for clusters. Error:", err)
				}

				iu := eks.DescribeAddonVersionsInput{
					AddonName:         &v,
					KubernetesVersion: clusterGet.Cluster.Version,
					MaxResults:        nil,
					NextToken:         nil,
				}
				u, err := e.DescribeAddonVersions(context.TODO(), &iu)
				if err != nil {
					fmt.Println("Error while trying to find versions. Error:", err)
					os.Exit(1)
				}

				for _, w := range u.Addons {
					for k1, v1 := range w.AddonVersions {
						if k1 == 0 {
							if aws.ToString(d.Addon.AddonVersion) == aws.ToString(v1.AddonVersion) {
								fmt.Println(aws.ToString(w.AddonName), "is running version:", styleGreen.Render(aws.ToString(d.Addon.AddonVersion)), "and is up to date.")
							} else {
								fmt.Println(aws.ToString(w.AddonName), "is running version:", styleRed.Render(aws.ToString(d.Addon.AddonVersion)), "and can be upgraded to version:", styleBlue.Render(aws.ToString(v1.AddonVersion)))
							}
						}
					}
				}
			}
		}
	}

	return err
}
