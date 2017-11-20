# Mainflux CLI

[![License](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/mainflux/bashflux.svg?branch=master)](https://travis-ci.org/mainflux/bashflux)
[![Go Report Card](https://goreportcard.com/badge/github.com/mainflux/bashflux)](https://goreportcard.com/report/github.com/mainflux/bashflux)
[![Join the chat at https://gitter.im/mainflux/mainflux](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Command line interface (CLI) for Mainflux system.

### Installation
#### Prerequisite
If not set already, please set your `GOPATH` and `GOBIN` environment variables. For example:
```bash
mkdir -p ~/go
export GOPATH=~/go
export GOBIN=$GOPATH/bin
# It's often useful to add $GOBIN to $PATH
export PATH=$PATH:$GOBIN
```

#### Get the code
Use [`go`](https://golang.org/cmd/go/) tool to "get" (i.e. fetch and build) `bashflux` package:
```bash
go get github.com/mainflux/bashflux
```

This will download the code to `$GOPATH/src/github.com/mainflux/bashflux` directory,
and then compile it and install the binary in `$GOBIN` directory.

Now you can run the program with:
```
bashflux
```
if `$GOBIN` is in `$PATH` (otherwise use `$GOBIN/bashflux`)

### Documentation
Development documentation can be found [here](http://mainflux.io/).

### Community
#### Mailing lists
[mainflux](https://groups.google.com/forum/#!forum/mainflux) Google group.

#### IRC
[Mainflux Gitter](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

#### Twitter
[@mainflux](https://twitter.com/mainflux)

### License
[Apache License, version 2.0](LICENSE)
