package thirdparty

import (
	"TueKan-backend/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"time"
)

type FileItem struct {
	Name         *string    `json:"name"`
	LastModified *time.Time `json:"last_modified"`
	Size         *int64     `json:"size"`
	StorageClass *string    `json:"storage_class"`
}

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

func ListItems() ([]*FileItem, error) {
	svc := s3.New(Sess)
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String("tuekan")})
	if err != nil {
		return nil, err
	}

	fileItems := make([]*FileItem, 0)

	for _, item := range resp.Contents {
		fileItem := new(FileItem)

		fileItem.Name = item.Key
		fileItem.Size = item.Size
		fileItem.LastModified = item.LastModified
		fileItem.StorageClass = item.StorageClass

		fileItems = append(fileItems, fileItem)
	}

	return fileItems, nil
}

func DownloadItem(filename string) error {
	filepath := "./img/" + filename

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(Sess)

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String("tuekan"),
			Key:    aws.String(filename),
		})
	if err != nil {
		return err
	}

	return nil
}

func UploadItem(filename string) error {
	filepath := "./img/" + filename

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	uploader := s3manager.NewUploader(Sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("tuekan"),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}
