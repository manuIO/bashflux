/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"fmt"
	"log"

	"github.com/mainflux/mainflux-cli/cmd"
	"github.com/spf13/cobra"
)

func main() {

	var httpHost = "0.0.0.0"
	var httpPort = 7070

	var s string

	var startTime string
	var endTime string

	// Set HTTP server address
	cmd.SetServerAddr(httpHost, httpPort)

	// print mainflux-cli banner
	fmt.Println(banner)

	////
	// Status
	////
	var cmdStatus = &cobra.Command{
		Use:   "status",
		Short: "Server status",
		Long:  `Mainflux server health checkt.`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			s := cmd.Status()
			fmt.Println(s)
		},
	}

	////
	// Devices
	////
	var cmdDevices = &cobra.Command{
		Use:   "devices",
		Short: "Manipulation with devices.",
		Long:  `Manipulation with devices: create, delete or update devices.`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			s := cmd.GetDevices()
			fmt.Println(s)
		},
	}

	// Create Device
	var cmdCreateDevice = &cobra.Command{
		Use:   "create",
		Short: "create or create <JSON_device>",
		Long:  `Creates new device and generates it's UUID`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			s := cmd.CreateDevice("")
			fmt.Println(s)
		},
	}

	// Get Device
	var cmdGetDevice = &cobra.Command{
		Use:   "get",
		Short: "get or get <device_id>",
		Long:  `Gets all devices or Gets device by id`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l == 0 {
				s = cmd.GetDevices()
			} else {
				for i := 0; i < l; i++ {
					s = s + cmd.GetDevice(args[i])
				}
			}
			fmt.Println(s)
		},
	}

	// Delete Device
	var cmdDeleteDevice = &cobra.Command{
		Use:   "delete",
		Short: "delete all or delete <device_id>",
		Long:  `Removes device from DB`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l != 0 {
				if args[0] == "all" {
					s = cmd.DeleteAllDevices();
				} else {
					for i := 0; i < l; i++ {
						s = s + cmd.DeleteDevice(args[i])
					}
				}
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	// Plug Device
	var cmdPlugDevice = &cobra.Command{
		Use:   "plug",
		Short: "plug <device_d> <JSON_channels> ",
		Long:  `Plugs device into the channel(s)`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l == 2 {
				s = cmd.PlugDevice(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	// Update Device
	var cmdUpdateDevice = &cobra.Command{
		Use:   "update",
		Short: "update <device_id> <JSON_string>",
		Long:  `Update device record`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l == 2 {
				s = cmd.UpdateDevice(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
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
			s := cmd.GetChannels()
			fmt.Println(s)
		},
	}

	// Create Channel
	var cmdCreateChannel = &cobra.Command{
		Use:   "create",
		Short: "create or create <JSON_channel>",
		Long:  `Creates new channel and generates it's UUID`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			s = cmd.CreateChannel("")
			fmt.Println(s)
		},
	}

	// Get Channel
	var cmdGetChannel = &cobra.Command{
		Use:   "get",
		Short: "get or get <channel_id>",
		Long:  `Gets all channels or Gets channel by id`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l == 0 {
				s = cmd.GetChannels()
			} else {
				for i := 0; i < l; i++ {
					s = s + cmd.GetChannel(args[i])
				}
			}
			fmt.Println(s)
		},
	}

	// Update Channel
	var cmdUpdateChannel = &cobra.Command{
		Use:   "update",
		Short: "update <channel_id> <JSON_string>",
		Long:  `Updates channel record`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l == 2 {
				s = cmd.UpdateChannel(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	// Delete Channel
	var cmdDeleteChannel = &cobra.Command{
		Use:   "delete",
		Short: "delete <channel_id>",
		Long:  `Removes channel`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l != 0 {
				if args[0] == "all" {
					s = cmd.DeleteAllChannels();
				} else {
					for i := 0; i < l; i++ {
						s = s + cmd.DeleteChannel(args[i])
					}
				}
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	// Plug Channel
	var cmdPlugChannel = &cobra.Command{
		Use:   "plug",
		Short: "plug <channel_id> <JSON_device>",
		Long:  `Plugs device(s) into the channel`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l == 2 {
				s = cmd.PlugChannel(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	////
	// Messages
	////
	var cmdMessages = &cobra.Command{
		Use:   "msg",
		Short: "Send or retrieve messages",
		Long:  `Send or retrieve messages: controll message flow on the channel`,
	}

	// Get Message
	var cmdGetMessage = &cobra.Command{
		Use:   "get",
		Short: "get <channel_id>",
		Long:  `Gets all messages from a given channel`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l > 0 {
				s = cmd.GetMsg(args[0], startTime, endTime)
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	// Send Message
	var cmdSendMessage = &cobra.Command{
		Use:   "send",
		Short: "send <channel_id> <JSON_string>",
		Long:  `Sends message on the channel`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l > 1 {
				s = cmd.SendMsg(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	////
	// Users
	////
	var cmdCreateUser = &cobra.Command{
		Use:   "user",
		Short: "user <name> <password>",
		Long:  `Creates new user`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			var httpPort = 8180
			cmd.SetServerAddr(httpHost, httpPort)
			l := len(args)
			if l > 1 {
				s = cmd.CreateUser(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	////
	// Sessions
	////
	var cmdInitSession = &cobra.Command{
		Use:   "session",
		Short: "session <name> <password>",
		Long:  `Creates new user`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			var httpPort = 8180
			cmd.SetServerAddr(httpHost, httpPort)
			l := len(args)
			if l > 1 {
				s = cmd.LogginInUser(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	////
	// ApiKeys
	////
	var cmdApiKeys = &cobra.Command{
		Use:   "apikeys",
		Short: "apikeys <authorization>",
		Long:  `Get API key`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			var httpPort = 8180
			cmd.SetServerAddr(httpHost, httpPort)
			l := len(args)
			if l == 1 {
				s = cmd.GetApiKeys(args[0])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	////
	// Create API key
	////
	var cmdCreateApiKeys = &cobra.Command{
		Use:   "create",
		Short: "apikeys <authorization> <JSON_owner>",
		Long:  `Get API key`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			var httpPort = 8180
			cmd.SetServerAddr(httpHost, httpPort)
			l := len(args)
			if l == 2 {
				s = cmd.CreateApiKeys(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	////
	// Delete ApiKeys
	////
	var cmdDeleteApiKeys = &cobra.Command{
		Use:   "delete",
		Short: "delete <authorization> <key>",
		Long:  `Delete API key`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			var httpPort = 8180
			cmd.SetServerAddr(httpHost, httpPort)
			l := len(args)
			if l == 2 {
				s = cmd.DeleteApiKeys(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	////
	// Get ApiKeys
	////
	var cmdGetApiKeys = &cobra.Command{
		Use:   "owner",
		Short: "owner <authorization> <key>",
		Long:  `Get API key Owner`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			var httpPort = 8180
			cmd.SetServerAddr(httpHost, httpPort)
			l := len(args)
			if l == 2 {
				s = cmd.GetOwnerApiKeys(args[0], args[1])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	////
	// Update ApiKeys
	////
	var cmdUpdateApiKeys = &cobra.Command{
		Use:   "update",
		Short: "update <authorization> <key> <JSON_owner>",
		Long:  `Get Owner`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			var httpPort = 8180
			cmd.SetServerAddr(httpHost, httpPort)
			l := len(args)
			if l == 3 {
				s = cmd.UpdateOwnerApiKeys(args[0], args[1], args[2])
			} else {
				s = "Usage: " + cmdCobra.Short
			}
			fmt.Println(s)
		},
	}

	var rootCmd = &cobra.Command{Use: "maninflux-cli"}

	rootCmd.AddCommand(cmdStatus)
	rootCmd.AddCommand(cmdDevices)
	rootCmd.AddCommand(cmdChannels)
	rootCmd.AddCommand(cmdMessages)
	rootCmd.AddCommand(cmdCreateUser)
	rootCmd.AddCommand(cmdInitSession)
	rootCmd.AddCommand(cmdCreateUser)
	rootCmd.AddCommand(cmdApiKeys)

	cmdDevices.AddCommand(cmdCreateDevice)
	cmdDevices.AddCommand(cmdGetDevice)
	cmdDevices.AddCommand(cmdUpdateDevice)
	cmdDevices.AddCommand(cmdDeleteDevice)
	cmdDevices.AddCommand(cmdPlugDevice)

	cmdChannels.AddCommand(cmdCreateChannel)
	cmdChannels.AddCommand(cmdGetChannel)
	cmdChannels.AddCommand(cmdUpdateChannel)
	cmdChannels.AddCommand(cmdDeleteChannel)
	cmdChannels.AddCommand(cmdPlugChannel)

	cmdMessages.AddCommand(cmdGetMessage)
	cmdMessages.AddCommand(cmdSendMessage)

	cmdApiKeys.AddCommand(cmdCreateApiKeys)
	cmdApiKeys.AddCommand(cmdDeleteApiKeys)
	cmdApiKeys.AddCommand(cmdGetApiKeys)
	cmdApiKeys.AddCommand(cmdUpdateApiKeys)

	cmdGetMessage.Flags().StringVarP(
		&startTime, "start", "s", "", "start_time query parameter")
	cmdGetMessage.Flags().StringVarP(
		&endTime, "end", "e", "", "end_time query parameter")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var banner = `
███╗   ███╗ █████╗ ██╗███╗   ██╗███████╗██╗     ██╗   ██╗██╗  ██╗
████╗ ████║██╔══██╗██║████╗  ██║██╔════╝██║     ██║   ██║╚██╗██╔╝
██╔████╔██║███████║██║██╔██╗ ██║█████╗  ██║     ██║   ██║ ╚███╔╝
██║╚██╔╝██║██╔══██║██║██║╚██╗██║██╔══╝  ██║     ██║   ██║ ██╔██╗
██║ ╚═╝ ██║██║  ██║██║██║ ╚████║██║     ███████╗╚██████╔╝██╔╝ ██╗
╚═╝     ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚═╝     ╚══════╝ ╚═════╝ ╚═╝  ╚═╝

                == Industrial IoT System ==
               Made with <3 by Mainflux Team

[w] http://mainflux.io
[t] @mainflux
`
