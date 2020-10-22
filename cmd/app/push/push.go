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
	zipFile     bool
	rename      string
)

func New() app.AppService {
	return &App{}
}

func (a *App) Init(ctx *cobra.Command) error {
	ctx.Flags().StringVarP(&awsRegion, "region", "", "", "set s3 bucket region")
	ctx.Flags().StringVarP(&awsBucket, "bucket", "", "", "set s3 bucket name")
	ctx.Flags().StringVarP(&rename, "rename", "", "", "set the zip file name")
	ctx.Flags().BoolVarP(&zipFile, "zip", "", false, "type ture or false. if true, and the source arguments is a correct folder or file path, it will compress to a zip file and upload")
	return nil
}

func (a *App) Start(ctx *cobra.Command, args []string) error {
	source := args[0]
	destination := args[1]

	var path string
	var err error
	if zipFile {
		path = getZipName(source)
		err = file.Zip(source, path)
	} else {
		path = source
	}

	if err != nil {
		return err
	}

	s3Svr, err := aws_tool.New(awsRegion)
	if err != nil {
		return err
	}

	buffer, err := file.ReadFileBuffer(path)
	if err != nil {
		return err
	}

	if err := s3Svr.Upload(awsBucket, destination, buffer); err != nil {
		return err
	}

	if zipFile {
		err = file.RemoveFile(path)
	}

	if err != nil {
		return err
	}

	return nil
}

func (a *App) Close(ctx *cobra.Command) error {
	return nil
}

func getZipName(source string) string {
	if len(rename) != 0 {
		return rename + ".zip"
	}

	var zipName string
	if file.IsDir(source) {
		zipName = file.GetFolderName(source, true)
	} else {
		zipName = file.GetFileName(source)
	}
	return zipName + ".zip"
}
