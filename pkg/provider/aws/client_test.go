package awsClient

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/viper"
	"testing"
)

func TestLoadClientConfig(t *testing.T) {
	v := viper.New()
	v.SetConfigFile("../../../testdata/config-test.yaml")
	err := v.ReadInConfig()
	if err != nil {
		t.Fatalf("Error while reading the config file: %v", err)
	}

	credentials := v.GetString("aws.auth.credentials")

	expectedCredentials := "true"
	if credentials != expectedCredentials {
		t.Errorf("Incorrect value for the 'aws.auth.credentials' key. Expected: %s, Got: %s", expectedCredentials, credentials)
	}

	profile := v.GetString("aws.auth.profile")

	expectedProfile := "false"
	if profile != expectedProfile {
		t.Errorf("Incorrect value for the 'aws.auth.credentials' key. Expected: %s, Got: %s", expectedCredentials, credentials)
	}

	if credentials == profile {
		t.Errorf("Only one authentication method can be used.")
	}
}

func TestNewEksClient(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		t.Fatalf("Error while loading the default config: %v", err)
	}

	clt := eks.NewFromConfig(cfg)

	if clt == nil {
		t.Errorf("Couldn't create a client to EKS service.")
	}
}

type EKSClientInterface interface {
	ListClusters(cfg aws.Config, params *eks.ListClustersInput, optFns ...func(*eks.Options)) (*eks.ListClustersOutput, error)
}
