package cmd

import (
	"fmt"
	"os"

	"github.com/rajatjindal/ghcr-cleanup/pkg/github"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	token       string
	username    string
	packageName string
	dryrun      bool
	debug       bool
	minRetain   int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ghcr-cleanup",
	Short: "ghcr cleanup is a cli tool to cleanup old images from ghcr container registry",
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		client := github.NewClient(token, true)
		err := client.CleanupPackages(username, packageName, minRetain)
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "github token, defaults to env variable GITHUB_TOKEN")
	rootCmd.Flags().StringVar(&username, "username", "", "rajatjindal")
	rootCmd.MarkFlagRequired("username")

	rootCmd.Flags().StringVar(&packageName, "package-name", "", "translatethread")
	rootCmd.MarkFlagRequired("package-name")

	rootCmd.Flags().BoolVar(&debug, "debug", false, "enable debug logging")
	rootCmd.Flags().BoolVar(&dryrun, "dryrun", true, "run in dry run mode")
	rootCmd.Flags().IntVar(&minRetain, "min-retain", 10, "retain min versions")
}
