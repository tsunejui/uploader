package aws

import (
	"bytes"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	pkg_s3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	session    *session.Session
	downloader *s3manager.Downloader
	client     *pkg_s3.S3
}

func New(awsRegion string) (*S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		return nil, err
	}

	return &S3{
		session:    sess,
		downloader: s3manager.NewDownloader(sess),
		client:     pkg_s3.New(sess),
	}, nil
}

func (s3 *S3) Download(bucket string, source string, file io.WriterAt) (int64, error) {
	fmt.Printf("download...")
	numBytes, err := s3.downloader.Download(file,
		&pkg_s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(source),
		},
	)
	if err != nil {
		return 0, err
	}
	fmt.Printf("done\n")
	return numBytes, nil
}

func (s3 *S3) Upload(bucket string, key string, file []byte) error {
	fmt.Printf("uploading...")
	if _, err := s3.client.PutObject(&pkg_s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),
	}); err != nil {
		return err
	}
	fmt.Printf("done\n")
	return nil
}
