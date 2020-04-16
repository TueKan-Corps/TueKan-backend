package util

import (
	"TueKan-backend/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func InitAWSSession(c *config.Config) error {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(c.AmazonID, c.AmazonSecret, ""),
	})
	if err != nil {
		return err
	}

	_, err = sess.Config.Credentials.Get()
	if err != nil {
		return err
	}

	return nil
}
