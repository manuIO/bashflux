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

	"github.com/abiosoft/ishell"
	"github.com/mainflux/mainflux-cli/cmd"
	"github.com/mainflux/mainflux-cli/config"
	"github.com/mainflux/mainflux-core/models"
)

func main() {

	// Parse config
	var cfg config.Config
	cfg.Parse()

	// Set HTTP server address
	cmd.SetServerAddr(cfg.HTTPHost, cfg.HTTPPort)

	shell := ishell.New()

	// display info
	shell.Println(banner)

	////
	// Status
	////
	shell.Register("status", func(args ...string) (string, error) {
		s := cmd.Status()
		return s, nil
	})

	////
	// Help
	////
	shell.Register("help", func(args ...string) (string, error) {
		return cliHelp, nil
	})

	////
	// Devices
	////
	shell.Register("devices", func(args ...string) (string, error) {
		var s string
		l := len(args)

		if l == 0 {
			s = cmd.GetDevices()
			return s, nil
		}

		switch args[0] {
		case "get":
			if l > 1 {
				s = cmd.GetDevice(args[1])
				break
			}
			s = cmd.GetDevices()
			break
		case "create":
			if l > 1 {
				s = cmd.CreateDevice(args[1])
			} else {
				s = cmd.CreateDevice("")
			}
			break
		case "update":
			if l > 2 {
				s = cmd.UpdateDevice(args[1], args[2])
			}
			break
		case "delete":
			if args[1] == "all" {
				var devices []models.Device
				s = cmd.GetDevices()
				json.Unmarshal([]byte(s), &devices);
				for i := 0; i  < len(devices); i++ {
					s = s + cmd.DeleteDevice(devices[i].ID)
				}
			} else if l > 1 {
				for i := 1; i < l; i++ {
					s = s + cmd.DeleteDevice(args[i])
				}
				break
			} else {
				s = "usage: devices delete [all] device_id1 device_id2 ..."
			}
			break
		case "plug":
			if l > 2 {
				s = cmd.PlugDevice(args[1], args[2])
			}
			break
		case "help":
			s = devicesHelp

		default:
			s = "Unrecognized command"
		}

		return s, nil
	})

	////
	// Channels
	////
	shell.Register("channels", func(args ...string) (string, error) {
		var s string
		l := len(args)

		if l == 0 {
			s = cmd.GetChannels()
			return s, nil
		}

		switch args[0] {
		case "get":
			if l > 1 {
				s = cmd.GetChannel(args[1])
				break
			}
			s = cmd.GetChannels()
			break
		case "create":
			if l > 1 {
				s = cmd.CreateChannel(args[1])
			} else {
				s = cmd.CreateChannel("")
			}
			break
		case "update":
			if l > 2 {
				s = cmd.UpdateChannel(args[1], args[2])
			}
			break
		case "delete":
			if args[1] == "all" {
				var channels []models.Channel
				s = cmd.GetChannels()
				json.Unmarshal([]byte(s), &channels);
				for i := 0; i  < len(channels); i++ {
					s = s + cmd.DeleteChannel(channels[i].ID)
				}
			} else if l > 1 {
				for i := 1; i < l; i++ {
					s = s + cmd.DeleteChannel(args[i])
				}
				break
			} else {
				s = "usage: channels delete [all] channel_id1 channel_id2 ..."
			}
			break
		case "plug":
			if l > 2 {
				s = cmd.PlugChannel(args[1], args[2])
			}
			break
		case "help":
			s = channelsHelp

		default:
			s = "Unrecognized command"
		}

		return s, nil
	})

	// start shell
	shell.Start()
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

var cliHelp = `
Mainflux CLI is a command line tool for administration and provisioning of
Mainflux IoT server. More information can be found on project's website:
http://mainflux.io

COMMANDS
	status      Mainflux server health check
	devices     Manipulation with devices. Do 'devices help' for info.
	channels    Manipulation with channels. Do 'channels help' for info.

	clear       Clears the screen
	exit        Exits the CLI

	help        Prints this help
`

var devicesHelp = `
COMMANDS
	create                                  Creates new device and generates it's UUID
	get                                     Gets all devices
	get <device_id>                         Gets device by id
	update <device_id> <JSON_string>        Updates device record
	delete <device_id                       Removes device
	plug <device_d> <JSON_channels_list>    Plugs device into the channel(s)
`

var channelsHelp = `
COMMANDS
	create                                  Creates new channel and generates it's UUID
	get                                     Gets all channels
	get <channel_id>                        Gets channel by id
	update <channel_id> <JSON_string>       Updates channel record
	delete <channel_id>                     Removes channel
	plug <channel_id> <JSON_devices_list>   Plugs device(s) into the channel
`
