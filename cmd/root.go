package cmd

import (
	"context"
	"embed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/install"
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
	rootCmd.AddCommand(install.Cmd)
}

func Execute(docs embed.FS) error {
	ctx := context.WithValue(context.Background(), "docs", docs)
	return rootCmd.ExecuteContext(ctx)
}
