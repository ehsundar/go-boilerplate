package cmd

import (
	"github.com/spf13/cobra"
)

func RegisterVersionCommand(root *cobra.Command) {
	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number and build information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cmd.Printf("%s version %s\n", AppName, Version)
			cmd.Printf("commit: %s\n", Commit)
			cmd.Printf("date: %s\n", Date)

			return nil
		},
	})
}
