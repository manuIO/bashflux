package main

import (
	"log"

	"github.com/mainflux/bashflux/cmd"
	"github.com/spf13/cobra"
)

type Config struct {
	HTTPHost string
	HTTPPort int
}

func main() {

	conf := &Config{
		HTTPHost: "localhost",
		HTTPPort: 0,
	}

	// Root
	var rootCmd = &cobra.Command{
		Use: "bashflux",
		PersistentPreRun: func(cmdCobra *cobra.Command, args []string) {
			// Set HTTP server address
			cmd.SetServerAddr(conf.HTTPHost, conf.HTTPPort)
		},
	}

	// Root Commands
	rootCmd.AddCommand(cmd.CmdVersion)
	rootCmd.AddCommand(cmd.CmdClients)
	rootCmd.AddCommand(cmd.CmdChannels)
	rootCmd.AddCommand(cmd.CmdMessages)
	rootCmd.AddCommand(cmd.CmdSession)
	rootCmd.AddCommand(cmd.CmdUsers)

	// Clients
	cmd.CmdClients.AddCommand(cmd.CmdCreateClient)
	cmd.CmdClients.AddCommand(cmd.CmdGetClient)
	cmd.CmdClients.AddCommand(cmd.CmdUpdateClient)
	cmd.CmdClients.AddCommand(cmd.CmdDeleteClient)
	cmd.CmdClients.AddCommand(cmd.CmdConnectClient)
	cmd.CmdClients.AddCommand(cmd.CmdDisconnectClient)

	// Channels
	cmd.CmdChannels.AddCommand(cmd.CmdCreateChannel)
	cmd.CmdChannels.AddCommand(cmd.CmdGetChannel)
	cmd.CmdChannels.AddCommand(cmd.CmdUpdateChannel)
	cmd.CmdChannels.AddCommand(cmd.CmdDeleteChannel)

	// Messages
	cmd.CmdMessages.AddCommand(cmd.CmdSendMessage)

	// Users
	cmd.CmdUsers.AddCommand(cmd.CmdCreateUser)

	// Token
	cmd.CmdSession.AddCommand(cmd.CmdCreateToken)

	// Root Flags
	rootCmd.PersistentFlags().StringVarP(
		&conf.HTTPHost, "host", "m", conf.HTTPHost, "HTTP Host address")
	rootCmd.PersistentFlags().IntVarP(
		&conf.HTTPPort, "port", "p", conf.HTTPPort, "HTTP Host Port")

	// Client and Channels Flags
	rootCmd.PersistentFlags().IntVarP(
		&cmd.Limit, "limit", "l", 100, "limit query parameter")
	rootCmd.PersistentFlags().IntVarP(
		&cmd.Offset, "offset", "o", 0, "offset query parameter")

	// Set TLS certificates
	cmd.SetCerts()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
