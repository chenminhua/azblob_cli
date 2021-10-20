package blob

import (
	"github.com/spf13/cobra"
)

// byeCmd represents the bye command
func NewBlobCmd() *cobra.Command {
	blobCmd := &cobra.Command{
		Use:   "blob",
		Short: "upload, download blob file",
		Long:  `upload, download blob file`,
	}
	blobCmd.AddCommand(NewBlobUploadCmd())
	return blobCmd
}
