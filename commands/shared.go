package commands

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "control",
	Short: "Утилита для управления",
	Long:  "Утилита для управления",
}

var logger *log.Logger

// Execute executes
func Execute(version string) {
	logger = log.Default()
	logger.SetPrefix("")
	logger.SetOutput(os.Stdout)
	if version == "development" {
		logger.SetFlags(log.Llongfile | log.Ltime)
	} else {
		logger.SetFlags(log.Lmsgprefix)
	}
	logger.Printf("Control v: %s", version)
	rootCmd.Version = version
	rootCmd.AddCommand(startAgentCommand)
	rootCmd.AddCommand(startServerCommand)
	rootCmd.Execute()
}
