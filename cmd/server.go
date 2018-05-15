package cmd

import (
	"github.com/sfluor/blazedb/server"
	"github.com/spf13/cobra"
)

var configPath *string

func init() {
	rootCmd.AddCommand(srvCmd)
	// Flags
	configPath = srvCmd.Flags().StringP("config", "c", "/etc/blazedb.toml", "Config path")
}

var srvCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a blazedb server",
	Run: func(cmd *cobra.Command, args []string) {
		config := server.LoadConfig(*configPath)
		srv := server.New(config.Port)
		srv.Start()
	},
}
