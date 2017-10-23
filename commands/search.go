package commands

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
	pb "github.com/sonm-io/core/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func parseOrderType(ty string) (pb.OrderType, error) {
	switch ty {
	case "ANY":
		return pb.OrderType_ANY, nil
	case "BID":
		return pb.OrderType_BID, nil
	case "ASK":
		return pb.OrderType_ASK, nil
	default:
		return pb.OrderType_ANY, errors.New("Unknown order type")
	}
}

func searchBySlot(cmd *cobra.Command, data []byte) error {
	slot := &pb.Slot{}
	err := json.Unmarshal(data, &slot)
	if err != nil {
		return err
	}

	cc, err := initGrpcClient()
	if err != nil {
		return err
	}

	ordType, err := parseOrderType(orderType)
	if err != nil {
		return err
	}

	req := &pb.GetOrdersRequest{
		Slot:      slot,
		OrderType: ordType,
		Count:     searchLimit,
	}

	result, err := cc.GetOrders(context.Background(), req)
	if err != nil {
		return err
	}

	ordersCount := len(result.Orders)

	if ordersCount == 0 {
		cmd.Println("No orders found")
		return nil
	}

	cmd.Printf("Found %d orders:\r\n", len(result.Orders))
	for i, order := range result.Orders {
		cmd.Printf("%d) %s\r\n", i, order.Id)
	}

	return nil
}

var searchOrderCmd = &cobra.Command{
	Use:   "search <slotTemplate.json>",
	Short: "search for orders by Slot spec",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errSlotFileRequired
		}

		slotFile := args[0]
		f, err := ioutil.ReadFile(slotFile)
		if err != nil {
			return err
		}

		err = searchBySlot(cmd, f)
		return err
	},
}
