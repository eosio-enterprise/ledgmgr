package main

import (
	"os"

	"github.com/eosio-enterprise/ledgmgr/cmd"
	"github.com/spf13/cobra"
)

func main() {

	cmdWatch := cmd.Watch()

	var rootCmd = &cobra.Command{
		Use: "ledgmgr",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(cmdWatch)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
