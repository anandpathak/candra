package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// ServerList contains requires ec2 instance details
type ServerList struct {
	Name     string
	PublicIP string
	PemKey   string
}

// GetServersList provides list of aws server based on the filter params
func GetServersList(config aws.Config, tag string, value string) []ServerList {
	var serverList []ServerList
	sess, err := session.NewSession(&config)
	svc := ec2.New(sess)
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(tag),
				Values: []*string{
					aws.String(value),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}
	for _, reservation := range result.Reservations {
		var s1 ServerList
		for _, instance := range reservation.Instances {
			//check if instance is in  running state
			if *instance.State.Code != 16 {
				break
			}
			for _, TAG := range instance.Tags {
				if *(TAG.Key) == "Name" {
					s1.Name = *(TAG.Value)
				}
			}

			if instance.PublicIpAddress != nil {
				s1.PublicIP = *instance.PublicIpAddress
			} else {
				s1.PublicIP = *instance.PrivateIpAddress
			}
			s1.PemKey = *(instance.KeyName)
			serverList = append(serverList, s1)
		}

	}
	return serverList
}
