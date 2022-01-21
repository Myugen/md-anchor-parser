package cmd

import (
	"github.com/myugen/md-anchor-parser/reader"
	"github.com/spf13/cobra"
)

func NewRootCmd(r ...reader.Reader) *cobra.Command {
	return &cobra.Command{
		Use:   "md-anchor-parse",
		Short: "ASDF",
		Long:  "ASDFG",
		Run:   rootExecution(r...),
	}
}

func rootExecution(r ...reader.Reader) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var defaultReader reader.Reader

		if len(r) > 0 {
			defaultReader = r[0]
		} else {
			defaultReader = new(reader.MarkdownReader)
		}
		defaultReader.Read()
	}
}
