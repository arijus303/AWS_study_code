package main

import (
	"fmt"
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	// Specify your AWS region
	region := "us-east-1"

	// Specify your SQS queue URL
	queueURL := "https://sqs.us-east-1.amazonaws.com/549565409924/testq"

	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create an SQS service client
	sqsClient := sqs.New(sess)

	// Receive messages from the queue
	receiveParams := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(10),     // Adjust as needed
		VisibilityTimeout:   aws.Int64(1),      // Adjust as needed
		WaitTimeSeconds:     aws.Int64(3),     // Adjust as needed
		MessageAttributeNames: []*string{
			aws.String("All"),
		},
	}

	// Receive messages
	result, err := sqsClient.ReceiveMessage(receiveParams)
	if err != nil {
		fmt.Println("Error receiving messages:", err)
		return
	}

	if len(result.Messages) > 0 {

	// Print received messages
	fmt.Println("Received Messages:")
	for _, message := range result.Messages {
		fmt.Println(*message.Body)
	}

	fmt.Println("Simulating workload...")
	time.Sleep(time.Second * 5)
	// Note: In a real-world scenario, you would typically process and handle the received messages here.
	} else {
		fmt.Println("No Messages in SQS")
	}
}

