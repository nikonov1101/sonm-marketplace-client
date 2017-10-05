package commands

import (
	"github.com/spf13/cobra"
	"io/ioutil"
)

func searchBySlot(cmd *cobra.Command, slot []byte) error {
	// todo:  parse file to pb.Slot
	// todo: implement grpc call to market
	cmd.Printf("FILE: \r\n%s", string(slot))
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
