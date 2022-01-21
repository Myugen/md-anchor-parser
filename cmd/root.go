package cmd

import (
	"github.com/spf13/cobra"
)

type Reader interface {
	Read()
}

func NewRootCmd(reader Reader) *cobra.Command {
	return &cobra.Command{
		Use:   "md-anchor-parse",
		Short: "ASDF",
		Long:  "ASDFG",
		Run:   rootExecution(reader),
	}
}

func rootExecution(reader Reader) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		reader.Read()
	}
}
