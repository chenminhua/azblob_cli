package container

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)


type ContainerListCmdOption struct {
	accountName   string
	accountKey    string
}

func newContainerListCmdOption(cmd *cobra.Command) *ContainerListCmdOption {
	accountName := ""
	accountKey := ""
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
	return &ContainerListCmdOption{
		accountName: accountName,
		accountKey: accountKey,
	}
}

func (op *ContainerListCmdOption) Complete() {

}

func NewContainerListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list",
		Long:  `list`,
		Run: func(cmd *cobra.Command, args []string) {
			o := newContainerListCmdOption(cmd)
			if o != nil {
				o.Complete()
			}
		},
	}
	return cmd
}
