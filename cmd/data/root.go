package data

import (
	"github.com/spf13/cobra"
)

type dataConfig struct {
	dataDir string
}

var cfg = dataConfig{}

func RootCmd() *cobra.Command {
	index := &cobra.Command{
		Use:   "data",
		Short: "commands related to to the search index",
	}

	index.PersistentFlags().StringVarP(&cfg.dataDir, "data-dir", "d", "./var/data/episodes", "Path to the raw data files")

	index.AddCommand(InitCmd())
	index.AddCommand(ImportTranscriptsCmd())
	index.AddCommand(ImportSpotifyData())

	return index
}
