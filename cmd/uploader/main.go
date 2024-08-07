package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (

	s3client *s3.S3
	s3bucket string

)


func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Endpoint: aws.String("http://localhost:4566"),
			S3ForcePathStyle: aws.Bool(true),
			Credentials: credentials.NewStaticCredentials(
				"teste",
				"teste",
				"",	
			),
		},
	)
	if err != nil {
		panic(err)
		
	}
	s3client = s3.New(sess)
	s3bucket = "bucket-mensagens"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)		
	}
	defer dir.Close()

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		uploadFile(files[0].Name())
	}

}

func uploadFile(fimename string ) {
	completeFileName := fmt.Sprintf("./tmp/%s", fimename)
	fmt.Printf("Uploading file %s to bucket\n", completeFileName)
	
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFileName)
		return
	}
	defer f.Close()
	_, err = s3client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3bucket),
		Key: aws.String(fimename),
		Body: f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", err)
		return
	}
	fmt.Printf("File %s uploaded\n", completeFileName)


}