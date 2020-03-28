package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func main() {

	counter := 0
	var returnError error

	for {
		color.Cyan("Listening for updated transactions...")
		ledgerCmd := exec.Command("hledger", "--file="+"/Users/max/dev/eosio-enterprise/chappe/ledger.dat", os.Args[1]) //"bs")

		stdout, err := ledgerCmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		if returnError = ledgerCmd.Start(); err != nil {
			log.Fatal(err)
		}

		reportOutput, _ := ioutil.ReadAll(stdout)

		if err := ledgerCmd.Wait(); err != nil {
			log.Fatal(err)
		}

		sleepCmd := exec.Command("sleep", "3")
		err = sleepCmd.Run()

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

	log.Printf("LedgerMgr finished with return code: %v", returnError)
}
