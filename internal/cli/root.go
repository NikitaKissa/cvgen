package cli

import (
	"github.com/spf13/cobra"
)

func Execute() error {
	root := &cobra.Command{
		Use:   "cvgen",
		Short: "cvgen — CLI CV/resume generator",
	}

	root.AddCommand(newGenerateCmd())
	root.AddCommand(newVerifyCmd())
	root.AddCommand(newVersionCmd())

	return root.Execute()
}
