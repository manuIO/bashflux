# Bashflux

[![License](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/mainflux/bashflux.svg?branch=master)](https://travis-ci.org/mainflux/bashflux)
[![Go Report Card](https://goreportcard.com/badge/github.com/mainflux/bashflux)](https://goreportcard.com/report/github.com/mainflux/bashflux)
[![Join the chat at https://gitter.im/mainflux/mainflux](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Command line interface (CLI) for Mainflux system.

### Service
* Get the service verison: `bashflux version`

### User management
* `bashflux users create john.doe@email.com password`
* `bashflux tokens create john.doe@email.com password`

### System provisioning
* Provisioning devices: `bashflux clients create '{"type":"device", "name":"nyDevice"}' <user_auth_token>`
* Provisioning applications: `bashflux clients create '{"type":"app", "name":"nyDevice"}' <user_auth_token>`
* Retrieving provisioned clients: `bashflux clients get --offset=1 --limit=5 <user_auth_token>`
* Retrieving a specific client: `bashflux clients get <client_id>  --offset=1 --limit=5 <user_auth_token>`
* Removing clients: ``bashflux clients delete <client_id> <user_auth_token>``

* Provisioning devices: `bashflux channels create '{"name":"nyChannel"}' <user_auth_token>`
* Provisioning applications: `bashflux channels create '{"name":"nyChannel"}' <user_auth_token>`
* Retrieving provisioned channels: `bashflux channels get --offset=1 --limit=5 <user_auth_token>`
* Retrieving a specific channel: `bashflux channels get <channel_id>  --offset=1 --limit=5 <user_auth_token>`
* Removing channels: `bashflux channels delete <channel_id> <user_auth_token>`

### Access control
* Connect client to a channel: `bashflux client connect <client_id> <chanel_id <user_auth_token>`
* Disconnect client from channel: `bashflux client disconnect <client_id> <chanel_id <user_auth_token>`

* Send message: `bashflux msg send <channel_id> '[{"bn":"some-base-name:","bt":1.276020076001e+09, "bu":"A","bver":5, "n":"voltage","u":"V","v":120.1}, {"n":"current","t":-5,"v":1.2}, {"n":"current","t":-4,"v":1.3}]' <client_auth_token>`
