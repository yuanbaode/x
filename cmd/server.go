package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuanbaode/x/service/server"
	"log"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start a service registry center",
	Example: "",

	Run: serve,
}

var (
	serverHost  string
	serverStore string
)

func init() {
	serverCmd.Flags().StringVarP(&serverHost, "host", "", "127.0.0.1:7777", "specify the host of the goc server")
	serverCmd.Flags().StringVarP(&serverStore, "store", "", ".goc.kvstore", "specify the host of the goc server")

	rootCmd.AddCommand(serverCmd)
}

func serve(cmd *cobra.Command, args []string) {
	cfg := server.Config{

	}
	err := server.NewServer(cfg).Start()
	if err != nil {
		log.Fatal(err)
	}
}
