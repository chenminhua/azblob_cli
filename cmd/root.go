package cmd
import (
	"fmt"
	"github.com/chenminhua/azblob_cli/cmd/blob"
	"github.com/chenminhua/azblob_cli/cmd/container"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string

	rootCmd = &cobra.Command{
		Use:   "azblob",
		Short: "A azure blob cmdline tool",
		Long: `A azure blob cmdline tool`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.azblob_conf.yaml)")

	rootCmd.AddCommand(blob.NewBlobCmd())
	rootCmd.AddCommand(container.NewContainerCmd())
}

func initConfig() {
	fmt.Println("hello init config")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".azblob_conf.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
}