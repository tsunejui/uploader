package cmd

import (
	"backup-tool/internal/pkg/lib/app"
	"fmt"

	"github.com/spf13/cobra"
)

type Command struct {
	Use     string
	Short   string
	Long    string
	Service app.AppService
	Args    cobra.PositionalArgs
}

type CommandService struct {
	rootCmd *cobra.Command
}

func New(cmd *Command) (*CommandService, error) {
	command, err := newCommand(cmd)
	if err != nil {
		return nil, err
	}
	return &CommandService{
		rootCmd: command,
	}, nil
}

func (s *CommandService) AddCommand(cmds ...*Command) error {
	for _, cmd := range cmds {
		command, err := newCommand(cmd)
		if err != nil {
			return err
		}
		s.rootCmd.AddCommand(command)
	}
	return nil
}

func (s *CommandService) Execute() error {
	return s.rootCmd.Execute()
}

func newCommand(cmd *Command) (*cobra.Command, error) {
	command := &cobra.Command{
		Use:   cmd.Use,
		Short: cmd.Short,
		Long:  cmd.Long,
	}

	if cmd.Args != nil {
		command.Args = cmd.Args
	}

	if cmd.Service != nil {
		if err := cmd.Service.Init(command); err != nil {
			return nil, fmt.Errorf("Failed to initial app: %v", err)
		}

		command.RunE = startApp(cmd.Service)
	}

	return command, nil
}

func startApp(application app.AppService) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		defer application.Close(cmd)

		if err := application.Start(cmd, args); err != nil {
			return fmt.Errorf("Failed to start app: %v", err)
		}

		return nil
	}
}
