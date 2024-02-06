package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Create an Amazon S3 service client

	accessKey := ""
	secretKey := ""

	client := s3.New(s3.Options{
		Region:      "us-east-2",
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	})

	response, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("some-bucket"),
	})

	// fmt.Println(response.Contents)
	for _, contents := range response.Contents {
		fmt.Println(*contents.Key)
	}

	fmt.Println(err)

}
