package main

import (
    "fmt"

    "github.com/ndcampbell/conformitygopher/configs"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
    sess, err := session.NewSessionWithOptions(session.Options{
        Config:  aws.Config{Region: aws.String("us-east-1")},
        Profile: "sandbox",
    })
    if err != nil {
        fmt.Println("Error", err)
    }

    ec2client := ec2.New(sess)
    result, err := ec2client.DescribeInstances(nil)
    if err != nil {
        fmt.Println("Error", err)
    } else {
        fmt.Println("Success", result)
}

    configs.LoadConfigs()
}
