package commands

import (
	"github.com/spf13/cobra"
)

func deleteOrder(cmd *cobra.Command, id string) error {
	// todo: implement grpc call to market
	cmd.Printf("DELETE ORDER: %s\r\n", id)
	return nil
}

var deleteOrderCmd = &cobra.Command{
	Use:   "delete <order_id>",
	Short: "delete orderTemplate from Marketplace",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errOrderIDRequired
		}

		id := args[0]
		err := deleteOrder(cmd, id)
		return err
	},
}
