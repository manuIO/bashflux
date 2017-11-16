package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/mainflux/bashflux/cmd"
	"github.com/spf13/cobra"
)

// Config struct
type Config struct {
	// HTTP
	HTTPHost string
	HTTPPort int
}

func main() {

	var s string
	var limit int
	var conf Config

	conf.HTTPHost = "0.0.0.0"
	conf.HTTPPort = 8180

	// print bashflux banner
	color.Yellow(banner)

	////
	// Version
	////
	var cmdVersion= &cobra.Command{
		Use:   "version",
		Short: "Get manager version",
		Long:  `Mainflux server health checkt.`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			s = cmd.Version()
		},
	}

	////
	// Clients
	////
	var cmdClients = &cobra.Command{
		Use:   "clients",
		Short: "clients <options>",
		Long:  `Clients handling: create, delete or update clients.`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 1 {
				s = cmd.GetClients(args[0])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	// Create Client
	var cmdCreateClient = &cobra.Command{
		Use:   "create",
		Short: "create <JSON_client> <user_token>",
		Long:  `Create new client, generate his UUID and store it`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 2 {
				msg := args[0]
				token := args[1]
				s = cmd.CreateClient(msg, token)
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	// Get Client
	var cmdGetClient = &cobra.Command{
		Use:   "get",
		Short: "get <user_token> or get <client_id> <user_token>",
		Long:  `Get all clients or client by id`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 1 {
				s = cmd.GetClients(args[0])
			} else if len(args) == 2 {
				s = s + cmd.GetClient(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	// Delete Client
	var cmdDeleteClient = &cobra.Command{
		Use:   "delete",
		Short: "delete all or delete <client_id> <user_token>",
		Long:  `Removes client from database`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 2 {
				if args[0] == "all" {
					s = cmd.DeleteAllClients(args[1])
				} else {
					s = s + cmd.DeleteClient(args[0], args[1])
				}
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	// Update Client
	var cmdUpdateClient = &cobra.Command{
		Use:   "update",
		Short: "update <client_id> <JSON_string> <user_token>",
		Long:  `Update client record`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 2 {
				s = cmd.UpdateClient(args[0], args[1], args[2])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	////
	// Channels
	////
	var cmdChannels = &cobra.Command{
		Use:   "channels",
		Short: "Manipulation with channels",
		Long:  `Manipulation with channels: create, delete or update channels`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 1 {
				s = cmd.GetChannels(limit, args[0])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	// Create Channel
	var cmdCreateChannel = &cobra.Command{
		Use:   "create",
		Short: "create <JSON_channel> <user_token>",
		Long:  `Creates new channel and generates it's UUID`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 2 {
				msg := args[0]
				token := args[1]
				s = cmd.CreateChannel(msg, token)
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	// Get Channel
	var cmdGetChannel = &cobra.Command{
		Use:   "get",
		Short: "get <user_token> or get <channel_id> <user_token>",
		Long:  `Gets list of all channels or gets channel by id`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 1 {
				s = cmd.GetChannels(limit, args[0])
			} else if len(args) == 2 {
				s = s + cmd.GetChannel(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	// Update Channel
	var cmdUpdateChannel = &cobra.Command{
		Use:   "update",
		Short: "update <channel_id> <JSON_string> <user_token>",
		Long:  `Updates channel record`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 3 {
				s = cmd.UpdateChannel(args[0], args[1], args[2])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	// Delete Channel
	var cmdDeleteChannel = &cobra.Command{
		Use:   "delete",
		Short: "delete <channel_id> <user_token>",
		Long:  `Delete channel by ID`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 2 {
				if args[0] == "all" {
					s = cmd.DeleteAllChannels(args[1])
				} else {
					s = s + cmd.DeleteChannel(args[0], args[1])
				}
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	////
	// Messages
	////
	var cmdMessages = &cobra.Command{
		Use:   "msg",
		Short: "Send or retrieve messages",
		Long:  `Send or retrieve messages: control message flow on the channel`,
	}

	// Send Message
	var cmdSendMessage = &cobra.Command{
		Use:   "send",
		Short: "send <channel_id> <JSON_string> <user_token>",
		Long:  `Sends message on the channel`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			// TODO: implement nginx and remove this
			cmd.SetServerAddr(conf.HTTPHost, 7070)
			if len(args) == 3 {
				s = cmd.SendMsg(args[0], args[1], args[2])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	////
	// Users
	////
	var cmdUsers = &cobra.Command{
		Use:   "users",
		Short: "User management",
		Long:  `Manages users in the system (creation, deletition and other system admin)`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			s = "Usage: " + cmdCobra.Short + ". Needs additional commands (see --help)"
		},
	}

	// Create User
	var cmdCreateUser = &cobra.Command{
		Use:   "create",
		Short: "create <username> <password>",
		Long:  `Creates new user`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 2 {
				s = cmd.CreateUser(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
	}

	////
	// Sessions
	////
	var cmdSession = &cobra.Command{
		Use:   "tokens",
		Short: "Tokens creation",
		Long:  `Used for tokens manipulation`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			s = "Usage: " + cmdCobra.Short + ". Need additional commands (see --help)"
		},
	}

	// Init Session
	var cmdCreateToken = &cobra.Command{
		Use:   "create",
		Short: "create <username> <password>",
		Long:  `Creates new token`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			if len(args) == 2 {
				s = cmd.CreateToken(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
		},
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
	rootCmd.AddCommand(cmdVersion)
	rootCmd.AddCommand(cmdClients)
	rootCmd.AddCommand(cmdChannels)
	rootCmd.AddCommand(cmdMessages)
	rootCmd.AddCommand(cmdSession)
	rootCmd.AddCommand(cmdUsers)

	// Clients
	cmdClients.AddCommand(cmdCreateClient)
	cmdClients.AddCommand(cmdGetClient)
	cmdClients.AddCommand(cmdUpdateClient)
	cmdClients.AddCommand(cmdDeleteClient)

	// Channels
	cmdChannels.AddCommand(cmdCreateChannel)
	cmdChannels.AddCommand(cmdGetChannel)
	cmdChannels.AddCommand(cmdUpdateChannel)
	cmdChannels.AddCommand(cmdDeleteChannel)

	// Messages
	cmdMessages.AddCommand(cmdSendMessage)

	// Users
	cmdUsers.AddCommand(cmdCreateUser)

	// Token
	cmdSession.AddCommand(cmdCreateToken)

	// Root Flags
	rootCmd.PersistentFlags().StringVarP(
		&conf.HTTPHost, "host", "m", conf.HTTPHost, "HTTP Host address")
	rootCmd.PersistentFlags().IntVarP(
		&conf.HTTPPort, "port", "p", conf.HTTPPort, "HTTP Host Port")

	// Channels Flags
	cmdChannels.PersistentFlags().IntVarP(
		&limit, "limit", "l", 0, "limit query parameter")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	// Print response
	fmt.Println(s + "\n\n")
}

var banner =
` __                 __     ___ __
|  |--.---.-.-----.|  |--.'  _|  |.--.--.--.--.
|  _  |  _  |__ --||     |   _|  ||  |  |_   _|
|_____|___._|_____||__|__|__| |__||_____|__.__|

             == Mainflux CLI ==

`
