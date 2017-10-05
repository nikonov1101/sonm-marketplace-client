package commands

import "github.com/spf13/cobra"

func init() {
	templateRootCmd.AddCommand(templateOrderCmd, templateSlotCmd)
}

var templateRootCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate template",
}

var templateSlotCmd = &cobra.Command{
	Use:   "slotTemplate",
	Short: "Generate slotTemplate template",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Printf("%s\r\n", getSlotTemplate())
	},
}

var templateOrderCmd = &cobra.Command{
	Use: "orderTemplate",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Printf("%s\r\n", getOrderTemplate())
	},
}
