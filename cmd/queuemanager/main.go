package main

import (
	"queuemanager/config"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Application entry point
func main() {
	queuemanagerCmd().Execute()
}

// New constructs a new CLI interface for execution
func queuemanagerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "queuemanager",
		Short: "Run the gRPC service for the Article Token Service",
		Run:   queuemanagerRun,
	}
	// Global flags
	pflags := cmd.PersistentFlags()
	pflags.StringP("config-file", "c", "", "path to configuration file")
	pflags.String("log-format", "", "log format [console|json]")
	// Local Flags
	flags := cmd.Flags()
	flags.StringP("listen", "l", "", "server listen address")
	// Bind flags to config options
	config.BindPFlags(map[string]*pflag.Flag{
		config.CONFIG_PATH_KEY: pflags.Lookup("config-file"),
		config.LOG_FORMAT_KEY:  pflags.Lookup("log-format"),
	})
	// Add sub commands
	cmd.AddCommand(versionCmd())
	return cmd
}

// queuemanagerRun is executed when the CLI executes
// the queuemanager command
func queuemanagerRun(cmd *cobra.Command, _ []string) {
	cmd.Help()
}
