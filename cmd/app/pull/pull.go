package pull

import (
	aws_tool "backup-tool/cmd/app/utils/aws"
	file_tool "backup-tool/cmd/app/utils/file"
	"backup-tool/internal/pkg/lib/app"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type App struct{}

var (
	source      string
	destination string
	awsBucket   string
	awsRegion   string
	unzipFile   bool
)

func New() app.AppService {
	return &App{}
}

func (a *App) Init(ctx *cobra.Command) error {
	ctx.Flags().StringVarP(&awsRegion, "region", "", "", "set s3 bucket region")
	ctx.Flags().StringVarP(&awsBucket, "bucket", "", "", "set s3 bucket name")
	ctx.Flags().BoolVarP(&unzipFile, "unzip", "", false, "unzip object file")
	return nil
}

func (a *App) Start(ctx *cobra.Command, args []string) error {
	source := args[0]
	destination := args[1]

	s3Svr, err := aws_tool.New(awsRegion)
	if err != nil {
		return err
	}

	file, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("Unable to open file %q, %v", destination, err)
	}

	defer file.Close()

	numBytes, err := s3Svr.Download(awsBucket, source, file)
	if err != nil {
		return fmt.Errorf("Unable to download item %q, %v", destination, err)
	}

	if unzipFile {
		err = file_tool.Unzip(
			destination, file_tool.GetFolderName(destination, false),
		)
	}

	if err != nil {
		return err
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	return nil
}

func (a *App) Close(ctx *cobra.Command) error {
	return nil
}
