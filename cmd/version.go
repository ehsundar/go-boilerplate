package cmd

import (
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var Version = "0.0.1"

func RegisterVersionCommand(root *cobra.Command) {
	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of boilerplate",
		Long:  "Display the current version of the boilerplate application.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cmd.Printf("boilerplate version %s\n", Version)

			return nil
		},
	})
}
