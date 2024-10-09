package awsclient

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func FindInstancesByTags(tags []string, region string) []types.Instance {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("Unable to load AWS config: %v", err)
	}

	ec2Svc := ec2.NewFromConfig(cfg)
	var instances []types.Instance
	var filters []types.Filter
	for _, tag := range tags {
		filters = append(filters, types.Filter{
			Name:   aws.String("tag:" + tag),
			Values: []string{"*"},
		})
	}

	result, err := ec2Svc.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		Filters: filters,
	})
	if err != nil {
		log.Fatalf("Error querying EC2 instances: %v", err)
	}

	for _, res := range result.Reservations {
		instances = append(instances, res.Instances...)
	}
	return instances
}
