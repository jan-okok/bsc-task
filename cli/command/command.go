package command

import (
	"github.com/spf13/cobra"
	"github/bsc-task/action"
	"os"
)

var bscTaskCmd = &cobra.Command{
	Use:   "bscTaskCLI",
	Short: "BSC Task Command Line Client.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		}
	},
}

var (
	txCmd = &cobra.Command{
		Use:   "sendTransactions",
		Short: "batch send transactions",
		Run:   SendTransaction,
	}
)

func Execute() {
	bscTaskCmd.AddCommand(txCmd)
	if _, err := bscTaskCmd.ExecuteC(); err != nil {
		os.Exit(1)
	}
}

func SendTransaction(cmd *cobra.Command, args []string) {
	action.SendTransaction()
}
