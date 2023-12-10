package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func createSession() *session.Session {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewCredentials(&credentials.EnvProvider{}),
		},
	})
	if err != nil {
		log.Fatalf("FAILED CREATE SHARED SESSION: %v", err)
	}

	return sess
}

func createEC2Instance(sess *session.Session) {
	var (
		region                   = "us-east-1"             // N.Virginia
		imageId                  = "ami-0230bd60aa48260c6" // Amazon Linux 2023 AMI
		instanceType             = "t2.micro"
		keyPairName              = "mymachinekey" // generate key-pair .pem file for ssh access
		securityGroupName        = "launch-wizard-2"
		minCount, maxCount int64 = 1, 1
	)

	instance := ec2.New(sess, &aws.Config{
		Region: &region,
	})

	runningInstance, err := instance.RunInstances(&ec2.RunInstancesInput{
		ImageId:        &imageId,
		InstanceType:   &instanceType,
		MinCount:       &minCount,
		MaxCount:       &maxCount,
		KeyName:        &keyPairName,
		SecurityGroups: []*string{&securityGroupName},
	})
	if err != nil {
		log.Fatalf("FAILED CREATE EC2 INSTANCES IN REGION %s, IMAGEID %s and INSTANCE TYPE %s: %v", region, imageId, instanceType, err)
	}

	time.Sleep(2 * time.Minute)

	_, err = instance.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: []*string{runningInstance.Instances[0].InstanceId},
	})
	if err != nil {
		log.Fatalf("FAILED STOP INSTANCE WITH ID %s", *runningInstance.Instances[0].InstanceId)
	}

	log.Printf("SUCCESS STOP INSTANCE WITH ID %s", *runningInstance.Instances[0].InstanceId)
}

func main() {
	sess := createSession()
	createEC2Instance(sess)
}
