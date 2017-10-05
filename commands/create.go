package commands

import (
	"encoding/json"
	"io/ioutil"

	pb "github.com/sonm-io/core/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func sendOrderToMarket(cmd *cobra.Command, data []byte) error {
	order := &pb.Order{}
	err := json.Unmarshal(data, &order)
	if err != nil {
		return err
	}

	cc, err := initGrpcClient()
	if err != nil {
		return err
	}

	order, err = cc.CreateOrder(context.Background(), order)
	if err != nil {
		return err
	}

	cmd.Printf("Order created. ID = %s\r\n", order.Id)
	return nil
}

var createOrderCmd = &cobra.Command{
	Use:   "create <orderTemplate.json>",
	Short: "create orderTemplate on Marketplace",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errOrderFileRequired
		}

		orderFile := args[0]
		f, err := ioutil.ReadFile(orderFile)
		if err != nil {
			return err
		}

		err = sendOrderToMarket(cmd, f)
		return err
	},
}
