package cmd

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/indexer"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download files from 0g storage",
	Run: func(cmd *cobra.Command, args []string) {
		w3client := blockchain.MustNewWeb3(os.Getenv("EVM_RPC"), os.Getenv("PRIVATE_KEY"))
		defer w3client.Close()

		indexerClient, err := indexer.NewClient(os.Getenv("INDEXER_RPC"))
		if err != nil {
			log.Fatalf("create indexer client error: %v", err)
		}

		ctx := context.Background()

		roots := os.Getenv("ROOTS")

		if err := indexerClient.DownloadFragments(ctx, strings.Split(roots, ","), "downloaded_file.bin", true); err != nil {
			log.Fatalf("Download file error: %v", err)
		}
		log.Printf("Download successful!\n")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}