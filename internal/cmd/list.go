package cmd

import (
	"file_management/internal/core/file"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var detailed bool
	// Initialize logger
	//log := logger.New()
	cmd := &cobra.Command{
		Use:   "ls [path]",
		Short: "List files and directories",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}

			manager := file.New(log)
			return manager.List(path, detailed)
		},
	}

	cmd.Flags().BoolVarP(&detailed, "detailed", "l", false, "Show detailed information")
	return cmd
}
