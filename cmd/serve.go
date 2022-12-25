package cmd

import (
	"fmt"

	"github.com/ash2shukla/go-http-server/config"
	"github.com/ash2shukla/go-http-server/lib/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve content from current directory.",
	Long:  "Serve content from current directory.",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.GetConfig()

		host := config.Serve.Host
		port := fmt.Sprint(config.Serve.Port)

		server.ServerLoop(host, port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntP("port", "p", 8989, "Set port for serving.")
	serveCmd.Flags().StringP("host", "b", "localhost", "Set host for serving.")
	serveCmd.Flags().String("statichome", ".", "Set home dir for static content.")

	viper.BindPFlag("serve.port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("serve.host", serveCmd.Flags().Lookup("host"))
	viper.BindPFlag("serve.staticdir", serveCmd.Flags().Lookup("statichome"))
}
