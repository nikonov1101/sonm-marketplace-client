package commands

import (
	pb "github.com/sonm-io/core/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func deleteOrder(cmd *cobra.Command, id string) error {
	cc, err := initGrpcClient()
	if err != nil {
		return err
	}

	_, err = cc.CancelOrder(context.Background(), &pb.Order{Id: id})
	if err != nil {
		return err
	}

	cmd.Println("Order deleted.")
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
