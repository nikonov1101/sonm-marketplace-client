package commands

import (
	"errors"
	"os"

	pb "github.com/sonm-io/core/proto"
	"github.com/sonm-io/core/util"
	"github.com/spf13/cobra"
)

var (
	errOrderFileRequired = errors.New("JSON file with Order spec is required")
	errSlotFileRequired  = errors.New("JSON file with Slot spec is required")
	errOrderIDRequired   = errors.New("Order ID is required")

	marketPlaceEndpt string
	searchLimit      uint64
	orderType        string
)

func Root() *cobra.Command {
	rootCmd := &cobra.Command{Use: "market-client"}

	searchOrderCmd.PersistentFlags().Uint64Var(&searchLimit, "limit", 15, "Order limit")
	searchOrderCmd.PersistentFlags().StringVar(&orderType, "type", "ANY",
		"Order type to search ANY, BID or ASK")

	rootCmd.AddCommand(
		templateRootCmd,
		searchOrderCmd,
		createOrderCmd,
		deleteOrderCmd,
	)

	rootCmd.PersistentFlags().StringVar(
		&marketPlaceEndpt,
		"addr",
		"127.0.0.1:9095",
		"Marketplace addr",
	)

	rootCmd.SetOutput(os.Stdout)
	return rootCmd
}

func initGrpcClient() (pb.MarketClient, error) {
	conn, err := util.GetClientConn(marketPlaceEndpt)
	if err != nil {
		return nil, err
	}

	client := pb.NewMarketClient(conn)
	return client, nil
}
