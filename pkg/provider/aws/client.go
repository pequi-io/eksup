package awsClient

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/viper"
	"os"
)

func loadClientConfig() (aws.Config, error) {
	if viper.GetString("aws.auth.credentials") == "true" {
		cfg, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(viper.GetString("aws.region")),
			config.WithRetryMaxAttempts(3),
		)
		if err != nil {
			fmt.Println("Error while trying to authenticate to AWS using credentials. Error:", err)
			os.Exit(1)
		}

		return cfg, err

	} else if viper.GetString("aws.auth.profile") == "true" {
		cfg, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithSharedConfigProfile(viper.GetString("aws.auth.profileName")),
			config.WithRegion(viper.GetString("aws.region")),
			config.WithRetryMaxAttempts(3),
		)
		if err != nil {
			fmt.Println("Error while trying to authenticate to AWS using SSO credentials. Error: ", err)
			os.Exit(1)
		}

		return cfg, err

	} else if viper.GetString("aws.auth.credentials") == viper.GetString("aws.auth.profile") {
		fmt.Println("You can't use both credentials and profile authentication methods at the same time.")
	} else {
		fmt.Println("Couldn't find a specified authentication method.")
		os.Exit(1)
	}

	return aws.Config{}, nil
}

func newEksClient() *eks.Client {
	cfg, err := loadClientConfig()
	if err != nil {
		fmt.Println("Couldn't create a client to EKS service. Error:", err)
	}

	clt := eks.NewFromConfig(cfg)

	return clt
}

var clusters []string

func getClusters(clt *eks.Client) ([]string, error) {
	i := &eks.ListClustersInput{
		Include:    []string{"all"},
		MaxResults: nil,
		NextToken:  nil,
	}

	l, err := clt.ListClusters(context.TODO(), i)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range l.Clusters {
		clusters = append(clusters, v)
	}

	return clusters, err
}

func Initialize() (*eks.Client, []string, error) {
	e := newEksClient()

	c, err := getClusters(e)
	if err != nil {
		fmt.Println("Couldn't get clusters. Error:", err)
	}

	return e, c, err
}
