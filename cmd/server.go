package cmd

import (
	"github.com/sfluor/blazedb/server"
	"github.com/spf13/cobra"
)

var srvPort *uint

func init() {
	rootCmd.AddCommand(srvCmd)
	// Flags
	srvPort = srvCmd.Flags().UintP("port", "p", 9876, "Port that blazedb server need to listen on")
}

var srvCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a blazedb server",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New(*srvPort)
		srv.Start()
	},
}
