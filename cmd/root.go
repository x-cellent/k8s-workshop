package cmd

import (
	"context"
	"embed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/cluster"
	"github.com/x-cellent/k8s-workshop/cmd/exercise"
	"github.com/x-cellent/k8s-workshop/cmd/install"
	"github.com/x-cellent/k8s-workshop/cmd/slides"
)

var rootCmd = &cobra.Command{
	Use:   "w6p",
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
	rootCmd.AddCommand(cluster.Cmd)
	rootCmd.AddCommand(exercise.Cmd)
}

func Execute(docs, exercises embed.FS) error {
	ctx := context.WithValue(context.Background(), "docs", docs)
	ctx = context.WithValue(ctx, "exercises", exercises)
	return rootCmd.ExecuteContext(ctx)
}
