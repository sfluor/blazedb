package cmd

import (
	"github.com/spf13/cobra"
)

var addr string

func init() {
	rootCmd.AddCommand(clientCmd)
	// Flags
	addr = *srvCmd.Flags().StringP("addr", "a", "localhost:9876", "Address of the blazedb server")
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Connect to a blazedb server",
	Run: func(cmd *cobra.Command, args []string) {
		client := client.New(addr)
		client.Start()
	},
}
