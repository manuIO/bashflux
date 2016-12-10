/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"github.com/abiosoft/ishell"
	"github.com/mainflux/mainflux-cli/cmd"
	"github.com/mainflux/mainflux-cli/config"
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
			if l > 1 {
				s = cmd.DeleteDevice(args[1])
				break
			}
			break
		default:
			s = cmd.GetDevice(args[0])
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
			s = cmd.GetDevices()
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
			if l > 1 {
				s = cmd.DeleteChannel(args[1])
				break
			}
			break
		default:
			s = cmd.GetChannel(args[0])
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
