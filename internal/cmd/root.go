package cmd

import (
	"file_management/internal/config"
	"file_management/internal/utils/logger"
	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
	log *logger.Logger
)

func NewRootCmd(config *config.Config, logger *logger.Logger) *cobra.Command {
	cfg = config
	log = logger

	rootCmd := &cobra.Command{
		Use:   "hinq",
		Short: "A modern file manager CLI tool",
		Long:  `A file manager CLI tool with features for managing files and directories.`,
	}

	// Add all subcommands
	rootCmd.AddCommand(
		newListCmd(),
		//newCopyCmd(),
		//newMoveCmd(),
		//newDeleteCmd(),
		//newMkdirCmd(),
		//newSearchCmd(),
	)

	return rootCmd
}

func Execute(config *config.Config, logger *logger.Logger) error {
	return NewRootCmd(config, logger).Execute()
}
