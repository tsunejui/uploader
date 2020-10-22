package push

import (
	aws_tool "backup-tool/cmd/app/utils/aws"
	"backup-tool/cmd/app/utils/file"
	"backup-tool/internal/pkg/lib/app"

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
	ctx.Flags().StringVarP(&awsRegion, "region", "", "", "set s3 bucket region")
	ctx.Flags().StringVarP(&awsBucket, "bucket", "", "", "set s3 bucket name")
	return nil
}

func (a *App) Start(ctx *cobra.Command, args []string) error {
	source := args[0]
	destination := args[1]

	s3Svr, err := aws_tool.New(awsRegion)
	if err != nil {
		return err
	}

	buffer, err := file.ReadFileBuffer(source)
	if err != nil {
		return err
	}

	if err := s3Svr.Upload(awsBucket, destination, buffer); err != nil {
		return err
	}

	return nil
}

func (a *App) Close(ctx *cobra.Command) error {
	return nil
}
