package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Watch ...
func Watch() *cobra.Command {
	var command = &cobra.Command{
		Use:          "watch",
		Short:        "Watch a ledger file for changes and run a report on change",
		Example:      ` ledgmgr watch --file myledger.dat --report balancesheet`,
		SilenceUsage: false,
	}

	command.Flags().StringP("file", "f", "", "Ledger (see hledger.org) file to watch for updates")
	command.MarkFlagRequired("file")
	command.Flags().StringP("report", "r", "balancesheet", "Ledger report to run on file change")

	command.RunE = func(command *cobra.Command, args []string) error {

		ledgerFile, _ := command.Flags().GetString("file")
		// if len(ledgerFile) == 0 {
		// 	return fmt.Errorf("--cid required")
		// }
		report, _ := command.Flags().GetString("report")
		counter := 0

		priorUpdatedTime := time.Now()
		color.Cyan("Listening for updates to " + ledgerFile + "...")

		for {
			fileInfo, err := os.Stat(ledgerFile)
			if err != nil {
				log.Fatal(err)
			}
			lastUpdatedTime := fileInfo.ModTime()
			if lastUpdatedTime.After(priorUpdatedTime) {
				priorUpdatedTime = lastUpdatedTime

				ledgerCmd := exec.Command("hledger", "--file="+ledgerFile, report)

				stdout, err := ledgerCmd.StdoutPipe()
				if err != nil {
					log.Fatal(err)
				}

				if err = ledgerCmd.Start(); err != nil {
					log.Fatal(err)
				}

				reportOutput, _ := ioutil.ReadAll(stdout)
				if err := ledgerCmd.Wait(); err != nil {
					log.Fatal(err)
				}

				switch counter % 3 {
				case 0:
					color.Cyan(string(reportOutput))
				case 1:
					color.Red(string(reportOutput))
				case 2:
					color.Green(string(reportOutput))
				}
				counter++
			}
		}
	}

	return command
}
