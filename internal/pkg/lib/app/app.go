package app

import (
	"github.com/spf13/cobra"
)

type AppService interface {
	Init(*cobra.Command) error
	Start(*cobra.Command, []string) error
	Close(*cobra.Command) error
}
