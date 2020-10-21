package pull

import (
	"backup-tool/internal/pkg/lib/app"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/spf13/cobra"
)

type App struct{}

var (
	source      string
	destination string
	awsBucket   string
	awsRegion   string
)

func New() app.AppService {
	return &App{}
}

func (a *App) Init(ctx *cobra.Command) error {
	ctx.Flags().StringVarP(&awsRegion, "region", "r", "", "set s3 bucket region")
	ctx.Flags().StringVarP(&awsBucket, "bucket", "b", "", "set s3 bucket name")
	return nil
}

func (a *App) Start(ctx *cobra.Command, args []string) error {

	source := args[0]
	destination := args[1]

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		return err
	}
	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("Unable to open file %q, %v", destination, err)
	}

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(awsBucket),
			Key:    aws.String(source),
		})
	if err != nil {
		return fmt.Errorf("Unable to download item %q, %v", destination, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	return nil
}

func (a *App) Close(ctx *cobra.Command) error {
	return nil
}
