package container

import "github.com/spf13/cobra"

func NewContainerCmd() *cobra.Command {
	blobCmd := &cobra.Command{
		Use:   "container",
		Short: "container",
		Long:  `container`,
	}
	blobCmd.AddCommand(NewContainerListCmd())
	return blobCmd
}