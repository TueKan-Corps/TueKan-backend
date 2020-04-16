package thirdparty

import (
	"TueKan-backend/config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

var Sess *session.Session

func InitAWSSession(c *config.Config) error {
	var err error
	Sess, err = session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(c.AmazonID, c.AmazonSecret, ""),
	})
	if err != nil {
		return err
	}

	_, err = Sess.Config.Credentials.Get()
	if err != nil {
		return err
	}

	return nil
}

func ListBuckets() {
	svc := s3.New(Sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
