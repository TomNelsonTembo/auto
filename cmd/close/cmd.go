package close

import (
	"github.com/spf13/cobra"
)

// closeCmd represents the close command
func NewCloseCmd() *cobra.Command {
	closeCmd := &cobra.Command{
		Use:   "Auto Close",
		Short: "Close is a palette that contains cloudtrail commands",
		Long:  `Close is a palette that contains cloudtrail commands`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	closeCmd.AddCommand()

	return closeCmd
}
