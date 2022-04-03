package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/slides"
)

var rootCmd = &cobra.Command{
	Use:   "k8s-workshop",
	Short: "A CLI tool for the x-cellent k8s workshop",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	defer func() {
		_ = viper.BindPFlags(rootCmd.PersistentFlags())
	}()

	rootCmd.AddCommand(slides.Cmd)
}

func Execute() error {
	return rootCmd.Execute()
}
