/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mainflux/mainflux-cli/cmd"
	"github.com/mainflux/mainflux-core/models"
	"github.com/spf13/cobra"
)

func main() {

	var httpHost = "0.0.0.0"
	var httpPort = 7070

	var s string

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
				var devices []models.Device
				s = cmd.GetDevices()
				json.Unmarshal([]byte(s), &devices);
				for i := 0; i  < len(devices); i++ {
					s = s + cmd.DeleteDevice(devices[i].ID)
				}
			} else {
				for i := 0; i < l; i++ {
					s = s + cmd.DeleteDevice(args[i])
				}
			}
			fmt.Println(s)
			}
		},
	}

	// Plug Device
	var cmdPlugDevice = &cobra.Command{
		Use:   "plug",
		Short: "plug <device_d> <JSON_channels> ",
		Long:  `Plugs device into the channel(s)`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l > 2 {
				s = cmd.PlugDevice(args[0], args[1])
				fmt.Println(s)
			}
		},
	}

	// Update Device
	var cmdUpdateDevice = &cobra.Command{
		Use:   "update",
		Short: "update <device_id> <JSON_string>",
		Long:  `Update device record`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l > 2 {
				s = cmd.UpdateDevice(args[0], args[1])
				fmt.Println(s)
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

	// Get Device
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
			if l > 2 {
				s = cmd.UpdateChannel(args[0], args[1])
				fmt.Println(s)
			}
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
					var channels []models.Device
					s = cmd.GetChannels()
					json.Unmarshal([]byte(s), &channels);
					for i := 0; i  < len(channels); i++ {
						s = s + cmd.DeleteChannel(channels[i].ID)
					}
				} else {
					for i := 0; i < l; i++ {
						s = s + cmd.DeleteChannel(args[i])
					}
				}
				fmt.Println(s)
			}
		},
	}

	// Plug Channel
	var cmdPlugChannel = &cobra.Command{
		Use:   "plug",
		Short: "plug <channel_id> <JSON_device>",
		Long:  `Plugs device(s) into the channel`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			l := len(args)
			if l > 2 {
				s = cmd.PlugChannel(args[0], args[1])
				fmt.Println(s)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "maninflux-cli"}

	rootCmd.AddCommand(cmdStatus)
	rootCmd.AddCommand(cmdDevices)
	rootCmd.AddCommand(cmdChannels)

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

Mainflux CLI is a command line tool for administration and provisioning of
Mainflux IoT server. More information can be found on project's website:
http://mainflux.io
`
