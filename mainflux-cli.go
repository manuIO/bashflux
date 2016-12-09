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

	// handle login
	shell.Register("status", func(args ...string) (string, error) {
		s := cmd.Status()
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
