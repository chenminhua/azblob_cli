package blob

import (
	"fmt"
	"github.com/chenminhua/azblob_cli/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

type UploadCmdOptions struct {
	accountName   string
	accountKey    string
	containerName string
	fileName      string
	filePath      string
}

func newUploadCmdOptions(cmd *cobra.Command) *UploadCmdOptions {
	accountName := ""
	accountKey := ""
	containerName := ""
	if an, ok := viper.Get("accountName").(string); ok && an != "" {
		accountName = an
	} else {
		log.Fatal("No accountName")
		return nil
	}
	if an, ok := viper.Get("accountKey").(string); ok && an != "" {
		accountKey = an
	} else {
		log.Fatal("No accountKey")
		return nil
	}
	containerName, _ = cmd.Flags().GetString("containerName")
	if containerName == "" {
		// fallback to config
		if an, ok := viper.Get("containerName").(string); ok && an != "" {
			containerName = an
		} else {
			log.Fatal("No containerName")
			return nil
		}
	}
	fileName, _ := cmd.Flags().GetString("fileName")
	filePath, _ := cmd.Flags().GetString("filePath")
	if fileName == "" {
		log.Fatal("No fileName")
		return nil
	}
	if filePath == "" {
		log.Fatal("No filePath")
		return nil
	}
	fmt.Println(accountName, accountKey, containerName, fileName, filePath)
	return &UploadCmdOptions{
		accountName:   accountName,
		accountKey:    accountKey,
		containerName: containerName,
		fileName:      fileName,
		filePath:      filePath,
	} 
}

func NewBlobUploadCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "upload",
		Short: "upload blob file",
		Long:  `upload blob file`,
		Run: func(cmd *cobra.Command, args []string) {
			o := newUploadCmdOptions(cmd)
			if o != nil {
				o.Complete()
			}
		},
	}
	cmd.Flags().String("containerName", "", "container name")
	cmd.Flags().String("filePath", "", "file path")
	cmd.Flags().String("fileName", "", "file name")
	return cmd
}

func (upo *UploadCmdOptions) Complete() {
	blobCli, _ := pkg.NewBlobClient(upo.accountName, upo.accountKey, upo.containerName)
	blobCli.Upload(upo.filePath, upo.fileName)
}
