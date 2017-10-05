package commands

import (
	"errors"
	"os"

	pb "github.com/sonm-io/core/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	errOrderFileRequired = errors.New("JSON file with Order spec is required")
	errSlotFileRequired  = errors.New("JSON file with Slot spec is required")
	errOrderIDRequired   = errors.New("Order ID is required")

	marketPlaceEndpt string
)

func Root() *cobra.Command {
	rootCmd := &cobra.Command{Use: "market-client"}
	rootCmd.AddCommand(
		templateRootCmd,
		createOrderCmd,
		searchOrderCmd,
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
	conn, err := grpc.Dial(marketPlaceEndpt, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewMarketClient(conn)
	return client, nil
}
