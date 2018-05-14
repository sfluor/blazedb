package cmd

import (
	"github.com/sfluor/blazedb/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Args:  cobra.ExactArgs(1),
	Short: "Connect to a blazedb server",
	Run: func(cmd *cobra.Command, args []string) {
		client := client.New(args[0])
		client.Start()
	},
}
